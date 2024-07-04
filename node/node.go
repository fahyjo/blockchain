package node

import (
	"bytes"
	"context"
	"encoding/hex"
	"net"

	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/consensus"
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

	*crypto.Keys

	consensus *consensus.Consensus

	*Cache
	*Store
	BlockList *blocks.BlockList
	Mempool   *transactions.Mempool

	logger *zap.Logger

	proto.UnimplementedNodeServer
}

func NewNode(
	listenAddr string,
	height int64,
	keys *crypto.Keys,
	consensus *consensus.Consensus,
	cache *Cache,
	store *Store,
	blockList *blocks.BlockList,
	mempool *transactions.Mempool,
	logger *zap.Logger) *Node {
	return &Node{
		ListenAddr: listenAddr,
		Height:     height,
		Keys:       keys,
		consensus:  consensus,
		Cache:      cache,
		Store:      store,
		BlockList:  blockList,
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

	// check that we are in the proposal phase
	if n.consensus.CurrentRound.CurrentPhase.Value() != "proposal" {
		n.logger.Error("Received block, not in proposal phase", zap.String("currentPhase", n.consensus.CurrentRound.CurrentPhase.Value()))
		return ack, nil
	}

	// get block, blockID, blockIDStr
	block := blocks.ConvertProtoBlock(protoBlock)
	blockID, err := block.Hash()
	if err != nil {
		n.logger.Error("Error hashing block", zap.Error(err))
		return ack, err
	}
	blockIDStr := hex.EncodeToString(blockID)

	// check if already seen this block
	ok := n.BlockCache.Has(blockID)
	if ok {
		n.logger.Error("Received block already in cache", zap.String("blockID", blockIDStr))
		return ack, nil
	} else {
		n.BlockCache.Put(blockID, true)
	}

	n.logger.Info("Received new block, validating block ...", zap.String("blockID", blockIDStr))

	// validate block
	ok = n.validateBlock(block, blockID)
	if !ok {
		return ack, nil
	}

	n.logger.Info("Successfully validated block, broadcasting block ...", zap.String("blockID", blockIDStr))

	// broadcast block
	err = n.broadcastBlock(protoBlock)
	if err != nil {
		n.logger.Error("Error broadcasting block", zap.String("blockID", blockIDStr), zap.Error(err))
	}

	n.logger.Info("Broadcasted block, moving to preVote phase ...", zap.String("blockID", blockIDStr))

	// move to preVote phase
	n.consensus.CurrentRound.BlockID = blockID
	n.consensus.CurrentRound.Block = block
	n.consensus.CurrentRound.NextPhase()

	// broadcast preVote if validator
	if n.consensus.AmValidator {
		protoPreVote := &proto.PreVote{
			BlockID: blockID,
			Sig:     n.PrivateKey.Sign(blockID).Bytes(),
			PubKey:  n.PublicKey.Bytes(),
		}
		err = n.broadcastPreVote(protoPreVote)
		if err != nil {
			n.logger.Error("Error broadcasting preVote", zap.String("blockID", blockIDStr), zap.Error(err))
		}
		n.logger.Info("Broadcasted preVote", zap.String("blockID", blockIDStr))

		err = n.consensus.CurrentRound.CurrentPhase.AddValidatorID(hex.EncodeToString(n.PublicKey.Bytes()))
		if err != nil {
			n.logger.Error("Error adding validator", zap.String("blockID", blockIDStr), zap.Error(err))
			return ack, nil
		}
		err = n.consensus.CurrentRound.CurrentPhase.IncrementAttestationCount()
		if err != nil {
			n.logger.Error("Error incrementing attestation count", zap.String("blockID", blockIDStr), zap.Error(err))
			return ack, nil
		}
	}

	return ack, nil
}

func (n *Node) validateBlock(block *blocks.Block, blockID []byte) bool {
	blockIDStr := hex.EncodeToString(blockID)

	// check that the creator of the block is the designated proposer for this round
	if n.consensus.CurrentRound.ProposerID != hex.EncodeToString(block.PubKey.Bytes()) {
		n.logger.Error("Error validating block, creator of block is not the designated proposer for this round",
			zap.String("proposerID", n.consensus.CurrentRound.ProposerID),
			zap.String("blockCreator", hex.EncodeToString(block.PubKey.Bytes())),
			zap.String("blockID", blockIDStr))
		return false
	}
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

func (n *Node) broadcastPreVote(protoPreVote *proto.PreVote) error {
	for _, p := range n.PeerCache.Cache {
		client := p.Client
		_, err := client.HandlePreVote(context.Background(), protoPreVote)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) broadcastPreCommit(protoPreCommit *proto.PreCommit) error {
	for _, p := range n.PeerCache.Cache {
		client := p.Client
		_, err := client.HandlePreCommit(context.Background(), protoPreCommit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) HandlePreVote(ctx context.Context, protoPreVote *proto.PreVote) (*proto.Ack, error) {
	ack := &proto.Ack{}

	// convert protoPreVote to preVote
	preVote := consensus.ConvertProtoPreVote(protoPreVote)
	preVoteSenderIDStr := hex.EncodeToString(preVote.PubKey.Bytes())
	blockIDStr := hex.EncodeToString(preVote.BlockID)

	n.logger.Info("Received preVote, validating ...", zap.String("preVote sender ID", preVoteSenderIDStr), zap.String("blockID", blockIDStr))

	// check that we are in preVote phase
	currentPhase := n.consensus.CurrentRound.CurrentPhase.Value()
	if currentPhase != "preVote" {
		n.logger.Error("Received preVote, not in preVote phase", zap.String("currentPhase", currentPhase))
		return ack, nil
	}

	// check that preVote sender is a validator
	if !n.consensus.IsValidator(preVoteSenderIDStr) {
		n.logger.Error("Received preVote from non validator node", zap.String("preVote sender ID", preVoteSenderIDStr))
		return ack, nil
	}

	// check if we have already seen this preVote from this validator this round
	ok, err := n.consensus.CurrentRound.CurrentPhase.HasValidatorID(preVoteSenderIDStr)
	if err != nil {
		n.logger.Error("Error checking if already seen validator this phase", zap.String("preVote sender ID", preVoteSenderIDStr), zap.Error(err))
		return ack, nil
	}
	if ok {
		n.logger.Error("Already received preVote from this validator this phase", zap.String("preVote sender ID", preVoteSenderIDStr))
		return ack, nil
	}

	// check that the preVote is for the correct block
	if !bytes.Equal(preVote.BlockID, n.consensus.CurrentRound.BlockID) {
		n.logger.Error("Received preVote for block not being considered this round",
			zap.String("proposedBlockID", hex.EncodeToString(n.consensus.CurrentRound.BlockID)),
			zap.String("preVoteBlockID", blockIDStr))
		return ack, nil
	}

	// verify preVote signature
	if !preVote.Sig.Verify(preVote.PubKey, preVote.BlockID) {
		n.logger.Error("PreVote contains invalid signature", zap.String("preVote sender ID", preVoteSenderIDStr))
		return ack, nil
	}

	n.logger.Info("Successfully validated preVote, broadcasting ...", zap.String("preVote sender ID", preVoteSenderIDStr))

	// broadcast valid preVote
	err = n.broadcastPreVote(protoPreVote)
	if err != nil {
		n.logger.Error("Error broadcasting preVote", zap.String("preVote sender ID", preVoteSenderIDStr), zap.Error(err))
	}

	n.logger.Info("Updating phase ...", zap.String("preVote sender ID", preVoteSenderIDStr))

	// add id of new validator
	err = n.consensus.CurrentRound.CurrentPhase.AddValidatorID(preVoteSenderIDStr)
	if err != nil {
		n.logger.Error("Error adding validator", zap.Error(err))
		return ack, nil
	}
	// increment the preVote count
	err = n.consensus.CurrentRound.CurrentPhase.IncrementAttestationCount()
	if err != nil {
		n.logger.Error("Error incrementing attestation count", zap.Error(err))
		return ack, nil
	}

	n.logger.Info("Checking if have received required number of preVotes ...")

	threshold := len(n.consensus.Validators) * 2 / 3
	ok, err = n.consensus.CurrentRound.CurrentPhase.AtAttestationThreshold(threshold)
	if err != nil {
		n.logger.Error("Error checking if at attestation threshold", zap.Error(err))
		return ack, nil
	}
	if !ok {
		n.logger.Info("Did not reach threshold to move to preCommit phase")
		return ack, nil
	}

	n.logger.Info("Received required number of preVotes, moving to preCommit phase ...")

	// move to preCommit phase
	n.consensus.CurrentRound.NextPhase()
	if n.consensus.AmValidator {
		preCommit := &proto.PreCommit{
			BlockID: n.consensus.CurrentRound.BlockID,
			Sig:     n.PrivateKey.Sign(n.consensus.CurrentRound.BlockID).Bytes(),
			PubKey:  n.PublicKey.Bytes(),
		}
		err = n.broadcastPreCommit(preCommit)
		if err != nil {
			n.logger.Error("Error broadcasting own preCommit", zap.Error(err))
		}
		err = n.consensus.CurrentRound.CurrentPhase.AddValidatorID(hex.EncodeToString(n.PublicKey.Bytes()))
		if err != nil {
			n.logger.Error("Error adding validator", zap.Error(err))
			return ack, nil
		}
		err = n.consensus.CurrentRound.CurrentPhase.IncrementAttestationCount()
		if err != nil {
			n.logger.Error("Error incrementing attestation count", zap.Error(err))
			return ack, nil
		}
	}
	return ack, nil
}

func (n *Node) HandlePreCommit(ctx context.Context, protoPreCommit *proto.PreCommit) (*proto.Ack, error) {
	ack := &proto.Ack{}

	// convert protoPreCommit to preCommit
	preCommit := consensus.ConvertProtoPreCommit(protoPreCommit)
	preCommitSenderIDStr := hex.EncodeToString(preCommit.PubKey.Bytes())
	blockIDStr := hex.EncodeToString(preCommit.BlockID)

	n.logger.Info("Received preCommit, validating ...", zap.String("preCommit sender ID", preCommitSenderIDStr), zap.String("blockID", blockIDStr))

	// check that we are in preCommit phase
	currentPhase := n.consensus.CurrentRound.CurrentPhase.Value()
	if currentPhase != "preCommit" {
		n.logger.Error("Received preCommit, not in preCommit phase", zap.String("preCommit sender ID", preCommitSenderIDStr))
		return ack, nil
	}

	// check that the preCommit sender is a validator
	if !n.consensus.IsValidator(preCommitSenderIDStr) {
		n.logger.Error("Received preCommit from non validator node", zap.String("preCommit sender ID", preCommitSenderIDStr))
		return ack, nil
	}

	// check if we have already seen this preCommit from this validator this round
	ok, err := n.consensus.CurrentRound.CurrentPhase.HasValidatorID(preCommitSenderIDStr)
	if err != nil {
		n.logger.Error("Error checking if already seen validator this phase", zap.String("preCommit sender ID", preCommitSenderIDStr), zap.Error(err))
		return ack, nil
	}
	if ok {
		n.logger.Error("Already received preCommit from this validator this phase", zap.String("preCommit sender ID", preCommitSenderIDStr))
		return ack, nil
	}

	// check that the preCommit is for the correct block
	if !bytes.Equal(preCommit.BlockID, n.consensus.CurrentRound.BlockID) {
		n.logger.Error("Received preCommit for block not being considered this round",
			zap.String("proposedBlockID", hex.EncodeToString(n.consensus.CurrentRound.BlockID)),
			zap.String("preCommitBlockID", blockIDStr))
		return ack, nil
	}

	// verify preCommit signature
	if !preCommit.Sig.Verify(preCommit.PubKey, preCommit.BlockID) {
		n.logger.Error("PreCommit contains invalid signature", zap.String("preCommit sender ID", preCommitSenderIDStr))
		return ack, nil
	}

	n.logger.Info("Successfully validated preCommit, broadcasting ...", zap.String("preCommit sender ID", preCommitSenderIDStr))

	// broadcast valid preCommit
	err = n.broadcastPreCommit(protoPreCommit)
	if err != nil {
		n.logger.Error("Error broadcasting preCommit", zap.String("preCommit sender ID", preCommitSenderIDStr), zap.Error(err))
	}

	n.logger.Info("Updating phase ...", zap.String("preCommit sender ID", preCommitSenderIDStr))

	// add id of new validator
	err = n.consensus.CurrentRound.CurrentPhase.AddValidatorID(preCommitSenderIDStr)
	if err != nil {
		n.logger.Error("Error adding validator", zap.Error(err))
		return ack, nil
	}
	// increment the preCommit count
	err = n.consensus.CurrentRound.CurrentPhase.IncrementAttestationCount()
	if err != nil {
		n.logger.Error("Error incrementing attestation count", zap.Error(err))
		return ack, nil
	}

	n.logger.Info("Checking if have received required number of preCommits ...")

	// check if received required number of preCommits
	threshold := len(n.consensus.Validators) * 2 / 3
	ok, err = n.consensus.CurrentRound.CurrentPhase.AtAttestationThreshold(threshold)
	if err != nil {
		n.logger.Error("Error checking if at attestation threshold", zap.Error(err))
		return ack, nil
	}
	if !ok {
		n.logger.Info("Did not reach threshold to commit block")
		return ack, nil
	}

	n.logger.Info("Received required number of preCommits, committing block ...", zap.String("blockID", blockIDStr))

	ok = n.commitBlock(n.consensus.CurrentRound.Block, n.consensus.CurrentRound.BlockID)

	// clean mempool

	if !ok {
		panic("Block commit failed, state fucked")
	}

	// move to next round
	n.consensus.NextRound()

	return ack, nil
}

func (n *Node) commitBlock(block *blocks.Block, blockID []byte) bool {
	// add block to block store
	err := n.BlockStore.Put(n.consensus.CurrentRound.Block)
	if err != nil {
		n.logger.Error("Error adding block to block store", zap.Error(err))
		return false
	}

	// add block to block list
	err = n.BlockList.Add(blockID)
	if err != nil {
		n.logger.Error("Error adding block to block list, walking back block commit ...", zap.Error(err))
		err = n.walkBackCommit(blockID, nil, nil, nil, nil)
		if err != nil {
			n.logger.Error("Error walking back block commit", zap.Error(err))
			return false
		}
		return false
	}

	var committedTXs [][]byte
	var consumedUTXOs map[string]*utxos.UTXO
	var producedUTXOs [][]byte
	for _, tx := range block.Transactions {
		txID, err := tx.Hash()
		if err != nil {
			n.logger.Error("Error hashing transaction, walking back block commit ...", zap.Error(err))
			err = n.walkBackCommit(blockID, blockID, committedTXs, consumedUTXOs, producedUTXOs)
			if err != nil {
				n.logger.Error("Error walking back block commit", zap.Error(err))
				return false
			}
			return false
		}
		// add transactions to transaction store
		err = n.TransactionStore.Put(tx)
		if err != nil {
			n.logger.Error("Error adding transaction to transaction store, walking back block commit ...", zap.Error(err))
			err = n.walkBackCommit(blockID, blockID, committedTXs, consumedUTXOs, producedUTXOs)
			if err != nil {
				n.logger.Error("Error walking back block commit", zap.Error(err))
				return false
			}
			return false
		}
		committedTXs = append(committedTXs, txID)

		// remove consumed utxos from utxo store
		for _, input := range tx.Inputs {
			utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
			utxo, err := n.UtxoStore.Delete(utxoID)
			if err != nil {
				n.logger.Error("Error removing consumed utxo from utxo store, walking back block commit ...", zap.Error(err))
				err = n.walkBackCommit(blockID, blockID, committedTXs, consumedUTXOs, producedUTXOs)
				if err != nil {
					n.logger.Error("Error walking back block commit", zap.Error(err))
					return false
				}
				return false
			}
			consumedUTXOs[hex.EncodeToString(utxoID)] = utxo
		}
		// add produced utxos to utxo store
		for i, output := range tx.Outputs {
			utxoID := utxos.CreateUTXOID(txID, int64(i))
			utxo := utxos.NewUTXO(output.Amount, output.LockingScript)
			err = n.UtxoStore.Put(utxoID, utxo)
			if err != nil {
				n.logger.Error("Error adding produced utxo to utxo store, walking back block commit ...", zap.Error(err))
				err = n.walkBackCommit(blockID, blockID, committedTXs, consumedUTXOs, producedUTXOs)
				if err != nil {
					n.logger.Error("Error walking back block commit", zap.Error(err))
					return false
				}
				return false
			}
			producedUTXOs = append(producedUTXOs, utxoID)
		}
	}
}

func (n *Node) walkBackCommit(blockIDStore []byte, blockIDList []byte, txIDs [][]byte, consumedUTXOs map[string]*utxos.UTXO, producedUTXOs [][]byte) error {
	if blockIDStore != nil {
		err := n.BlockStore.Delete(blockIDStore)
		if err != nil {
			return err
		}
	}
	if blockIDList != nil {
		err := n.BlockList.Delete(blockIDList)
		if err != nil {
			return err
		}
	}
	for _, txID := range txIDs {
		err := n.TransactionStore.Delete(txID)
		if err != nil {
			return err
		}
	}
	for utxoIDStr, utxo := range consumedUTXOs {
		utxoID, err := hex.DecodeString(utxoIDStr)
		if err != nil {
			return err
		}
		err = n.UtxoStore.Put(utxoID, utxo)
		if err != nil {
			return err
		}
	}
	for _, utxoID := range producedUTXOs {
		_, err := n.UtxoStore.Delete(utxoID)
		if err != nil {
			return err
		}
	}
	return nil
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
