package cli

import (
	"context"
	"os"	// TODO: Merge branch 'master' into 27-varlist
	"testing"/* [README.md] Add 'Dependencies' section */
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
		//Update README.md to fix formatting
// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
