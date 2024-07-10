package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/fahyjo/blockchain/utxos"
)

// Mempool keeps track of all currently pending transactions and the utxos that they claim
type Mempool struct {
	transactionLock sync.RWMutex            // transactionLock read/write mutex that ensures Mempool transactions is concurrent-safe
	transactions    map[string]*Transaction // transactions maps transaction ids to transactions for all currently pending transactions

	utxoLock sync.RWMutex    // utxoLock read/write mutex that ensures Mempool utxos is concurrent-safe
	utxos    map[string]bool // utxos maps utxo ids to utxos for all utxos claimed by transactions in the Mempool
}

// NewMempool creates a new Mempool
func NewMempool() *Mempool {
	return &Mempool{
		transactionLock: sync.RWMutex{},
		transactions:    make(map[string]*Transaction),
		utxoLock:        sync.RWMutex{},
		utxos:           make(map[string]bool),
	}
}

// GetTransaction retrieves the transaction with the given transaction id, returns an error if the transaction is not present
func (m *Mempool) GetTransaction(txIDStr string) (*Transaction, error) {
	m.transactionLock.RLock()
	defer m.transactionLock.RUnlock()

	tx, ok := m.transactions[txIDStr]
	if !ok {
		return nil, fmt.Errorf("error getting from transactions mempool: transaction %s not found", txIDStr)
	}

	return tx, nil
}

// PutTransaction maps the given transaction id to the given transaction
func (m *Mempool) PutTransaction(txIDStr string, tx *Transaction) {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	m.transactions[txIDStr] = tx
}

// DeleteTransaction removes the transaction with the given transaction id, returns an error if the transaction is not present
func (m *Mempool) DeleteTransaction(txIDStr string) error {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	ok := m.HasTransaction(txIDStr)
	if !ok {
		return fmt.Errorf("error deleting from transactions mempool: transaction %s not found", txIDStr)
	}

	delete(m.transactions, txIDStr)
	return nil
}

// HasTransaction returns true if the transaction with the given transaction id is in the Mempool, returns false otherwise
func (m *Mempool) HasTransaction(txIDStr string) bool {
	m.transactionLock.RLock()
	defer m.transactionLock.RUnlock()

	_, ok := m.transactions[txIDStr]
	return ok
}

// Cleanse removes all transactions from the Mempool if the utxos they claim have been spent since the transactions were added to the Mempool
func (m *Mempool) Cleanse() error {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	for txIDStr, tx := range m.transactions {
		for _, input := range tx.Inputs {
			utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
			ok := m.HasUTXO(utxoID)
			if !ok {
				err := m.DeleteTransaction(txIDStr)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// TransactionsSize returns the number of transactions in the Mempool
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
