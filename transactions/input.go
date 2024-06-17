package transactions

import (
	"bytes"

	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

type Input struct {
	TxID            []byte
	UTXOIndex       int64
	UnlockingScript *UnlockingScript
}

func NewInput(txID []byte, utxoIndex int64, script *UnlockingScript) *Input {
	return &Input{
		TxID:            txID,
		UTXOIndex:       utxoIndex,
		UnlockingScript: script,
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

type UnlockingScript struct {
	PubKey *crypto.PublicKey
	Sig    *crypto.Signature
}

func NewUnlockingScript(pubKey *crypto.PublicKey, sig *crypto.Signature) *UnlockingScript {
	return &UnlockingScript{
		PubKey: pubKey,
		Sig:    sig,
	}
}

func (u *UnlockingScript) Unlock(l *LockingScript) bool {
	uPubKeyHash := u.PubKey.Hash()
	return bytes.Equal(uPubKeyHash, l.PubKeyHash)
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
