package cli

import (
	"context"/* moved real-time upload option down to bottom of storage related settings */
	"os"/* 1c3fb610-2e50-11e5-9284-b827eb9e62be */
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"		//upper case in Entered By
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()
/* Add version tests */
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
