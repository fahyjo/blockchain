package consensus

import (
	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

type PreCommitPhase struct {
	PreCommits   int
	ValidatorIDs map[string]bool
}

func NewPreCommitPhase() Phase {
	return &PreCommitPhase{
		PreCommits:   0,
		ValidatorIDs: make(map[string]bool, 10),
	}
}

func (p *PreCommitPhase) IncrementAttestationCount() error {
	p.PreCommits++
	return nil
}

func (p *PreCommitPhase) AtAttestationThreshold(threshold int) (bool, error) {
	return p.PreCommits >= threshold, nil
}

func (p *PreCommitPhase) AddValidatorID(validatorID string) error {
	p.ValidatorIDs[validatorID] = true
	return nil
}

func (p *PreCommitPhase) HasValidatorID(validatorID string) (bool, error) {
	_, ok := p.ValidatorIDs[validatorID]
	return ok, nil
}

func (p *PreCommitPhase) Value() string {
	return "preCommit"
}

type PreCommit struct {
	BlockID []byte
	Sig     *crypto.Signature
	PubKey  *crypto.PublicKey
}

func NewPreCommit(blockID []byte, sig *crypto.Signature, pub *crypto.PublicKey) *PreCommit {
	return &PreCommit{
		BlockID: blockID,
		Sig:     sig,
		PubKey:  pub,
	}
}

func ConvertProtoPreCommit(protoPreCommit *proto.PreCommit) *PreCommit {
	return NewPreCommit(
		protoPreCommit.BlockID,
		crypto.NewSignature(protoPreCommit.Sig),
		crypto.NewPublicKey(protoPreCommit.PubKey),
	)
}
