package transactions

import (
	"crypto/sha256"
)

// Transaction consumes utxos and produces new utxos
type Transaction struct {
	Inputs  []*Input  // Inputs specifies the utxos to be spent
	Outputs []*Output // Outputs produces new utxos and specifies their owners
}

// NewTransaction creates a Transaction struct
func NewTransaction(inputs []*Input, outputs []*Output) *Transaction {
	return &Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}
}

// Hash hashes the given Transaction
// To hash a Transaction we set the Input unlocking scripts to nil, hash the transaction, and then restore the Input unlocking scripts
func (tx *Transaction) Hash() ([]byte, error) {
	var unlockingScripts []*UnlockingScript
	for i, input := range tx.Inputs {
		unlockingScripts = append(unlockingScripts, input.UnlockingScript)
		tx.Inputs[i] = nil
	}

	b, err := EncodeTransaction(tx)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	for i, _ := range tx.Inputs {
		tx.Inputs[i].UnlockingScript = unlockingScripts[i]
	}

	return hash[:], nil
}
