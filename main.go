package main

import (
	"os"

	"github.com/fahyjo/blockchain/blocks"
	n "github.com/fahyjo/blockchain/node"
	"github.com/fahyjo/blockchain/peer"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
	"go.uber.org/zap"
)

func main() {
	listenAddr := os.Args[1]
	var peerAddrs []string
	for _, peerAddr := range os.Args[2:] {
		peerAddrs = append(peerAddrs, peerAddr)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var (
		peerCache         = peer.NewPeerCache()
		transactionsCache = transactions.NewTransactionCache()
		mempool           = transactions.NewMempool()
		cache             = n.NewCache(peerCache, transactionsCache, mempool)
	)

	blockStore, err := blocks.NewLevelsBlockStore(listenAddr)
	if err != nil {
		logger.Fatal("Failed to initialize block store", zap.Error(err))
	}
	transactionStore, err := transactions.NewLevelsTransactionStore(listenAddr)
	if err != nil {
		logger.Fatal("Failed to initialize transaction store", zap.Error(err))
	}
	utxoStore, err := utxos.NewLevelsUTXOStore(listenAddr)
	if err != nil {
		logger.Fatal("Failed to initialize utxo store", zap.Error(err))
	}
	store := n.NewStore(blockStore, transactionStore, utxoStore)

	node := n.NewNode(listenAddr, 0, cache, store, logger)
	err = node.Start(peerAddrs)
	if err != nil {
		logger.Fatal("Failed to start node", zap.Error(err))
	}
}
