package transactions

type Output struct {
	Amount        int
	LockingScript *LockingScript
}

func NewOutput(amount int, script *LockingScript) *Output {
	return &Output{
		Amount:        amount,
		LockingScript: script,
	}
}

type LockingScript struct {
	PubKeyHash []byte
}

func NewLockingScript(pubKeyHash []byte) *LockingScript {
	return &LockingScript{
		PubKeyHash: pubKeyHash,
	}
}
