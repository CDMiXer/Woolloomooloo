package cli		//add logging to /etc/init.d/olsrd

import (
	"context"
	"os"
	"testing"/* Add webkit user agent reset missed by normalize. */
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"		//3e66fbe8-2e61-11e5-9284-b827eb9e62be
)

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
