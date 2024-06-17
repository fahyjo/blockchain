package transactions

import (
	"bytes"
	"encoding/gob"

	proto "github.com/fahyjo/blockchain/proto"
)

func EncodeTransaction(tx *Transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
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

func ConvertProtoTransaction(protoTx *proto.Transaction) *Transaction {
	var inputs []*Input
	for _, input := range protoTx.Inputs {
		inputs = append(inputs, convertProtoInput(input))
	}

	var outputs []*Output
	for _, output := range protoTx.Outputs {
		outputs = append(outputs, convertProtoOutput(output))
	}

	return NewTransaction(inputs, outputs)
}

func ConvertTransaction(tx *Transaction) *proto.Transaction {
	var protoInputs []*proto.TxInput
	for _, input := range tx.Inputs {
		protoInputs = append(protoInputs, convertInput(input))
	}

	var protoOutputs []*proto.TxOutput
	for _, output := range tx.Outputs {
		protoOutputs = append(protoOutputs, convertOutput(output))
	}

	return &proto.Transaction{
		Inputs:  protoInputs,
		Outputs: protoOutputs,
	}
}
