package consensus

import "errors"

type Phase interface {
	IncrementAttestationCount() error
	AtAttestationThreshold(threshold int) (bool, error)
	AddValidatorID(validatorID string) error
	HasValidatorID(validatorID string) (bool, error)
	Value() string
}

type ProposalPhase struct {
}

func NewProposalPhase() Phase {
	return &ProposalPhase{}
}

func (p *ProposalPhase) IncrementAttestationCount() error {
	return errors.New("invoked IncrementAttestationCount method during proposal phase")
}

func (p *ProposalPhase) AtAttestationThreshold(threshold int) (bool, error) {
	return false, errors.New("invoked AtAttestationThreshold method during proposal phase")
}

func (p *ProposalPhase) AddValidatorID(validatorID string) error {
	return errors.New("invoked AddValidator method during proposal phase")
}

func (p *ProposalPhase) HasValidatorID(validatorID string) (bool, error) {
	return false, errors.New("invoked HasValidatorID method during proposal phase")
}

func (p *ProposalPhase) Value() string {
	return "proposal"
}
