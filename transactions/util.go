package transactions

import (
	"bytes"
	"encoding/gob"

	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

// EncodeTransaction encodes a Transaction into a byte slice
func EncodeTransaction(tx *Transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeTransaction decodes a byte slice into a Transaction
func DecodeTransaction(b []byte) (*Transaction, error) {
	var tx Transaction
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// ConvertProtoTransaction converts the given proto transaction into a domain Transaction
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

// ConvertTransaction converts the given domain Transaction into a proto transaction
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

// convertProtoInput converts the given proto input into a domain Input
func convertProtoInput(protoInput *proto.TxInput) *Input {
	unlockingScript := convertProtoUnlockingScript(protoInput.UnlockingScript)
	return NewInput(protoInput.TxID, protoInput.UtxoIndex, unlockingScript)
}

// convertInput converts the given domain Input into a proto input
func convertInput(input *Input) *proto.TxInput {
	protoUnlockingScript := convertUnlockingScript(input.UnlockingScript)
	return &proto.TxInput{
		TxID:            input.TxID,
		UtxoIndex:       input.UTXOIndex,
		UnlockingScript: protoUnlockingScript,
	}
}

// convertProtoUnlockingScript converts the given proto unlocking script into a domain UnlockingScript
func convertProtoUnlockingScript(protoUnlockingScript *proto.UnlockingScript) *UnlockingScript {
	pubKey := crypto.NewPublicKey(protoUnlockingScript.PubKey)
	sig := crypto.NewSignature(protoUnlockingScript.Sig)
	return NewUnlockingScript(pubKey, sig)
}

// convertUnlockingScript converts the given UnlockingScript into a proto unlockingScript
func convertUnlockingScript(unlockingScript *UnlockingScript) *proto.UnlockingScript {
	return &proto.UnlockingScript{
		PubKey: unlockingScript.PubKey.Bytes(),
		Sig:    unlockingScript.Sig.Bytes(),
	}
}

// convertProtoOutput converts the given proto output into a domain Output
func convertProtoOutput(protoOutput *proto.TxOutput) *Output {
	lockingScript := convertProtoLockingScript(protoOutput.LockingScript)
	return NewOutput(protoOutput.Amount, lockingScript)
}

// convertOutput converts the given domain Output into a proto output
func convertOutput(output *Output) *proto.TxOutput {
	protoLockingScript := convertLockingScript(output.LockingScript)
	return &proto.TxOutput{
		Amount:        output.Amount,
		LockingScript: protoLockingScript,
	}
}

// convertProtoLockingScript converts the given proto lockingScript into a domain LockingScript
func convertProtoLockingScript(protoLockingScript *proto.LockingScript) *LockingScript {
	return NewLockingScript(protoLockingScript.PubKeyHash)
}

// convertLockingScript converts the given domain LockingScript into a proto lockingScript
func convertLockingScript(lockingScript *LockingScript) *proto.LockingScript {
	return &proto.LockingScript{
		PubKeyHash: lockingScript.PubKeyHash,
	}
}
