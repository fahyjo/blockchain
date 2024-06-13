package blocks

import "github.com/syndtr/goleveldb/leveldb"

const (
	blockStorePath = "blocks/db/blocks.db"
)

type BlockStore interface {
	Get([]byte) (*Block, error)
	Put(*Block) error
}

type LevelsBlockStore struct {
	db *leveldb.DB
}

func NewLevelsBlockStore(db *leveldb.DB) (*LevelsBlockStore, error) {
	db, err := leveldb.OpenFile(blockStorePath, nil)
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
