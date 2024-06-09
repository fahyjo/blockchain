package transactions

import "sync"

type Mempool struct {
	lock sync.RWMutex
	txs  map[string][]byte
}

func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string][]byte),
	}
}
