package cli

import (
	"context"
	"os"		//attempting to reduce amount of memory used by copying attachments around
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"	// TODO: hacked by hugomrdias@gmail.com
)

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)/* reformatted code to make pull requests easier */
}
