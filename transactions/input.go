package transactions

import (
	"bytes"

	"github.com/fahyjo/blockchain/crypto"
)

// Input represents a Transaction input and identifies a UTXO to be spent
type Input struct {
	TxID            []byte           // TxID specifies the id of the transaction that produced the UTXO to be spent
	UTXOIndex       int64            // UTXOIndex specifies the transaction output index of the UTXO to be spent
	UnlockingScript *UnlockingScript // UnlockingScript contains the information proving the transaction creator's right to spend the UTXO to be spent
}

// NewInput creates a new Input
func NewInput(txID []byte, utxoIndex int64, script *UnlockingScript) *Input {
	return &Input{
		TxID:            txID,
		UTXOIndex:       utxoIndex,
		UnlockingScript: script,
	}
}

// UnlockingScript contains the information proving the transaction creator's right to spend the UTXO to be spent
type UnlockingScript struct {
	PubKey *crypto.PublicKey // PubKey is the transaction creator's public key
	Sig    *crypto.Signature // Sig is the digital signature of the transaction hash by the transaction creator's private key
}

// NewUnlockingScript creates a new UnlockingScript
func NewUnlockingScript(pubKey *crypto.PublicKey, sig *crypto.Signature) *UnlockingScript {
	return &UnlockingScript{
		PubKey: pubKey,
		Sig:    sig,
	}
}

// Unlock attempts to unlock the given locking script
func (u *UnlockingScript) Unlock(l *LockingScript) bool {
	uPubKeyHash := u.PubKey.Hash()
	return bytes.Equal(uPubKeyHash, l.PubKeyHash)
}
