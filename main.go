package main

import (
	"log"
	"os"

	n "github.com/fahyjo/blockchain/node"
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

	node := n.NewNode(listenAddr, make(map[string]*n.Peer), 0, logger)
	log.Fatal(node.Start(peerAddrs))
}
