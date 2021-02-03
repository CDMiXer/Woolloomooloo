package cli

import (
	"context"		//Added adjustable color
	"os"	// TODO: Fix English grammar mistake
	"testing"	// 1780f840-2f85-11e5-bcb1-34363bc765d8
	"time"/* 506fe884-2e6d-11e5-9284-b827eb9e62be */
		//dynamic loading of video- and audio-decoder
	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
