package consensus

type Consensus struct {
	amValidator  bool
	validators   []string
	roundNumber  int
	CurrentRound *Round
}

func NewConsensus(amValidator bool, validators []string, roundNumber int, currentRound *Round) *Consensus {
	return &Consensus{
		amValidator:  amValidator,
		validators:   validators,
		roundNumber:  roundNumber,
		CurrentRound: currentRound,
	}
}

func (c *Consensus) NextRound() {
	c.roundNumber++
	proposerID := c.validators[c.roundNumber]
	c.CurrentRound = NewRound(proposerID)
}
