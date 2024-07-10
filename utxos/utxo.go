package utxos

import (
	"github.com/fahyjo/blockchain/transactions"
)

// UTXO represents an unspent transaction output
type UTXO struct {
	Amount        int64                       // Amount specifies number of tokens this UTXO is worth
	LockingScript *transactions.LockingScript // LockingScript specifies the conditions under which this utxo can be spent
}

// NewUTXO creates a new UTXO
func NewUTXO(amount int64, lockingScript *transactions.LockingScript) *UTXO {
	return &UTXO{
		Amount:        amount,
		LockingScript: lockingScript,
	}
}
