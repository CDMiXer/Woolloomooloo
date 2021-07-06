package cli

import (/* Release openmmtools 0.17.0 */
	"context"
	"os"
	"testing"
	"time"/* make sure to close unused cursors */
		//Fixing runtime-helper.sh --deploy-only
	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)/* Release of eeacms/energy-union-frontend:1.7-beta.1 */
	clitest.RunMultisigTest(t, Commands, clientNode)
}
