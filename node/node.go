package node

import (
	"context"
	"encoding/hex"
	"net"

	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/crypto"
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

	validatorSet map[string]bool
	round        int

	*crypto.Keys

	*Cache
	*Store
	Mempool *transactions.Mempool

	logger *zap.Logger

	proto.UnimplementedNodeServer
}

func NewNode(
	listenAddr string,
	height int64,
	keys *crypto.Keys,
	cache *Cache,
	store *Store,
	mempool *transactions.Mempool,
	logger *zap.Logger) *Node {
	return &Node{
		ListenAddr: listenAddr,
		Height:     height,
		Keys:       keys,
		Cache:      cache,
		Store:      store,
		Mempool:    mempool,
		logger:     logger,
	}
}

func (n *Node) Start(peerAddrs []string) error {
	errChan := make(chan error, 2)

	n.logger.Info("Starting peer discovery go routine", zap.Strings("peers", peerAddrs))

	go func(peerAddrs []string, errCh chan error) {
		err := n.peerDiscovery(peerAddrs)
		if err != nil {
			n.logger.Error("Error in start peer discovery, exiting", zap.Error(err))
			errCh <- err
			return
		}
	}(peerAddrs, errChan)

	n.logger.Info("Starting gRPC go routine", zap.String("listenAddr", n.ListenAddr))

	go func(errCh chan error) {
		grpcServer := grpc.NewServer()
		proto.RegisterNodeServer(grpcServer, n)

		lis, err := net.Listen("tcp", n.ListenAddr)
		if err != nil {
			n.logger.Error("Error creating tcp listener", zap.Error(err))
			errCh <- err
			return
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			n.logger.Error("Error starting grpc server", zap.Error(err))
			errCh <- err
			return
		}
	}(errChan)

	n.logger.Info("Completed start sequence")

	for {
		err := <-errChan
		if err != nil {
			return err
		}
	}
}

func (n *Node) HandleBlock(ctx context.Context, protoBlock *proto.Block) (*proto.Ack, error) {
	ack := &proto.Ack{}
	block := blocks.ConvertProtoBlock(protoBlock)
	blockID, err := block.Hash()
	if err != nil {
		n.logger.Error("Handle Block: error hashing block", zap.Error(err))
		return ack, err
	}
	blockIDStr := hex.EncodeToString(blockID)

	// check if already seen this block
	ok := n.BlockCache.Has(blockID)
	if ok {
		return ack, nil
	} else {
		n.BlockCache.Put(blockID, true)
	}

	n.logger.Info("Received new block, validating block ...", zap.String("blockID", blockIDStr))

	// validate block
	ok = n.validateBlock(block)
	if !ok {
		return ack, nil
	}

	n.logger.Info("Successfully validated block, broadcasting block ...", zap.String("blockID", blockIDStr))

	// step 3: broadcast block
	err = n.broadcastBlock(protoBlock)
	if err != nil {
		n.logger.Error("Error broadcasting block", zap.String("blockID", blockIDStr), zap.Error(err))
	}

	n.logger.Info("Successfully broadcasted block, broadcasting vote ...", zap.String("blockID", blockIDStr))

	// step 2: vote for block
	err = n.broadcastVote(blockID, true)
	if err != nil {
		n.logger.Error("Error broadcasting vote", zap.String("blockID", blockIDStr), zap.Error(err))
	}

	n.logger.Info("Successfully broadcasted vote", zap.String("blockID", blockIDStr))

	return ack, nil
}

func (n *Node) validateBlock(b *blocks.Block) bool {
	// step 1: verify signature

	// step 2: verify hash of merkle tree root

	// step 3: validate prev block hash + height

	// step 4: validate timestamp

	// step 5: validate transactions
	// - double spend (txs gone through, txs in mempool, txs in same block)
	return false
}

