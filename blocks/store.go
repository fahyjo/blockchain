package blocks

import "github.com/syndtr/goleveldb/leveldb"

type BlockStore interface {
	Get([]byte) (*Block, error)
	Put(*Block) error
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
	hash, err := b.Hash()
	if err != nil {
		return err
	}

	by, err := EncodeBlock(b)
	if err != nil {
		return err
	}

	err = s.db.Put(hash, by, nil)
	if err != nil {
		return err
	}

	return nil
}
