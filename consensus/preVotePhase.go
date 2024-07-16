package consensus

import (
	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

// PreVotePhase the node is waiting for 2/3 of the validator nodes to attest to the validity of the block to be committed during the current Round
type PreVotePhase struct {
	preVotes     int             // preVotes the number of validator nodes that have approved of the current Round's block
	validatorIDs map[string]bool // validatorIDs the ids of the validator nodes that have approved of the current Round's block
}

// NewPreVotePhase creates a new PreVotePhase struct
func NewPreVotePhase() Phase {
	return &PreVotePhase{
		preVotes:     0,
		validatorIDs: make(map[string]bool, 10),
	}
}

// IncrementAttestationCount increments the number of validator nodes that have approved of the current Round's block
func (p *PreVotePhase) IncrementAttestationCount() error {
	p.preVotes++
	return nil
}

// AtAttestationThreshold returns true if the node has received block attestations from 2/3 of the validator nodes, returns false otherwise
func (p *PreVotePhase) AtAttestationThreshold(threshold int) (bool, error) {
	return p.preVotes >= threshold, nil
}

// AddValidatorID adds the given validator id to the map of validator ids that have approved of the current Round's block
func (p *PreVotePhase) AddValidatorID(validatorID string) error {
	p.validatorIDs[validatorID] = true
	return nil
}

// HasValidatorID returns true if the validator with the given id has approved of the current Round's block
func (p *PreVotePhase) HasValidatorID(validatorID string) (bool, error) {
	_, ok := p.validatorIDs[validatorID]
	return ok, nil
}

// Value returns the string value of the PreVotePhase
func (p *PreVotePhase) Value() string {
	return "preVote"
}

// PreVote broadcast by a validator after they have validated the current Round's block
type PreVote struct {
	BlockID []byte            // BlockID the id of the block to be committed during the current Round
	Sig     *crypto.Signature // Sig the validator node's signature created by the ed25519 private key of the validator and the BlockID
	PubKey  *crypto.PublicKey // PubKey the ed25519 public key corresponding to the ed25519 private key of the validator that broadcast this PreVote
}

// NewPreVote creates a new PreVote struct
func NewPreVote(blockID []byte, sig *crypto.Signature, pub *crypto.PublicKey) *PreVote {
	return &PreVote{
		BlockID: blockID,
		Sig:     sig,
		PubKey:  pub,
	}
}

// ConvertProtoPreVote converts the given proto preVote into a domain PreVote
func ConvertProtoPreVote(protoPreVote *proto.PreVote) *PreVote {
	return NewPreVote(
		protoPreVote.BlockID,
		crypto.NewSignature(protoPreVote.Sig),
		crypto.NewPublicKey(protoPreVote.PubKey),
	)
}
