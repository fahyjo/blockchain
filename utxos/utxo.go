package utxos

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"strconv"

	"github.com/fahyjo/blockchain/transactions"
)

// UTXO represents an unspent transaction output
type UTXO struct {
	Amount        int64                       // Amount specifies number of tokens this utxo is worth
	LockingScript *transactions.LockingScript // LockingScript specifies the conditions under which this utxo can be spent
}

// NewUTXO creates a new UTXO
func NewUTXO(amount int64, lockingScript *transactions.LockingScript) *UTXO {
	return &UTXO{
		Amount:        amount,
		LockingScript: lockingScript,
	}
}

// EncodeUTXO encodes a UTXO into a byte slice
func EncodeUTXO(utxo *UTXO) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(utxo)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeUTXO decodes a byte slice into a UTXO
func DecodeUTXO(b []byte) (*UTXO, error) {
	var utxo UTXO
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&utxo)
	if err != nil {
		return nil, err
	}
	return &utxo, nil
}

// CreateUTXOID creates an utxo id by concatenating and hashing the txID and the utxoIndex.
// The txID specifies the transaction that produced the utxo.
// The utxoIndex specifies the index of the transaction output that produced the utxo.
func CreateUTXOID(txID []byte, utxoIndex int64) []byte {
	txIDStr := hex.EncodeToString(txID)
	utxoIndexStr := strconv.Itoa(int(utxoIndex))

	var buffer bytes.Buffer
	buffer.WriteString(txIDStr)
	buffer.WriteString(":")
	buffer.WriteString(utxoIndexStr)
	return buffer.Bytes()
}
