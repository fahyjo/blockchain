package consensus

import (
	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

// PreCommitPhase the node is waiting for 2/3 of the validator nodes to indicate they will commit the current Round's block
type PreCommitPhase struct {
	preCommits   int             // preCommits the number of validator nodes that will commit the current Round's block
	validatorIDs map[string]bool // validatorIDs the ids of the validator nodes that will commit the current Round's block
}

// NewPreCommitPhase creates a new PreCommitPhase struct
func NewPreCommitPhase() Phase {
	return &PreCommitPhase{
		preCommits:   0,
		validatorIDs: make(map[string]bool, 10),
	}
}

// IncrementAttestationCount increments the number of validator nodes that will commit the current Round's block
func (p *PreCommitPhase) IncrementAttestationCount() error {
	p.preCommits++
	return nil
}

// AtAttestationThreshold returns true if the node has received block attestations from 2/3 of the validator nodes, returns false otherwise
func (p *PreCommitPhase) AtAttestationThreshold(threshold int) (bool, error) {
	return p.preCommits >= threshold, nil
}

// AddValidatorID adds the given validator id to the map of validator ids that will commit the current Round's block
func (p *PreCommitPhase) AddValidatorID(validatorID string) error {
	p.validatorIDs[validatorID] = true
	return nil
}

// HasValidatorID returns true if the validator with the given id will commit the current Round's block
func (p *PreCommitPhase) HasValidatorID(validatorID string) (bool, error) {
	_, ok := p.validatorIDs[validatorID]
	return ok, nil
}

// Value returns the string value of the PreVotePhase
func (p *PreCommitPhase) Value() string {
	return "preCommit"
}

// PreCommit broadcast by a validator after they have received PreVote messages from 2/3 of the validators
type PreCommit struct {
	BlockID []byte
	Sig     *crypto.Signature
	PubKey  *crypto.PublicKey
}

// NewPreCommit creates a new PreCommit struct
func NewPreCommit(blockID []byte, sig *crypto.Signature, pub *crypto.PublicKey) *PreCommit {
	return &PreCommit{
		BlockID: blockID,
		Sig:     sig,
		PubKey:  pub,
	}
}

// ConvertProtoPreCommit converts the given proto preCommit into a domain PreCommit
func ConvertProtoPreCommit(protoPreCommit *proto.PreCommit) *PreCommit {
	return NewPreCommit(
		protoPreCommit.BlockID,
		crypto.NewSignature(protoPreCommit.Sig),
		crypto.NewPublicKey(protoPreCommit.PubKey),
	)
}
