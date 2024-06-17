package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/fahyjo/blockchain/blocks"
	c "github.com/fahyjo/blockchain/config"
	"github.com/fahyjo/blockchain/crypto"
	n "github.com/fahyjo/blockchain/node"
	"github.com/fahyjo/blockchain/peer"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	configFile, err := os.Open("config.json")
	if err != nil {
		logger.Fatal("Failed to read config.json", zap.Error(err))
	}
	defer configFile.Close()

	var config []c.Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatal("Failed to parse config.json", zap.Error(err))
	}

	nodeIDStr := os.Args[1]
	nodeID, err := strconv.Atoi(nodeIDStr)
	if err != nil {
		logger.Fatal("Failed to parse nodeID", zap.Error(err))
	}

	listenAddr := config[nodeID].ListenAddr
	peerAddrs := config[nodeID].Peers

	privKey, err := crypto.NewPrivateKey()
	if err != nil {
		logger.Fatal("Failed to generate private key", zap.Error(err))
	}
	pubKey := privKey.PublicKey()
	keys := crypto.NewKeys(privKey, pubKey)

	var (
		peerCache         = peer.NewPeerCache()
		transactionsCache = transactions.NewTransactionCache()
		mempool           = transactions.NewMempool()
		cache             = n.NewCache(peerCache, transactionsCache, mempool)
	)

	blockPath := config[nodeID].Block
	blockStore, err := blocks.NewLevelsBlockStore(blockPath)
	if err != nil {
		logger.Fatal("Failed to initialize block store", zap.Error(err))
	}
	transactionPath := config[nodeID].Transaction
	transactionStore, err := transactions.NewLevelsTransactionStore(transactionPath)
	if err != nil {
		logger.Fatal("Failed to initialize transaction store", zap.Error(err))
	}
	utxoPath := config[nodeID].Utxo
	utxoStore, err := utxos.NewLevelsUTXOStore(utxoPath)
	if err != nil {
		logger.Fatal("Failed to initialize utxo store", zap.Error(err))
	}
	store := n.NewStore(blockStore, transactionStore, utxoStore)

	node := n.NewNode(listenAddr, 0, keys, cache, store, logger)
	err = node.Start(peerAddrs)
	if err != nil {
		logger.Fatal("Failed to start node", zap.Error(err))
	}
}
