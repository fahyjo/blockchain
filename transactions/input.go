package transactions

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
	PubKey []byte
	Sig    []byte
}

func NewUnlockingScript(pubKey []byte, sig []byte) *UnlockingScript {
	return &UnlockingScript{
		PubKey: pubKey,
		Sig:    sig,
	}
}
