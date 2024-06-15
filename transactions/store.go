package transactions

import (
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	transactionStorePath = "transactions/db/transactions.db"
)

type TransactionStore interface {
	Get([]byte) (*Transaction, error)
	Put(*Transaction) error
}

type LevelsTransactionStore struct {
	db *leveldb.DB
}

func NewLevelsTransactionStore() (TransactionStore, error) {
	db, err := leveldb.OpenFile(transactionStorePath, nil)
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

func (s *LevelsTransactionStore) Put(tx *Transaction) error {
	hash, err := tx.Hash()
	if err != nil {
		return err
	}

	b, err := EncodeTransaction(tx)
	if err != nil {
		return err
	}

	err = s.db.Put(hash, b, nil)
	if err != nil {
		return err
	}

	return nil
}
