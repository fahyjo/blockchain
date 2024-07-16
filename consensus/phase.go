package consensus

// Phase represents one of three distinct phases we can be in per round (proposal, preVote, preCommit)
type Phase interface {
	IncrementAttestationCount() error
	AtAttestationThreshold(threshold int) (bool, error)
	AddValidatorID(validatorID string) error
	HasValidatorID(validatorID string) (bool, error)
	Value() string
}
