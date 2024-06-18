package blocks

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type BlockStore interface {
	Get([]byte) (*Block, error)
	Put(*Block) error
}

type MemoryBlockStore struct {
	db map[string]*Block
}

func NewMemoryBlockStore() BlockStore {
	return &MemoryBlockStore{
		db: make(map[string]*Block),
	}
}

func (s *MemoryBlockStore) Get(blockID []byte) (*Block, error) {
	blockIDStr := hex.EncodeToString(blockID)
	block, ok := s.db[blockIDStr]
	if !ok {
		return nil, fmt.Errorf("block not found: %s", blockIDStr)
	}
	return block, nil
}

func (s *MemoryBlockStore) Put(block *Block) error {
	blockID, err := block.Hash()
	if err != nil {
		return err
	}

	blockIDStr := hex.EncodeToString(blockID)
	s.db[blockIDStr] = block
	return nil
}

type LevelsBlockStore struct {
	db *leveldb.DB
}

func NewLevelsBlockStore(path string) (BlockStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelsBlockStore{
		db: db,
	}, nil
}

func (s *LevelsBlockStore) Get(blockID []byte) (*Block, error) {
	b, err := s.db.Get(blockID, nil)
	block, err := DecodeBlock(b)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (s *LevelsBlockStore) Put(b *Block) error {
	blockID, err := b.Hash()
	if err != nil {
		return err
	}

	by, err := EncodeBlock(b)
	if err != nil {
		return err
	}

	err = s.db.Put(blockID, by, nil)
	if err != nil {
		return err
	}

	return nil
}
