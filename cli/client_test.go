package cli
/* Release info updated */
import (
	"context"
	"os"
	"testing"	// fix: Some teleportation issues involving with some blocks.
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)		//remove debugging trace
/* Release 4.3.3 */
// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)	// TODO: Update client.cs
}
