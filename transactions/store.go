package transactions

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// TransactionStore stores all committed transactions in a key-value store
type TransactionStore interface {
	Get([]byte) (*Transaction, error) // Get retrieves the Transaction with the given transaction id, returns error if the Transactions is not found
	Put([]byte, *Transaction) error   // Put maps the given transaction id to the given Transaction
	Delete([]byte) error              // Delete removes the Transaction with the given transaction id, returns an error if the Transaction is not found
}

// MemoryTransactionStore is an in-memory TransactionStore implementation
type MemoryTransactionStore struct {
	db map[string]*Transaction //db maps the transaction id (encoded as a string) to the transaction
}

// NewMemoryTransactionStore creates a new MemoryTransactionStore
func NewMemoryTransactionStore() TransactionStore {
	return &MemoryTransactionStore{
		db: make(map[string]*Transaction),
	}
}

func (s *MemoryTransactionStore) Get(txID []byte) (*Transaction, error) {
	txIDStr := hex.EncodeToString(txID)
	tx, ok := s.db[txIDStr]
	if !ok {
		return nil, fmt.Errorf("error getting from transaction store: transaction %s not found", txIDStr)
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
		return fmt.Errorf("error deleting from transaction store: transaction %s not found", txIDStr)
	}
	delete(s.db, txIDStr)
	return nil
}

// LevelsTransactionStore uses LevelDB to store transactions
type LevelsTransactionStore struct {
	db *leveldb.DB
}

// NewLevelsTransactionStore creates a new LevelsTransactionStore at the given path
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
	ok, err := s.db.Has(txID, nil)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("error deleting from transaction store: transaction %s not found", hex.EncodeToString(txID))
	}
	err = s.db.Delete(txID, nil)
	if err != nil {
		return err
	}
	return nil
}
