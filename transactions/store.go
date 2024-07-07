package transactions

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type TransactionStore interface {
	Get([]byte) (*Transaction, error)
	Put([]byte, *Transaction) error
	Delete([]byte) error
}

type MemoryTransactionStore struct {
	db map[string]*Transaction
}

func NewMemoryTransactionStore() TransactionStore {
	return &MemoryTransactionStore{
		db: make(map[string]*Transaction),
	}
}

func (s *MemoryTransactionStore) Get(txID []byte) (*Transaction, error) {
	txIDStr := hex.EncodeToString(txID)
	tx, ok := s.db[txIDStr]
	if !ok {
		return nil, fmt.Errorf("transaction not found: %s", txIDStr)
	}
	return tx, nil
}

func (s *MemoryTransactionStore) Put(txID []byte, tx *Transaction) error {
	txIDStr := hex.EncodeToString(txID)
	s.db[txIDStr] = tx
	return nil
}

func (s *MemoryTransactionStore) Delete(txID []byte) error {
	txIDStr := hex.EncodeToString(txID)
	_, ok := s.db[txIDStr]
	if !ok {
		return fmt.Errorf("transaction not found not found: %s", txIDStr)
	}
	delete(s.db, txIDStr)
	return nil
}

type LevelsTransactionStore struct {
	db *leveldb.DB
}

func NewLevelsTransactionStore(path string) (TransactionStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelsTransactionStore{
		db: db,
	}, nil
}

func (s *LevelsTransactionStore) Get(txID []byte) (*Transaction, error) {
	b, err := s.db.Get(txID, nil)
	if err != nil {
		return nil, err
	}
	tx, err := DecodeTransaction(b)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (s *LevelsTransactionStore) Put(txID []byte, tx *Transaction) error {
	b, err := EncodeTransaction(tx)
	if err != nil {
		return err
	}

	err = s.db.Put(txID, b, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *LevelsTransactionStore) Delete(txID []byte) error {
	b, err := s.db.Has(txID, nil)
	if err != nil {
		return err
	}
	if !b {
		return fmt.Errorf("error deleting tx from tx store: tx %s not found", hex.EncodeToString(txID))
	}
	err = s.db.Delete(txID, nil)
	if err != nil {
		return err
	}
	return nil
}
