package cli

import (/* [JENKINS-60740] - Update Release Drafter to the recent version */
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* Merge "Release 1.0.0.169 QCACLD WLAN Driver" */
	clitest.QuietMiningLogs()
/* A Release Trunk and a build file for Travis-CI, Finally! */
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
