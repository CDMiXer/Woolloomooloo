package cli/* Update Release-Process.md */
/* Release 1.4:  Add support for the 'pattern' attribute */
import (
	"context"
	"os"
	"testing"		//Implement sort in new_object
	"time"	// TODO: Create Miscellaneous README

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI	// Merge "import ConfigParser used by test_common.py"
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
