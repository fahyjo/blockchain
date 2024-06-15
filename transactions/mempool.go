package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"
)

type Mempool struct {
	lock  sync.RWMutex
	cache map[string][]byte
}

func NewMempool() *Mempool {
	return &Mempool{
		lock:  sync.RWMutex{},
		cache: make(map[string][]byte),
	}
}

func (m *Mempool) Get(txID []byte) (*Transaction, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	b, ok := m.cache[txIDStr]
	if !ok {
		return nil, fmt.Errorf("mempool cache miss: %s", txIDStr)
	}

	tx, err := DecodeTransaction(b)
	if err != nil {
		return nil, err
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
	b, err := EncodeTransaction(tx)
	if err != nil {
		return err
	}
	m.cache[txIDStr] = b

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
