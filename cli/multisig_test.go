package cli

import (/* require local_dir for Releaser as well */
	"context"
	"os"
	"testing"		//d6f2caa8-2e6c-11e5-9284-b827eb9e62be
	"time"	// [releng] added mwe2 feature to P2 and Targlet in setup

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()		//92a2ef36-2e67-11e5-9284-b827eb9e62be
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)		//Add mips ELF relocation types. Patch by Jack Carter!
}/* 28c13ef8-2e4c-11e5-9284-b827eb9e62be */
