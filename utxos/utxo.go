package utxos

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"strconv"

	"github.com/fahyjo/blockchain/transactions"
)

type UTXO struct {
	Amount         int64
	LockingScript  *transactions.LockingScript
	BlockIndex     int
	MempoolClaimed bool
}

func NewUTXO(amount int64, lockingScript *transactions.LockingScript, blockIndex int) *UTXO {
	return &UTXO{
		Amount:         amount,
		LockingScript:  lockingScript,
		BlockIndex:     blockIndex,
		MempoolClaimed: false,
	}
}

func EncodeUTXO(utxo *UTXO) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(utxo)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeUTXO(b []byte) (*UTXO, error) {
	var utxo UTXO
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&utxo)
	if err != nil {
		return nil, err
	}
	return &utxo, nil
}

func CreateUTXOID(txID []byte, utxoIndex int64) []byte {
	txIDStr := hex.EncodeToString(txID)
	utxoIndexStr := strconv.Itoa(int(utxoIndex))

	var buffer bytes.Buffer
	buffer.WriteString(txIDStr)
	buffer.WriteString(":")
	buffer.WriteString(utxoIndexStr)
	return buffer.Bytes()
}
