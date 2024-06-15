package main

import (
	"os"

	n "github.com/fahyjo/blockchain/node"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
	"go.uber.org/zap"
)

var (
	listenAddr string
	peerAddrs  []string
)

func init() {
	listenAddr = os.Args[1]
	for _, peerAddr := range os.Args[2:] {
		peerAddrs = append(peerAddrs, peerAddr)
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	mempool := transactions.NewMempool()

	utxoStore, err := utxos.NewLevelsUTXOStore(listenAddr)
	if err != nil {
		logger.Error(err.Error())
	}

	node := n.NewNode(listenAddr, make(map[string]*n.Peer), 0, make(map[string]bool), mempool, utxoStore, logger)
	err = node.Start(peerAddrs)
	if err != nil {
		logger.Error(err.Error())
	}
}
