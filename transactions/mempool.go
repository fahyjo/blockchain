package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/fahyjo/blockchain/utxos"
)

// Mempool keeps track of all currently pending transactions
type Mempool struct {
	transactionLock sync.RWMutex
	transactions    map[string]*Transaction

	utxoLock sync.RWMutex
	utxos    map[string]bool
}

func NewMempool() *Mempool {
	return &Mempool{
		transactionLock: sync.RWMutex{},
		transactions:    make(map[string]*Transaction),
		utxoLock:        sync.RWMutex{},
		utxos:           make(map[string]bool),
	}
}

func (m *Mempool) GetTransaction(txID []byte) (*Transaction, error) {
	m.transactionLock.RLock()
	defer m.transactionLock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	tx, ok := m.transactions[txIDStr]
	if !ok {
		return nil, fmt.Errorf("mempool transactions cache miss: %s", txIDStr)
	}

	return tx, nil
}

func (m *Mempool) PutTransaction(txID []byte, tx *Transaction) {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	txIDStr := hex.EncodeToString(txID)
	m.transactions[txIDStr] = tx
}

func (m *Mempool) DeleteTransaction(txIDStr string) {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	delete(m.transactions, txIDStr)
}

func (m *Mempool) HasTransaction(txID []byte) bool {
	m.transactionLock.RLock()
	defer m.transactionLock.RUnlock()

	txIDStr := hex.EncodeToString(txID)
	_, ok := m.transactions[txIDStr]
	return ok
}

func (m *Mempool) Cleanse() {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	for txIDStr, tx := range m.transactions {
		for _, input := range tx.Inputs {
			utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
			ok := m.HasUTXO(utxoID)
			if !ok {
				m.DeleteTransaction(txIDStr)
			}
		}
	}
}

func (m *Mempool) TransactionsSize() int {
	return len(m.transactions)
}

func (m *Mempool) PutUTXO(utxoID []byte, b bool) {
	m.utxoLock.Lock()
	defer m.utxoLock.Unlock()

	utxoIDStr := hex.EncodeToString(utxoID)
	m.utxos[utxoIDStr] = b
}

func (m *Mempool) DeleteUTXO(utxoID []byte) {
	m.utxoLock.Lock()
	defer m.utxoLock.Unlock()

	utxoIDStr := hex.EncodeToString(utxoID)
	delete(m.utxos, utxoIDStr)
}

func (m *Mempool) HasUTXO(utxoID []byte) bool {
	m.utxoLock.RLock()
	defer m.utxoLock.RUnlock()

	utxoIDStr := hex.EncodeToString(utxoID)
	_, ok := m.utxos[utxoIDStr]
	return ok
}

func (m *Mempool) UTXOsSize() int {
	return len(m.utxos)
}
