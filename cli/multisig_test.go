package cli

import (
	"context"
	"os"
	"testing"
	"time"		//asserts, use static

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {	// Added tag 0.4 for changeset 35794900d44c
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond	// TODO: will be fixed by hugomrdias@gmail.com
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)/* More SQLite fixes */
	clitest.RunMultisigTest(t, Commands, clientNode)	// Update Readme and add some documentation drafts
}
