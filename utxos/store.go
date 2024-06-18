package utxos

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type UTXOStore interface {
	Get([]byte) (*UTXO, error)
	Put([]byte, *UTXO) error
}

type MemoryUTXOStore struct {
	db map[string]*UTXO
}

func NewMemoryUTXOStore() UTXOStore {
	return &MemoryUTXOStore{
		db: make(map[string]*UTXO),
	}
}

func (s *MemoryUTXOStore) Get(utxoID []byte) (*UTXO, error) {
	utxoIDStr := hex.EncodeToString(utxoID)
	utxo, ok := s.db[utxoIDStr]
	if !ok {
		return nil, fmt.Errorf("utxo not found: %s", utxoIDStr)
	}
	return utxo, nil
}

func (s *MemoryUTXOStore) Put(utxoID []byte, utxo *UTXO) error {
	utxoIDStr := hex.EncodeToString(utxoID)
	s.db[utxoIDStr] = utxo
	return nil
}

type LevelsUTXOStore struct {
	db *leveldb.DB
}

func NewLevelsUTXOStore(path string) (UTXOStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelsUTXOStore{
		db: db,
	}, nil
}

func (s *LevelsUTXOStore) Get(utxoID []byte) (*UTXO, error) {
	b, err := s.db.Get(utxoID, nil)
	if err != nil {
		return nil, err
	}
	utxo, err := DecodeUTXO(b)
	if err != nil {
		return nil, err
	}
	return utxo, nil
}

func (s *LevelsUTXOStore) Put(utxoID []byte, utxo *UTXO) error {
	b, err := EncodeUTXO(utxo)
	if err != nil {
		return err
	}

	err = s.db.Put(utxoID, b, nil)
	if err != nil {
		return err
	}

	return nil
}
