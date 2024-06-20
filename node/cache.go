package node

import (
	"github.com/fahyjo/blockchain/peer"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
)

type Cache struct {
	PeerCache        *peer.PeerCache
	TransactionCache *transactions.TransactionCache
	Mempool          *transactions.Mempool
	UTXOCache        *utxos.UTXOCache
}

func NewCache(pc *peer.PeerCache, tc *transactions.TransactionCache, mp *transactions.Mempool, uc *utxos.UTXOCache) *Cache {
	return &Cache{
		PeerCache:        pc,
		TransactionCache: tc,
		Mempool:          mp,
		UTXOCache:        uc,
	}
}
