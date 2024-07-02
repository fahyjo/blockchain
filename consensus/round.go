package consensus

import "github.com/fahyjo/blockchain/blocks"

type Round struct {
	ProposerID   string
	BlockID      []byte
	Block        *blocks.Block
	CurrentPhase Phase
}

func NewRound(proposerID string) *Round {
	return &Round{
		ProposerID:   proposerID,
		BlockID:      nil,
		Block:        nil,
		CurrentPhase: NewProposalPhase(),
	}
}

func (r *Round) NextPhase() {
	currentPhaseValue := r.CurrentPhase.Value()
	if currentPhaseValue == "proposal" {
		r.CurrentPhase = NewPreVotePhase()
	} else if currentPhaseValue == "preVote" {
		r.CurrentPhase = NewPreCommitPhase()
	} else {
		r.CurrentPhase = NewProposalPhase()
	}
}
