package consensus

import "slices"

type Consensus struct {
	AmValidator  bool
	Validators   []string
	roundNumber  int
	CurrentRound *Round
}

func NewConsensus(amValidator bool, validators []string, roundNumber int, currentRound *Round) *Consensus {
	return &Consensus{
		AmValidator:  amValidator,
		Validators:   validators,
		roundNumber:  roundNumber,
		CurrentRound: currentRound,
	}
}

func (c *Consensus) IsValidator(validatorID string) bool {
	return slices.Contains(c.Validators, validatorID)
}

func (c *Consensus) NextRound() {
	c.roundNumber++
	proposerID := c.Validators[c.roundNumber%len(c.Validators)]
	c.CurrentRound = NewRound(proposerID)
}
