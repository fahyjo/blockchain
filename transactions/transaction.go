package transactions

import (
	"crypto/sha256"
)

type Transaction struct {
	Inputs  []*Input
	Outputs []*Output
}

func NewTransaction(inputs []*Input, outputs []*Output) *Transaction {
	return &Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func (tx *Transaction) Hash() ([]byte, error) {
	var unlockingScripts []*UnlockingScript
	for _, input := range tx.Inputs {
		unlockingScripts = append(unlockingScripts, input.UnlockingScript)
		input.UnlockingScript = nil
	}

	b, err := EncodeTransaction(tx)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	for i, input := range tx.Inputs {
		input.UnlockingScript = unlockingScripts[i]
	}

	return hash[:], nil
}
