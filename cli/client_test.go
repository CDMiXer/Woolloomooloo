package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI	// TODO: hacked by seth@sethvargo.com
// commands
func TestClient(t *testing.T) {/* Merge branch 'master' into thomas_semoss_dev6 */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()/* Release 0.48 */
/* Release notes for 0.4 */
	blocktime := 5 * time.Millisecond/* Serialize query execution during hupped query */
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
