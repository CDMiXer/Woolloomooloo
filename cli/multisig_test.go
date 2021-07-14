package cli

import (
	"context"/* Create docs/examples.md */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
	// Added description, creator and assists
// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* Updated build [ci skip] */
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond/* Update configDox */
	ctx := context.Background()/* Merge "msm: mdss: un map dsi transmit buffer properly" */
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)		//disable analyse option if there is no feature display
	clitest.RunMultisigTest(t, Commands, clientNode)
}
