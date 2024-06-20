package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"
)

// TransactionCache keeps track of all the transactions we have seen
type TransactionCache struct {
	lock  sync.RWMutex
	cache map[string]bool
}

func NewTransactionCache() *TransactionCache {
	return &TransactionCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]bool),
	}
}

func (c *TransactionCache) Get(txID []byte) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	b, ok := c.cache[txIDStr]
	if !ok {
		return false, fmt.Errorf("transaction cache miss: %s", txIDStr)
	}
	return b, nil
}

func (c *TransactionCache) Put(txID []byte, b bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	txIDStr := hex.EncodeToString(txID)
	c.cache[txIDStr] = b
}

func (c *TransactionCache) Has(txID []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	_, ok := c.cache[txIDStr]
	return ok
}

func (c *TransactionCache) Size() int {
	return len(c.cache)
}
