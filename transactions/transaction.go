package transactions

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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

func EncodeTransaction(tx *Transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&tx)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeTransaction(b []byte) (*Transaction, error) {
	var tx Transaction
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (t *Transaction) Hash() ([]byte, error) {
	var unlockingScripts []*UnlockingScript
	for _, input := range t.Inputs {
		unlockingScripts = append(unlockingScripts, input.UnlockingScript)
		input.UnlockingScript = nil
	}

	b, err := EncodeTransaction(t)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)

	for i, input := range t.Inputs {
		input.UnlockingScript = unlockingScripts[i]
	}

	return hash[:], nil
}
