package node

import (
	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/peers"
	"github.com/fahyjo/blockchain/transactions"
)

// Cache holds a node's PeerCache, BlockCache, and TransactionCache
type Cache struct {
	PeerCache        *peers.PeerCache               // PeerCache holds all known peers
	BlockCache       *blocks.BlockCache             // BlockCache holds the id of every block we have seen
	TransactionCache *transactions.TransactionCache // TransactionCache holds the id of every transaction we have seen
}

// NewCache creates a new Cache struct
func NewCache(pc *peers.PeerCache, bc *blocks.BlockCache, tc *transactions.TransactionCache) *Cache {
	return &Cache{
		PeerCache:        pc,
		BlockCache:       bc,
		TransactionCache: tc,
	}
}
