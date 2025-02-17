package steps

import (
	"github.com/git-town/git-town/v8/src/git"
	"github.com/git-town/git-town/v8/src/hosting"
)

// AbortRebaseStep represents aborting on ongoing merge conflict.
// This step is used in the abort scripts for Git Town commands.
type AbortRebaseStep struct {
	EmptyStep
}

func (step *AbortRebaseStep) Run(run *git.ProdRunner, connector hosting.Connector) error {
	return run.Frontend.AbortRebase()
}
