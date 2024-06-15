package node

import (
	"context"
	"encoding/hex"
	"net"

	"github.com/fahyjo/blockchain/peer"
	proto "github.com/fahyjo/blockchain/proto"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	ListenAddr string
	Height     int64

	*Cache
	*Store

	logger *zap.Logger

	proto.UnimplementedNodeServer
}

func NewNode(
	listenAddr string,
	height int64,
	cache *Cache,
	store *Store,
	logger *zap.Logger) *Node {
	return &Node{
		ListenAddr: listenAddr,
		Height:     height,
		Cache:      cache,
		Store:      store,
		logger:     logger,
	}
}

func (n *Node) Start(peerAddrs []string) error {
	errChan := make(chan error, 3)

	n.logger.Info("Starting peer discovery go routine", zap.Strings("peers", peerAddrs))
	go func(peerAddrs []string, errCh chan error) {
		err := n.peerDiscovery(peerAddrs)
		if err != nil {
			n.logger.Error("Error in peer discovery, exiting", zap.String("error", err.Error()))
			errCh <- err
		}
	}(peerAddrs, errChan)

	n.logger.Info("Starting gRPC server go routine", zap.String("listenAddr", n.ListenAddr))
	go func(errCh chan error) {
		grpcServer := grpc.NewServer()
		proto.RegisterNodeServer(grpcServer, n)

		lis, err := net.Listen("tcp", n.ListenAddr)
		if err != nil {
			errCh <- err
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			errCh <- err
		}
	}(errChan)

	for {
		err := <-errChan
		if err != nil {
			return err
		}
	}
}

func (n *Node) HandleTransaction(ctx context.Context, protoTx *proto.Transaction) (*proto.Ack, error) {
	tx := transactions.ConvertProtoTransaction(protoTx)
	txID, _ := tx.Hash()
	txIDStr := hex.EncodeToString(txID)
	ack := &proto.Ack{}

	// check if already seen this transaction
	ok := n.TransactionCache.Has(txID)
	if ok {
		return ack, nil
	}
	n.TransactionCache.Put(txID, true)

	n.logger.Info("Received transaction, validating ...", zap.String("txID", txIDStr))

	ok = n.validateTransactionInputs(tx)
	if !ok {
		return ack, nil
	}

	ok = n.validateTransactionOutputs(tx)
	if !ok {
		return ack, nil
	}

	n.logger.Info("Successfully validated transaction, adding to mempool", zap.String("txID", txIDStr))

	// add valid transaction to mempool
	err := n.Mempool.Put(tx)
	if err != nil {
		n.logger.Error("Error post transaction validation: error adding valid transaction to mempool",
			zap.String("txID", txIDStr))
		return ack, nil
	}

	n.logger.Info("Successfully added transaction to mempool, marking ref utxos as claimed", zap.String("txID", txIDStr), zap.Int("mempool size", n.Mempool.Size()))

	// mark all referenced utxos as mempool claimed
	for _, input := range tx.Inputs {
		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		utxoIDStr := hex.EncodeToString(utxoID)
		utxo, _ := n.UtxoStore.Get(utxoID)
		utxo.MempoolClaimed = true
		err := n.UtxoStore.Put(utxoID, utxo)
		if err != nil {
			n.logger.Error("Error post transaction validation: error overwriting utxo in utxo set after marking mempool claimed",
				zap.String("txID", txIDStr),
				zap.String("utxoID", utxoIDStr))
		}
	}

	n.logger.Info("Broadcasting transaction", zap.String("txID", txIDStr))

	ok = n.broadcastTransaction(tx)
	if !ok {
		return ack, nil
	}

	return ack, nil
}

func (n *Node) validateTransactionInputs(tx *transactions.Transaction) bool {
	txID, err := tx.Hash()
	if err != nil {
		return false
	}
	txIDStr := hex.EncodeToString(txID)

	var seenUTXOs map[string]bool
	for i, input := range tx.Inputs {
		// checks that the input signature is valid
		if !input.UnlockingScript.Sig.Verify(input.UnlockingScript.PubKey, txID) {
			n.logger.Error("Error validating transaction: transaction contains input with invalid signature, transaction not added to mempool",
				zap.String("txID", hex.EncodeToString(txID)),
				zap.Int("input index", i))
			return false
		}

		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		utxoIDStr := hex.EncodeToString(utxoID)

		// checks that no two inputs reference same utxo
		_, ok := seenUTXOs[utxoIDStr]
		if ok {
			n.logger.Error(
				"Error validating transaction: transaction contains multiple inputs referencing the same utxo, transaction not added to mempool",
				zap.String("txID", txIDStr),
				zap.String("utxoID", utxoIDStr))
			return false
		} else {
			seenUTXOs[utxoIDStr] = true
		}

		// checks that the referenced utxo exists/is not spent
		utxo, err := n.UtxoStore.Get(utxoID)
		if err != nil {
			n.logger.Error("Error validating transaction: transaction contains input referencing nonexistent utxo, transaction not added to mempool",
				zap.String("txID", txIDStr),
				zap.String("utxoID", utxoIDStr))
			return false
		}

		// checks that the referenced utxo is not claimed by other transaction already in mempool
		if utxo.MempoolClaimed {
			n.logger.Error("Error validating transaction: transaction contains mempool claimed utxo, transaction not added to mempool",
				zap.String("txID", txIDStr),
				zap.String("utxoID", utxoIDStr))
			return false
		}

		// checks that the tx owner has the right to spend referenced utxo
		if !input.UnlockingScript.Unlock(utxo.LockingScript) {
			uPubKeyHashStr := hex.EncodeToString(input.UnlockingScript.PubKey.Hash())
			lPubKeyHashStr := hex.EncodeToString(utxo.LockingScript.PubKeyHash)
			n.logger.Error("Error validating transaction: transaction contains input that does not unlock referenced utxo, transaction not added to mempool",
				zap.String("txID", txIDStr),
				zap.String("unlocking script pub key hash", uPubKeyHashStr),
				zap.String("utxoID", utxoIDStr),
				zap.String("locking pub key hash", lPubKeyHashStr))
			return false
		}
	}
	return true
}