func (n *Node) broadcastBlock(protoBlock *proto.Block) error {
	for _, p := range n.PeerCache.Cache {
		client := p.Client
		_, err := client.HandleBlock(context.Background(), protoBlock)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) broadcastVote(blockID []byte, vote bool) error {
	protoVote := &proto.Vote{
		BlockID: blockID,
		Vote:    vote,
	}
	for _, p := range n.PeerCache.Cache {
		client := p.Client
		_, err := client.HandleVote(context.Background(), protoVote)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) HandleVote(ctx context.Context, protoVote *proto.Vote) (*proto.Ack, error) {
	ack := &proto.Ack{}
	return ack, nil
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

	n.logger.Info("Received new transaction, validating transaction ...", zap.String("txID", txIDStr))

	// validate transaction
	ok = n.validateTransaction(tx, txID)
	if !ok {
		return ack, nil
	}

	n.logger.Info("Successfully validated transaction, adding to mempool", zap.String("txID", txIDStr))

	// add valid transaction to mempool
	n.Mempool.PutTransaction(txID, tx)

	n.logger.Info("Successfully added transaction to mempool, marking ref utxos as claimed", zap.String("txID", txIDStr), zap.Int("mempool size", n.Mempool.TransactionsSize()))

	// mark all referenced utxos as mempool claimed
	for _, input := range tx.Inputs {
		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		n.Mempool.PutUTXO(utxoID, true)
	}

	n.logger.Info("Successfully marked ref utxos, broadcasting transaction", zap.String("txID", txIDStr))

	// broadcast transactions to peers
	err := n.broadcastTransaction(protoTx)
	if err != nil {
		n.logger.Error("Error broadcasting transaction", zap.String("txID", txIDStr), zap.Error(err))
		return ack, nil
	}

	return ack, nil
}

func (n *Node) validateTransaction(tx *transactions.Transaction, txID []byte) bool {
	ok := n.validateTransactionInputs(tx.Inputs, txID)
	if !ok {
		return false
	}

	ok = n.validateTransactionOutputs(tx.Inputs, tx.Outputs, txID)
	if !ok {
		return false
	}

	return true
}

func (n *Node) validateTransactionInputs(inputs []*transactions.Input, txID []byte) bool {
	txIDStr := hex.EncodeToString(txID)

	var seenUTXOs map[string]bool
	for i, input := range inputs {
		// checks that the input signature is valid
		if !input.UnlockingScript.Sig.Verify(input.UnlockingScript.PubKey, txID) {
			n.logger.Error("Error validating transaction: transaction contains input with invalid signature",
				zap.String("txID", txIDStr),
				zap.Int("inputIndex", i))
			return false
		}

		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		utxoIDStr := hex.EncodeToString(utxoID)

		// checks that no two inputs reference same utxo
		_, ok := seenUTXOs[utxoIDStr]
		if ok {
			n.logger.Error(
				"Error validating transaction: transaction contains multiple inputs referencing the same utxo",
				zap.String("txID", txIDStr),
				zap.Int("inputIndex", i),
				zap.String("utxoID", utxoIDStr))
			return false
		} else {
			seenUTXOs[utxoIDStr] = true
		}

		// checks that the referenced utxo exists/is not spent
		utxo, err := n.UtxoStore.Get(utxoID)
		if err != nil {
			n.logger.Error("Error validating transaction: transaction contains input referencing nonexistent utxo",
				zap.Error(err),
				zap.String("txID", txIDStr),
				zap.Int("inputIndex", i),
				zap.String("utxoID", utxoIDStr))
			return false
		}

		// checks that the referenced utxo is not claimed by other transaction already in mempool
		if n.Mempool.HasUTXO(utxoID) {
			n.logger.Error("Error validating transaction: transaction contains mempool claimed utxo",
				zap.String("txID", txIDStr),
				zap.Int("inputIndex", i),
				zap.String("utxoID", utxoIDStr))
			return false
		}

		// checks that the tx owner has the right to spend referenced utxo
		if !input.UnlockingScript.Unlock(utxo.LockingScript) {
			uPubKeyHashStr := hex.EncodeToString(input.UnlockingScript.PubKey.Hash())
			lPubKeyHashStr := hex.EncodeToString(utxo.LockingScript.PubKeyHash)
			n.logger.Error("Error validating transaction: transaction contains input that does not unlock referenced utxo",
				zap.String("txID", txIDStr),
				zap.Int("inputIndex", i),
				zap.String("unlocking script pub key hash", uPubKeyHashStr),
				zap.String("utxoID", utxoIDStr),
				zap.String("locking script pub key hash", lPubKeyHashStr))
			return false
		}
	}
	return true
}

func (n *Node) validateTransactionOutputs(inputs []*transactions.Input, outputs []*transactions.Output, txID []byte) bool {
	txIDStr := hex.EncodeToString(txID)

	var totalAmount int64
	for _, input := range inputs {
		utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
		utxo, _ := n.UtxoStore.Get(utxoID)
		totalAmount += utxo.Amount
	}

	for i, output := range outputs {
		// checks output amount is non-negative
		amount := output.Amount
		if amount < 0 {
			n.logger.Error("Error validating transaction: transaction contains output with negative amount",
				zap.String("txID", txIDStr),
				zap.Int("outputIndex", i))
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

func (n *Node) broadcastTransaction(protoTx *proto.Transaction) error {
	for _, p := range n.PeerCache.Cache {
		client := p.Client
		_, err := client.HandleTransaction(context.Background(), protoTx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) HandleHandshake(ctx context.Context, peerMsg *proto.Handshake) (*proto.Handshake, error) {
	peerAddr := peerMsg.ListenAddr
	msg := &proto.Handshake{
		ListenAddr:      n.ListenAddr,
		PeerListenAddrs: n.PeerCache.GetPeerAddrs(),
		Height:          n.Height,
	}

	// checks that we do not already know this peer
	seen := n.PeerCache.Has(peerAddr)
	if seen {
		return msg, nil
	}

	// checks own handshake message has not been forwarded back to us
	if peerAddr == n.ListenAddr {
		return msg, nil
	}

	n.logger.Info("Received new handshake, dialing peer", zap.String("peer", peerAddr))

	// dial peer
	client, err := n.dialPeer(peerAddr)
	if err != nil {
		n.logger.Error("Error dialing peer", zap.Error(err), zap.String("peer", peerAddr))
		return msg, nil
	}

	n.logger.Info("Successfully dialed peer, adding peer", zap.String("peer", peerAddr))

	// add peer
	p := peer.NewPeer(peerMsg.Height, client)
	n.Cache.PeerCache.Put(peerAddr, p)

	n.logger.Info("Successfully added peer, starting peer discovery go routine", zap.String("peer", peerAddr), zap.Strings("peers", peerMsg.PeerListenAddrs))

	// start peer discovery go routine
	go func() {
		err := n.peerDiscovery(peerMsg.PeerListenAddrs)
		if err != nil {
			n.logger.Error("Error in peer discover", zap.Error(err), zap.String("peer", peerAddr), zap.Strings("peers", peerMsg.PeerListenAddrs))
		}
	}()

	// update list of peers
	msg.PeerListenAddrs = n.PeerCache.GetPeerAddrs()
	return msg, nil
}

func (n *Node) peerDiscovery(peerAddrs []string) error {
	for _, peerAddr := range peerAddrs {
		seen := n.PeerCache.Has(peerAddr)
		// checks we do not already know this peer
		if seen {
			continue
		}
		// checks peer addr is not our addr
		if peerAddr == n.ListenAddr {
			continue
		}

		// dial peer
		client, err := n.dialPeer(peerAddr)
		if err != nil {
			return err
		}

		msg := &proto.Handshake{
			ListenAddr:      n.ListenAddr,
			PeerListenAddrs: n.PeerCache.GetPeerAddrs(),
			Height:          n.Height,
		}

		// invoke peer handshake method
		peerMsg, err := client.HandleHandshake(context.Background(), msg)
		if err != nil {
			continue
		}

		// add new peer
		n.logger.Info("Adding peer", zap.String("peer", peerAddr))
		p := peer.NewPeer(peerMsg.Height, client)
		n.Cache.PeerCache.Put(peerAddr, p)

		err = n.peerDiscovery(peerMsg.PeerListenAddrs)
		if err != nil {
			return err
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
