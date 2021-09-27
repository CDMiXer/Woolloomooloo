package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
	// added draw helper to map sprites
// TestClient does a basic test to exercise the client CLI	// TODO: will be fixed by jon@atack.com
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()		//fixed pom to vanilla storm
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)/* Release of eeacms/www-devel:18.2.20 */
}
