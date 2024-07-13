package node

import (
	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/peers"
	"github.com/fahyjo/blockchain/transactions"
)

type Cache struct {
	PeerCache        *peers.PeerCache
	BlockCache       *blocks.BlockCache
	TransactionCache *transactions.TransactionCache
}

func NewCache(pc *peers.PeerCache, bc *blocks.BlockCache, tc *transactions.TransactionCache) *Cache {
	return &Cache{
		PeerCache:        pc,
		BlockCache:       bc,
		TransactionCache: tc,
	}
}
