package consensus

import "errors"

type Phase interface {
	IncrementAttestationCount() error
	AddValidator(validatorID string) error
	Value() string
}

type ProposalPhase struct {
}

func NewProposalPhase() Phase {
	return &ProposalPhase{}
}

func (p *ProposalPhase) IncrementAttestationCount() error {
	return errors.New("invoked IncrementAttestationCount method on ProposalPhase Phase")
}

func (p *ProposalPhase) AddValidator(validatorID string) error {
	return errors.New("invoked AddValidator method on ProposalPhase Phase")
}

func (p *ProposalPhase) Value() string {
	return "proposal"
}

type PreVotePhase struct {
	PreVotes     int
	ValidatorIDs map[string]bool
}

func NewPreVotePhase() Phase {
	return &PreVotePhase{
		PreVotes:     0,
		ValidatorIDs: make(map[string]bool, 10),
	}
}

func (p *PreVotePhase) IncrementAttestationCount() error {
	p.PreVotes++
	return nil
}

func (p *PreVotePhase) AddValidator(validatorID string) error {
	p.ValidatorIDs[validatorID] = true
	return nil
}

func (p *PreVotePhase) Value() string {
	return "preVote"
}

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

func (p *PreCommitPhase) AddValidator(validatorID string) error {
	p.ValidatorIDs[validatorID] = true
	return nil
}

func (p *PreCommitPhase) Value() string {
	return "preCommit"
}
