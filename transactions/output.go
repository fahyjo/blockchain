package transactions

// Output represents a Transaction output and is used to create a new UTXO
type Output struct {
	Amount        int64
	LockingScript *LockingScript
}

// NewOutput creates a new Output
func NewOutput(amount int64, script *LockingScript) *Output {
	return &Output{
		Amount:        amount,
		LockingScript: script,
	}
}

// LockingScript specifies the conditions under which the utxo to be created can be spent
type LockingScript struct {
	PubKeyHash []byte // PubKeyHash is the hash of the transaction creator's public key
}

// NewLockingScript creates a new LockingScript
func NewLockingScript(pubKeyHash []byte) *LockingScript {
	return &LockingScript{
		PubKeyHash: pubKeyHash,
	}
}
