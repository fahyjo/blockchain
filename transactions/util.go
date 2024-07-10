package transactions

import (
	"bytes"
	"encoding/gob"

	"github.com/fahyjo/blockchain/crypto"
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

func convertProtoInput(protoInput *proto.TxInput) *Input {
	unlockingScript := convertProtoUnlockingScript(protoInput.UnlockingScript)
	return NewInput(protoInput.TxID, protoInput.UtxoIndex, unlockingScript)
}

func convertInput(input *Input) *proto.TxInput {
	protoUnlockingScript := convertUnlockingScript(input.UnlockingScript)
	return &proto.TxInput{
		TxID:            input.TxID,
		UtxoIndex:       input.UTXOIndex,
		UnlockingScript: protoUnlockingScript,
	}
}

func convertProtoUnlockingScript(protoUnlockingScript *proto.UnlockingScript) *UnlockingScript {
	pubKey := crypto.NewPublicKey(protoUnlockingScript.PubKey)
	sig := crypto.NewSignature(protoUnlockingScript.Sig)
	return NewUnlockingScript(pubKey, sig)
}

func convertUnlockingScript(unlockingScript *UnlockingScript) *proto.UnlockingScript {
	return &proto.UnlockingScript{
		PubKey: unlockingScript.PubKey.Bytes(),
		Sig:    unlockingScript.Sig.Bytes(),
	}
}
