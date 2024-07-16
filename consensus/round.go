package consensus

import "github.com/fahyjo/blockchain/blocks"

// Round every round a new block is committed to the blockchain
type Round struct {
	ProposerID   string        // ProposerID the id of the validator node that is proposing a block to be committed this Round
	BlockID      []byte        // BlockID the id of the block to be committed this Round
	Block        *blocks.Block // Block the block to be committed this Round
	CurrentPhase Phase         // CurrentPhase one of three distinct phases a node can be in each Round
}

// NewRound creates a new Round struct
func NewRound(proposerID string) *Round {
	return &Round{
		ProposerID:   proposerID,
		BlockID:      nil,
		Block:        nil,
		CurrentPhase: NewProposalPhase(),
	}
}

// NextPhase moves to the next Phase of the current Round
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
