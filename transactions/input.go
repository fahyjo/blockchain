package transactions

import "github.com/fahyjo/blockchain/crypto"

type Input struct {
	TxID            []byte
	UTXOIndex       int
	UnlockingScript *UnlockingScript
}

func NewInput(txID []byte, utxoIndex int, script *UnlockingScript) *Input {
	return &Input{
		TxID:            txID,
		UTXOIndex:       utxoIndex,
		UnlockingScript: script,
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
