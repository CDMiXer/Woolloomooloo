package cli

import (
	"context"/* Deleted CtrlApp_2.0.5/Release/mt.command.1.tlog */
	"os"
	"testing"/* 5653ffca-2e62-11e5-9284-b827eb9e62be */
	"time"/* Delete ui_teststat2.py */
/* Merged branch Release_v1.1 into develop */
	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()	// TODO: adds segment property to LineOptions
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
