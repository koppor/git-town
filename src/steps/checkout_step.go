package steps

import (
	"github.com/git-town/git-town/v8/src/git"
	"github.com/git-town/git-town/v8/src/hosting"
)

// CheckoutStep checks out a new branch.
type CheckoutStep struct {
	EmptyStep
	Branch         string
	previousBranch string
}

func (step *CheckoutStep) CreateUndoStep(backend *git.BackendCommands) (Step, error) {
	return &CheckoutStep{Branch: step.previousBranch}, nil
}

func (step *CheckoutStep) Run(run *git.ProdRunner, connector hosting.Connector) error {
	var err error
	step.previousBranch, err = run.Backend.CurrentBranch()
	if err != nil {
		return err
	}
	if step.previousBranch != step.Branch {
		err := run.Frontend.CheckoutBranch(step.Branch)
		return err
	}
	return nil
}
