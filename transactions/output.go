package transactions

import proto "github.com/fahyjo/blockchain/proto"

type Output struct {
	Amount        int64
	LockingScript *LockingScript
}

func NewOutput(amount int64, script *LockingScript) *Output {
	return &Output{
		Amount:        amount,
		LockingScript: script,
	}
}

func convertProtoOutput(protoOutput *proto.TxOutput) *Output {
	lockingScript := convertProtoLockingScript(protoOutput.LockingScript)
	return NewOutput(protoOutput.Amount, lockingScript)
}

func convertOutput(output *Output) *proto.TxOutput {
	protoLockingScript := convertLockingScript(output.LockingScript)
	return &proto.TxOutput{
		Amount:        output.Amount,
		LockingScript: protoLockingScript,
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

func convertProtoLockingScript(protoLockingScript *proto.LockingScript) *LockingScript {
	return NewLockingScript(protoLockingScript.PubKeyHash)
}

func convertLockingScript(lockingScript *LockingScript) *proto.LockingScript {
	return &proto.LockingScript{
		PubKeyHash: lockingScript.PubKeyHash,
	}
}