func (n *Node) validateTransactionOutputs(tx *transactions.Transaction) bool {
	txID, err := tx.Hash()
	if err != nil {
		return false
	}
	txIDStr := hex.EncodeToString(txID)

	var totalAmount int64
	for _, input := range tx.Inputs {
		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		utxo, _ := n.UtxoStore.Get(utxoID)
		totalAmount += utxo.Amount
	}

	for i, output := range tx.Outputs {
		// checks output amount is non-negative
		amount := output.Amount
		if amount < 0 {
			n.logger.Error("Error validating transaction: transaction contains output with negative amount",
				zap.String("txID", txIDStr),
				zap.Int("utxoIndex", i))
			return false
		}
		// checks total output amount is not greater than total input amount
		totalAmount -= amount
		if totalAmount < 0 {
			n.logger.Error("Error validating transaction: transaction output amount greater than input amount",
				zap.String("txID", txIDStr))
			return false
		}
	}
	return true
}

func (n *Node) broadcastTransaction(tx *transactions.Transaction) bool {
	txID, err := tx.Hash()
	if err != nil {
		n.logger.Error("Error broadcasting transaction, unable to calculate txID")
		return false
	}
	txIDStr := hex.EncodeToString(txID)

	protoTx := transactions.ConvertTransaction(tx)
	for peerAddr, peer := range n.PeerCache.Cache {
		client := peer.Client
		_, err := client.HandleTransaction(context.Background(), protoTx)
		if err != nil {
			n.logger.Error("Error broadcasting transaction",
				zap.Error(err),
				zap.String("txID", txIDStr),
				zap.String("peer", peerAddr))
		}
	}
	return true
}

func (n *Node) HandleHandshake(ctx context.Context, peerMsg *proto.Handshake) (*proto.Handshake, error) {
	peerAddr := peerMsg.ListenAddr

	n.logger.Info("Received handshake", zap.String("peer", peerAddr))

	msg := &proto.Handshake{
		ListenAddr:      n.ListenAddr,
		PeerListenAddrs: n.PeerCache.GetPeerAddrs(),
		Height:          n.Height,
	}

	ok := n.PeerCache.Has(peerAddr)
	if !ok && peerAddr != n.ListenAddr {
		client, err := n.dialPeer(peerAddr)
		if err != nil {
			n.logger.Error("Error dialing peer", zap.Error(err), zap.String("peer", peerAddr))
			return msg, nil
		}
		peer := peer.NewPeer(peerMsg.Height, client)

		n.logger.Info("Adding peer", zap.String("peer", peerAddr))

		n.Cache.PeerCache.Put(peerAddr, peer)

		n.logger.Info("Successfully added peer, starting peer discovery go routine", zap.String("peer", peerAddr), zap.Strings("peers", peerMsg.PeerListenAddrs))

		go func() {
			err := n.peerDiscovery(peerMsg.PeerListenAddrs)
			if err != nil {
				n.logger.Error("Error in peer discover", zap.Error(err), zap.String("peer", peerAddr), zap.Strings("peers", peerMsg.PeerListenAddrs))
			}
		}()
	}

	return msg, nil
}

func (n *Node) peerDiscovery(peerAddrs []string) error {
	for _, peerAddr := range peerAddrs {
		ok := n.PeerCache.Has(peerAddr)
		if !ok && peerAddr != n.ListenAddr {
			client, err := n.dialPeer(peerAddr)
			if err != nil {
				return err
			}

			msg := &proto.Handshake{
				ListenAddr:      n.ListenAddr,
				PeerListenAddrs: n.PeerCache.GetPeerAddrs(),
				Height:          n.Height,
			}
			peerMsg, err := client.HandleHandshake(context.Background(), msg)
			if err != nil {
				return err
			}

			peer := peer.NewPeer(peerMsg.Height, client)
			n.Cache.PeerCache.Put(peerAddr, peer)

			err = n.peerDiscovery(peerMsg.PeerListenAddrs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) dialPeer(peerAddr string) (proto.NodeClient, error) {
	conn, err := grpc.NewClient(peerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewNodeClient(conn)
	return client, nil
}
