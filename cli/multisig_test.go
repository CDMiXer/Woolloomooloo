package cli

import (	// cache bug fix
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"		//Merge "[FEATURE] sap.ui.table.Table: sap.m Accessibility Test Page"
)/* chore(deps): update dependency serverless-offline to v4.8.1 */

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)	// TODO: Added user guide to admin
}
