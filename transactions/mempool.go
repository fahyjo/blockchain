package transactions

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/fahyjo/blockchain/utxos"
)

// Mempool keeps track of all the currently pending transactions and the utxos that they claim
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

// Cleanse removes transactions from the Mempool if the utxos they claim have been spent
func (m *Mempool) Cleanse() error {
	m.transactionLock.Lock()
	defer m.transactionLock.Unlock()

	for txIDStr, tx := range m.transactions {
		for _, input := range tx.Inputs {
			utxoID := utxos.CreateUTXOID(input.TxID, input.UTXOIndex)
			utxoIDStr := hex.EncodeToString(utxoID)
			ok := m.HasUTXO(utxoIDStr)
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

// TransactionsSize returns the number of currently pending transactions
func (m *Mempool) TransactionsSize() int {
	return len(m.transactions)
}

// PutUTXO adds the utxo id to the utxos Mempool
func (m *Mempool) PutUTXO(utxoIDStr string, b bool) {
	m.utxoLock.Lock()
	defer m.utxoLock.Unlock()

	m.utxos[utxoIDStr] = b
}

// DeleteUTXO removes the given utxo id from the utxos Mempool, returns an error if the given utxo id is not found
func (m *Mempool) DeleteUTXO(utxoIDStr string) error {
	m.utxoLock.Lock()
	defer m.utxoLock.Unlock()

	ok := m.HasUTXO(utxoIDStr)
	if !ok {
		return fmt.Errorf("error deleting from utxos mempool: utxo %s not found", utxoIDStr)
	}
	delete(m.utxos, utxoIDStr)
	return nil
}

// HasUTXO returns true if the given utxo id is in the utxos Mempool, returns false otherwise
func (m *Mempool) HasUTXO(utxoIDStr string) bool {
	m.utxoLock.RLock()
	defer m.utxoLock.RUnlock()

	_, ok := m.utxos[utxoIDStr]
	return ok
}

// UTXOsSize returns the number of utxos claimed by currently pending transactions
func (m *Mempool) UTXOsSize() int {
	return len(m.utxos)
}
