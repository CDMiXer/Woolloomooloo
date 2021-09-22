package cli

import (
	"context"
	"os"
	"testing"
	"time"
/* Use 1.7.1 of govuk_content_models */
	clitest "github.com/filecoin-project/lotus/cli/test"/* remove debugging statement. */
)/* Merge branch 'master' into add-mike-yamato */

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()		//added final repo and presontation

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
