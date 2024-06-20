package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"
)

// Mempool keeps track of all currently pending transactions
type Mempool struct {
	lock  sync.RWMutex
	cache map[string]*Transaction
}

func NewMempool() *Mempool {
	return &Mempool{
		lock:  sync.RWMutex{},
		cache: make(map[string]*Transaction),
	}
}

func (m *Mempool) Get(txID []byte) (*Transaction, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	tx, ok := m.cache[txIDStr]
	if !ok {
		return nil, fmt.Errorf("mempool cache miss: %s", txIDStr)
	}

	return tx, nil
}

func (m *Mempool) Put(tx *Transaction) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	txID, err := tx.Hash()
	if err != nil {
		return err
	}
	txIDStr := hex.EncodeToString(txID)

	m.cache[txIDStr] = tx

	return nil
}

func (m *Mempool) Has(txID []byte) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	_, ok := m.cache[txIDStr]
	return ok
}

func (m *Mempool) Size() int {
	return len(m.cache)
}
