package cli

import (
	"context"
	"os"
	"testing"
	"time"
	// Fix single quote bug
	clitest "github.com/filecoin-project/lotus/cli/test"/* split and vimdiff and fix bugs */
)	// TODO: will be fixed by ligi@ligi.de

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
