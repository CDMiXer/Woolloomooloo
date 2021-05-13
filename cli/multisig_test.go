package cli
	// TODO: Automatic changelog generation for PR #8367 [ci skip]
import (
	"context"
	"os"
	"testing"
	"time"		//a7b12730-2e76-11e5-9284-b827eb9e62be

	clitest "github.com/filecoin-project/lotus/cli/test"/* Restore and fix etc/ChangeLog entry. */
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {/* Merge "[INTERNAL] Release notes for version 1.28.31" */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
