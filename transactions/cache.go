package transactions

import (
	"sync"
)

// TransactionCache keeps track of all the transaction ids we have seen
type TransactionCache struct {
	lock  sync.RWMutex    // lock read/write mutex that ensures TransactionCache is concurrent-safe
	cache map[string]bool // cache holds all the transaction ids we have seen
}

// NewTransactionCache creates a new TransactionCache struct
func NewTransactionCache() *TransactionCache {
	return &TransactionCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]bool),
	}
}

// Put adds the given transaction id to the TransactionCache
func (c *TransactionCache) Put(txIDStr string, b bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[txIDStr] = b
}

// Has returns true if the given transaction id is in the TransactionCache, returns false otherwise
func (c *TransactionCache) Has(txIDStr string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.cache[txIDStr]
	return ok
}

// Size returns the number of Transactions in the TransactionCache
func (c *TransactionCache) Size() int {
	return len(c.cache)
}
