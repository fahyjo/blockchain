package utxos

import (
	"encoding/hex"
	"fmt"
	"sync"
)

// UTXOCache keeps track of all utxos claimed by transactions in mempool
type UTXOCache struct {
	lock  sync.RWMutex
	cache map[string]bool
}

func NewUTXOCache() *UTXOCache {
	return &UTXOCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]bool),
	}
}

func (c *UTXOCache) Get(utxoID []byte) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	utxoIDStr := hex.EncodeToString(utxoID)
	b, ok := c.cache[utxoIDStr]
	if !ok {
		return false, fmt.Errorf("utxo cache miss: %s", utxoIDStr)
	}
	return b, nil
}

func (c *UTXOCache) Put(utxoID []byte, b bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	utxoIDStr := hex.EncodeToString(utxoID)
	c.cache[utxoIDStr] = b
}

func (c *UTXOCache) Has(utxoID []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	txIDStr := hex.EncodeToString(utxoID)
	_, ok := c.cache[txIDStr]
	return ok
}

func (c *UTXOCache) Size() int {
	return len(c.cache)
}
