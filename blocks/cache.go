package blocks

import (
	"fmt"
	"sync"
)

// BlockCache keeps track of all block ids we have seen
type BlockCache struct {
	lock  sync.RWMutex    // lock read/write mutex that ensures BlockCache is concurrent-safe
	cache map[string]bool // cache holds all the block ids we have seen
}

// NewBlockCache creates a new BlockCache struct
func NewBlockCache() *BlockCache {
	return &BlockCache{
		lock:  sync.RWMutex{},
		cache: make(map[string]bool),
	}
}

// Get retrieves the Block with the given block id, returns an error if the Block with the given block id is not found
func (c *BlockCache) Get(blockIDStr string) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	b, ok := c.cache[blockIDStr]
	if !ok {
		return false, fmt.Errorf("error getting from block cache: block %s not found", blockIDStr)
	}
	return b, nil
}

// Put adds the given block id to the BlockCache
func (c *BlockCache) Put(blockIDStr string, b bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[blockIDStr] = b
}

// Has returns true if the given block id is in the BlockCache, returns false otherwise
func (c *BlockCache) Has(blockIDStr string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.cache[blockIDStr]
	return ok
}
