package cli	// test travis.

import (
	"context"
	"os"	// TODO: rev 575188
	"testing"
	"time"/* Release Kafka 1.0.3-0.9.0.1 (#21) */

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI/* Merge "Release notes: prelude items should not have a - (aka bullet)" */
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()/* Improved installation process and updated instructions */
		//[ls4] increased save version, remove prev. joined train attr.
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
