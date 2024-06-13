package transactions

import (
	"encoding/hex"
	"fmt"
)

type Mempool struct {
	txs map[string][]byte
}

func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string][]byte),
	}
}

func (m *Mempool) Get(txID []byte) (*Transaction, error) {
	b, ok := m.txs[hex.EncodeToString(txID)]
	if !ok {
		return nil, fmt.Errorf("transaction not found: %s", txID)
	}
	tx, err := DecodeTransaction(b)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (m *Mempool) Put(tx *Transaction) error {
	hash, err := tx.Hash()
	if err != nil {
		return err
	}
	txID := hex.EncodeToString(hash)

	b, err := EncodeTransaction(tx)
	if err != nil {
		return err
	}

	m.txs[txID] = b
	return nil
}
