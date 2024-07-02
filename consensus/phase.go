package consensus

type Phase interface {
	Value() string
}

type ProposalPhase struct {
}

func NewProposalPhase() Phase {
	return &ProposalPhase{}
}

func (p *ProposalPhase) Value() string {
	return "proposal"
}

type PreVotePhase struct {
	PreVotes    int
	PreVoterIDs []string
}

func NewPreVotePhase() Phase {
	return &PreVotePhase{
		PreVotes:    0,
		PreVoterIDs: make([]string, 10),
	}
}

func (p *PreVotePhase) Value() string {
	return "preVote"
}

type PreCommitPhase struct {
	PreCommits      int
	PreCommitterIDs []string
}

func NewPreCommitPhase() Phase {
	return &PreCommitPhase{
		PreCommits:      0,
		PreCommitterIDs: make([]string, 10),
	}
}

func (p *PreCommitPhase) Value() string {
	return "preCommit"
}
