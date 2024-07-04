package blocks

import (
	"encoding/hex"
	"fmt"
	"sync"
)

type BlockList struct {
	lock sync.RWMutex
	list []string
}

func NewBlockList() *BlockList {
	return &BlockList{
		lock: sync.RWMutex{},
		list: make([]string, 0, 50),
	}
}

func (l *BlockList) Add(blockID []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	blockIDStr := hex.EncodeToString(blockID)
	l.list = append(l.list, blockIDStr)
	return nil
}

func (l *BlockList) Get(index int) (string, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if index >= l.Size() || index < 0 {
		return "", fmt.Errorf("error retrieving blockID from block list, index out of range: %d", index)
	}

	return l.list[index], nil
}

func (l *BlockList) Size() int {
	return len(l.list)
}

func (l *BlockList) Delete(blockID []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	blockIDStr := hex.EncodeToString(blockID)

	for i, s := range l.list {
		if s == blockIDStr {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error deleting block from block list: block %s not found", blockIDStr)
}
