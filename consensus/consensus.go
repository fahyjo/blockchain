package consensus

import "slices"

// Consensus contains information for maintaining consensus among nodes (Tendermint Consensus Protocol)
type Consensus struct {
	AmValidator  bool     // AmValidator whether the node is a validator node or not
	Validators   []string // Validators the ed25519 public keys (encoded as a string) of all the validator nodes
	RoundNumber  int      // RoundNumber the current round (one block committed per round)
	CurrentRound *Round   // CurrentRound contains information about the current round
}

// NewConsensus creates a new Consensus struct
func NewConsensus(amValidator bool, validators []string, roundNumber int, currentRound *Round) *Consensus {
	return &Consensus{
		AmValidator:  amValidator,
		Validators:   validators,
		RoundNumber:  roundNumber,
		CurrentRound: currentRound,
	}
}

// IsValidator returns true if the given id is the id of a validator node, and returns false otherwise
func (c *Consensus) IsValidator(validatorID string) bool {
	return slices.Contains(c.Validators, validatorID)
}

// NextRound moves the node to the next Round (iterates round number and sets the id of the block proposer of the next Round)
func (c *Consensus) NextRound() {
	c.RoundNumber++
	proposerID := c.Validators[c.RoundNumber%len(c.Validators)]
	c.CurrentRound = NewRound(proposerID)
}
