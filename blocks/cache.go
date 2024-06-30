package blocks

import (
	"encoding/hex"
	"fmt"
	"sync"
)

type BlockCache struct {
	lock  sync.RWMutex
	cache map[string]bool
}

func NewBlockCache() *BlockCache {
	return &BlockCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]bool),
	}
}

func (c *BlockCache) Get(blockID []byte) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	blockIDStr := hex.EncodeToString(blockID)
	b, ok := c.cache[blockIDStr]
	if !ok {
		return false, fmt.Errorf("block cache miss: %s", blockIDStr)
	}
	return b, nil
}

func (c *BlockCache) Put(blockID []byte, b bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	blockIDStr := hex.EncodeToString(blockID)
	c.cache[blockIDStr] = b
}

func (c *BlockCache) Has(blockID []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	blockIDStr := hex.EncodeToString(blockID)
	_, ok := c.cache[blockIDStr]
	return ok
}
