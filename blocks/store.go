package blocks

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// BlockStore stores all blocks committed to the blockchain in a key-value store
type BlockStore interface {
	Get([]byte) (*Block, error) // Get retrieves the Block with the given block id, returns an error if the Block with the given block id is not found
	Put([]byte, *Block) error   // Put maps the given block id to the given Block
	Delete([]byte) error        // Delete removes the Block with the given block id, returns an error if the Block with the given block id is not found
}

// MemoryBlockStore is an in-memory implementation of BlockStore
type MemoryBlockStore struct {
	db map[string]*Block // db maps the block id (encoded as a string) to the block
}

// NewMemoryBlockStore creates a new MemoryBlockStore struct
func NewMemoryBlockStore() BlockStore {
	return &MemoryBlockStore{
		db: make(map[string]*Block),
	}
}

func (s *MemoryBlockStore) Get(blockID []byte) (*Block, error) {
	blockIDStr := hex.EncodeToString(blockID)
	block, ok := s.db[blockIDStr]
	if !ok {
		return nil, fmt.Errorf("error getting from block store: block %s not found", blockIDStr)
	}
	return block, nil
}

func (s *MemoryBlockStore) Put(blockID []byte, block *Block) error {
	blockIDStr := hex.EncodeToString(blockID)
	s.db[blockIDStr] = block
	return nil
}

func (s *MemoryBlockStore) Delete(blockID []byte) error {
	blockIDStr := hex.EncodeToString(blockID)
	_, ok := s.db[blockIDStr]
	if !ok {
		return fmt.Errorf("error deleting from block store: block %s not found", blockIDStr)
	}
	delete(s.db, blockIDStr)
	return nil
}

// LevelsBlockStore uses LevelDB to store blocks
type LevelsBlockStore struct {
	db *leveldb.DB // db client to interact with LevelDB store
}

// NewLevelsBlockStore creates a new LevelsBlockStore struct at the given path
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
	if err != nil {
		return nil, err
	}
	block, err := DecodeBlock(b)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (s *LevelsBlockStore) Put(blockID []byte, b *Block) error {
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

func (s *LevelsBlockStore) Delete(blockID []byte) error {
	ok, err := s.db.Has(blockID, nil)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("error deleting block from block store: block %s not found", hex.EncodeToString(blockID))
	}

	err = s.db.Delete(blockID, nil)
	if err != nil {
		return err
	}

	return nil
}
