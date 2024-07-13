package main

import (
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/fahyjo/blockchain/blocks"
	c "github.com/fahyjo/blockchain/config"
	"github.com/fahyjo/blockchain/consensus"
	"github.com/fahyjo/blockchain/crypto"
	n "github.com/fahyjo/blockchain/node"
	"github.com/fahyjo/blockchain/peers"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
	"go.uber.org/zap"
)

func main() {
	// make logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// open json config file
	configFile, err := os.Open("config/config.json")
	if err != nil {
		logger.Fatal("Failed to read config.json", zap.Error(err))
	}
	defer configFile.Close()

	// parse json config file
	var config map[string]c.Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatal("Failed to parse config.json", zap.Error(err))
	}

	nodeIDStr := os.Args[1]

	// network config
	networkConfig := config[nodeIDStr].Network
	listenAddr := networkConfig.ListenAddr
	peerAddrs := networkConfig.Peers

	// crypto config
	cryptoConfig := config[nodeIDStr].Crypto
	pubKeyStr := cryptoConfig.PublicKey
	pubKeyBytes, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		logger.Fatal("Failed to decode public key", zap.Error(err))
	}
	pubKey := crypto.NewPublicKey(pubKeyBytes)
	privKeyStr := cryptoConfig.PrivateKey
	privKeyBytes, err := hex.DecodeString(privKeyStr)
	if err != nil {
		logger.Fatal("Failed to decode private key", zap.Error(err))
	}
	privKey := crypto.NewPrivateKey(privKeyBytes)
	keys := crypto.NewKeys(privKey, pubKey)

	// consensus config
	consensusConfig := config[nodeIDStr].Consensus
	amValidator := consensusConfig.AmValidator
	validators := consensusConfig.Validators
	round := consensus.NewRound(validators[0])
	consensus := consensus.NewConsensus(amValidator, validators, 0, round)

	// make caches
	var (
		peerCache         = peers.NewPeerCache()
		blockCache        = blocks.NewBlockCache()
		transactionsCache = transactions.NewTransactionCache()
		cache             = n.NewCache(peerCache, blockCache, transactionsCache)
	)

	// make stores
	var (
		blockStore       = blocks.NewMemoryBlockStore()
		transactionStore = transactions.NewMemoryTransactionStore()
		utxoStore        = utxos.NewMemoryUTXOStore()
		store            = n.NewStore(blockStore, transactionStore, utxoStore)
	)

	// make blocklist and mempool
	blockList := blocks.NewBlockList()
	mempool := transactions.NewMempool()

	// make and start node
	node := n.NewNode(listenAddr, 0, keys, consensus, cache, store, blockList, mempool, logger)
	err = node.Start(peerAddrs)
	if err != nil {
		logger.Fatal("Failed to start node", zap.Error(err))
	}
}

/*
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
*/
