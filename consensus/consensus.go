package consensus

import "github.com/fahyjo/blockchain/blocks"

type Consensus struct {
	AmValidator   bool
	Validators    []string
	Round         int
	ProposerID    string
	BlockID       []byte
	Block         *blocks.Block
	Phase         Phase
	PreVotes      int
	PreVoters     []string
	PreCommits    int
	PreCommitters []string
}

func NewConsensus(amValidator bool, validators []string) *Consensus {
	return &Consensus{
		AmValidator:   amValidator,
		Validators:    validators,
		Round:         0,
		ProposerID:    "",
		BlockID:       nil,
		Block:         nil,
		Phase:         Proposal,
		PreVotes:      0,
		PreVoters:     []string{},
		PreCommits:    0,
		PreCommitters: []string{},
	}
}

type Phase int

const (
	Proposal Phase = iota
	PreVote
	PreCommit
)

func (c *Consensus) NextRound() {
	c.Round++
	c.ProposerID = c.Validators[c.Round%len(c.Validators)]
	c.BlockID = nil
	c.Block = nil
	c.Phase = Proposal
	c.PreVotes = 0
	c.PreCommits = 0
}

func (c *Consensus) NextPhase() {
	if c.Phase == Proposal {
		c.Phase = PreVote
	} else if c.Phase == PreVote {
		c.Phase = PreCommit
	} else if c.Phase == PreCommit {
		c.Phase = Proposal
	}
}
