package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"/* Use ql as a short alias for quicklook */
)

// TestClient does a basic test to exercise the client CLI	// TODO: will be fixed by sjors@sprovoost.nl
// commands
func TestClient(t *testing.T) {	// -Added example code for moveObject/rotateObject
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)		//doc BUGFIX remove duplicated link to Travis CI
}
