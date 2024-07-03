package consensus

type PreCommitPhase struct {
	PreCommits   int
	ValidatorIDs map[string]bool
}

func NewPreCommitPhase() Phase {
	return &PreCommitPhase{
		PreCommits:   0,
		ValidatorIDs: make(map[string]bool, 10),
	}
}

func (p *PreCommitPhase) IncrementAttestationCount() error {
	p.PreCommits++
	return nil
}

func (p *PreCommitPhase) AtAttestationThreshold(threshold int) (bool, error) {
	return p.PreCommits >= threshold, nil
}

func (p *PreCommitPhase) AddValidatorID(validatorID string) error {
	p.ValidatorIDs[validatorID] = true
	return nil
}

func (p *PreCommitPhase) HasValidatorID(validatorID string) (bool, error) {
	_, ok := p.ValidatorIDs[validatorID]
	return ok, nil
}

func (p *PreCommitPhase) Value() string {
	return "preCommit"
}
