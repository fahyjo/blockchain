package utxos

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// UTXOStore stores UTXOs in a key-value store
type UTXOStore interface {
	Get([]byte) (*UTXO, error)    // Get retrieves the UTXO with the given utxo ID, returns error if the UTXO is not found
	Put([]byte, *UTXO) error      // Put maps the given utxo ID to the given UTXO
	Delete([]byte) (*UTXO, error) // Delete removes the UTXO with the given utxo ID, returns the removed UTXO, and returns an error if the UTXO is not found
}

// MemoryUTXOStore is an in-memory UTXOStore implementation
type MemoryUTXOStore struct {
	db map[string]*UTXO // db maps the utxo ID (encoded as a string) to the UTXO
}

// NewMemoryUTXOStore creates a new MemoryUTXOStore
func NewMemoryUTXOStore() UTXOStore {
	return &MemoryUTXOStore{
		db: make(map[string]*UTXO),
	}
}

func (s *MemoryUTXOStore) Get(utxoID []byte) (*UTXO, error) {
	utxoIDStr := hex.EncodeToString(utxoID)
	utxo, ok := s.db[utxoIDStr]
	if !ok {
		return nil, fmt.Errorf("error getting from utxo store: utxo %s not found", utxoIDStr)
	}
	return utxo, nil
}

func (s *MemoryUTXOStore) Put(utxoID []byte, utxo *UTXO) error {
	utxoIDStr := hex.EncodeToString(utxoID)
	s.db[utxoIDStr] = utxo
	return nil
}

func (s *MemoryUTXOStore) Delete(utxoID []byte) (*UTXO, error) {
	utxoIDStr := hex.EncodeToString(utxoID)
	utxo, ok := s.db[utxoIDStr]
	if !ok {
		return nil, fmt.Errorf("error deleting from utxo store: utxo %s not found", hex.EncodeToString(utxoID))
	}
	delete(s.db, utxoIDStr)
	return utxo, nil
}

// LevelsUTXOStore uses LevelDB to store utxos
type LevelsUTXOStore struct {
	db *leveldb.DB
}

// NewLevelsUTXOStore creates a new LevelsUTXOStore at the given path
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

func (s *LevelsUTXOStore) Delete(utxoID []byte) (*UTXO, error) {
	utxo, err := s.Get(utxoID)
	if err != nil {
		return nil, fmt.Errorf("error deleting from utxo store: utxo %s not found", hex.EncodeToString(utxoID))
	}
	err = s.db.Delete(utxoID, nil)
	if err != nil {
		return nil, err
	}
	return utxo, err
}
