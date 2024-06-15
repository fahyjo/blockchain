package node

import (
	"github.com/fahyjo/blockchain/peer"
	"github.com/fahyjo/blockchain/transactions"
)

type Cache struct {
	PeerCache        *peer.PeerCache
	TransactionCache *transactions.TransactionCache
	Mempool          *transactions.Mempool
}

func NewCache(pc *peer.PeerCache, tc *transactions.TransactionCache, mp *transactions.Mempool) *Cache {
	return &Cache{
		PeerCache:        pc,
		TransactionCache: tc,
		Mempool:          mp,
	}
}
