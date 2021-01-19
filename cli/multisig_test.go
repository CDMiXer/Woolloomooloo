package cli

import (
	"context"	// TODO: RTSS: include OgreUnifiedShader.h unconditionally
	"os"	// Add permissions to 500 error possible causes
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI	// TODO: Reformatted readme
// commands
func TestMultisig(t *testing.T) {/* Merge "Release 1.0.0.201 QCACLD WLAN Driver" */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}		//add 3rd component
