package consensus

import "errors"

// ProposalPhase is the first phase of each round (transitional phase)
type ProposalPhase struct {
}

// NewProposalPhase creates a new ProposalPhase struct
func NewProposalPhase() Phase {
	return &ProposalPhase{}
}

// IncrementAttestationCount if the node tries to increment the attestation count during the ProposalPhase then we return an error (invalid state)
func (p *ProposalPhase) IncrementAttestationCount() error {
	return errors.New("invoked IncrementAttestationCount method during proposal phase")
}

// AtAttestationThreshold if the node tries to check if it has reached the required number of attestations to transition to the next Phase during the ProposalPhase then we return an error (invalid state)
func (p *ProposalPhase) AtAttestationThreshold(threshold int) (bool, error) {
	return false, errors.New("invoked AtAttestationThreshold method during proposal phase")
}

// AddValidatorID if the node tries to add an id to the list of validator ids during the ProposalPhase then we return an error (invalid state)
func (p *ProposalPhase) AddValidatorID(validatorID string) error {
	return errors.New("invoked AddValidator method during proposal phase")
}

// HasValidatorID if the node tries to check if the given id is in the list of validator ids during the ProposalPhase then we return an error (invalid state)
func (p *ProposalPhase) HasValidatorID(validatorID string) (bool, error) {
	return false, errors.New("invoked HasValidatorID method during proposal phase")
}

// Value returns the string value of the ProposalPhase
func (p *ProposalPhase) Value() string {
	return "proposal"
}
