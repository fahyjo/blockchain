package blocks

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"
)

type BlockList struct {
	lock sync.RWMutex
	list [][]byte
}

func NewBlockList() *BlockList {
	return &BlockList{
		lock: sync.RWMutex{},
		list: make([][]byte, 0, 50),
	}
}

func (l *BlockList) Add(blockID []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = append(l.list, blockID)
	return nil
}

func (l *BlockList) Get(index int) ([]byte, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if index >= l.Size() || index < 0 {
		return nil, fmt.Errorf("error retrieving blockID from block list, index out of range: %d", index)
	}

	return l.list[index], nil
}

func (l *BlockList) Size() int {
	return len(l.list)
}

func (l *BlockList) Delete(blockID []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	for i, s := range l.list {
		if bytes.Equal(s, blockID) {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error deleting block from block list: block %s not found", hex.EncodeToString(blockID))
}
