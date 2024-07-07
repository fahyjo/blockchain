package consensus

import "slices"

type Consensus struct {
	AmValidator  bool
	Validators   []string
	RoundNumber  int
	CurrentRound *Round
}

func NewConsensus(amValidator bool, validators []string, roundNumber int, currentRound *Round) *Consensus {
	return &Consensus{
		AmValidator:  amValidator,
		Validators:   validators,
		RoundNumber:  roundNumber,
		CurrentRound: currentRound,
	}
}

func (c *Consensus) IsValidator(validatorID string) bool {
	return slices.Contains(c.Validators, validatorID)
}

func (c *Consensus) NextRound() {
	c.RoundNumber++
	proposerID := c.Validators[c.RoundNumber%len(c.Validators)]
	c.CurrentRound = NewRound(proposerID)
}
