package consensus

import (
	"github.com/fahyjo/blockchain/crypto"
	proto "github.com/fahyjo/blockchain/proto"
)

type PreVotePhase struct {
	PreVotes     int
	ValidatorIDs map[string]bool
}

func NewPreVotePhase() Phase {
	return &PreVotePhase{
		PreVotes:     0,
		ValidatorIDs: make(map[string]bool, 10),
	}
}

func (p *PreVotePhase) IncrementAttestationCount() error {
	p.PreVotes++
	return nil
}

func (p *PreVotePhase) AtAttestationThreshold(threshold int) (bool, error) {
	return p.PreVotes >= threshold, nil
}

func (p *PreVotePhase) AddValidatorID(validatorID string) error {
	p.ValidatorIDs[validatorID] = true
	return nil
}

func (p *PreVotePhase) HasValidatorID(validatorID string) (bool, error) {
	_, ok := p.ValidatorIDs[validatorID]
	return ok, nil
}

func (p *PreVotePhase) Value() string {
	return "preVote"
}

type PreVote struct {
	BlockID []byte
	Sig     *crypto.Signature
	PubKey  *crypto.PublicKey
}

func NewPreVote(blockID []byte, sig *crypto.Signature, pub *crypto.PublicKey) *PreVote {
	return &PreVote{
		BlockID: blockID,
		Sig:     sig,
		PubKey:  pub,
	}
}

func ConvertProtoPreVote(protoPreVote *proto.PreVote) *PreVote {
	return NewPreVote(
		protoPreVote.BlockID,
		crypto.NewSignature(protoPreVote.Sig),
		crypto.NewPublicKey(protoPreVote.PubKey),
	)
}
