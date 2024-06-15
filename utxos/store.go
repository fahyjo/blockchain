package utxos

import "github.com/syndtr/goleveldb/leveldb"

const (
	utxoStorePath = "utxos/db/utxos"
)

type UTXOStore interface {
	Get([]byte) (*UTXO, error)
	Put([]byte, *UTXO) error
}

type LevelsUTXOStore struct {
	db *leveldb.DB
}

func NewLevelsUTXOStore(listenAddr string) (UTXOStore, error) {
	path := utxoStorePath + "/" + listenAddr
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
