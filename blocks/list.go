package blocks

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"
)

// BlockList holds a list of the block id of each Block in the blockchain
type BlockList struct {
	lock sync.RWMutex // lock read/write mutex that ensures BlockList is concurrent-safe
	list [][]byte     // lists holds the id of each Block in the blockchain
}

// NewBlockList creates a new BlockList struct
func NewBlockList() *BlockList {
	return &BlockList{
		lock: sync.RWMutex{},
		list: make([][]byte, 0, 50),
	}
}

// Add adds the given block id to the BlockList
func (l *BlockList) Add(blockID []byte) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = append(l.list, blockID)
}

// Get retrieves the id of the Block at the given index in the blockchain, returns an error if the given index is out of bounds
func (l *BlockList) Get(index int) ([]byte, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if index >= l.Size() || index < 0 {
		return nil, fmt.Errorf("error retrieving blockID from block list, index out of range: %d", index)
	}

	return l.list[index], nil
}

// Size returns the number of blocks in the blockchain
func (l *BlockList) Size() int {
	return len(l.list)
}

// Delete removes the given block id from the BlockLIst
func (l *BlockList) Delete(blockID []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	for i, s := range l.list {
		if bytes.Equal(s, blockID) {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error deleting blockID from block list: block id %s not found", hex.EncodeToString(blockID))
}

// Dump returns the number of blocks in the blockchain and the id of each Block in the blockchain
func (l *BlockList) Dump() (int, [][]byte, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	size := len(l.list)
	var blockIDs [][]byte
	for _, blockID := range l.list {
		blockIDs = append(blockIDs, blockID)
	}
	return size, blockIDs, nil
}
