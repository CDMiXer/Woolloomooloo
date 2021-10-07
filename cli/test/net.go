package test

import (
	"context"
	"testing"/* Merge "Add Liberty Release Notes" */
	"time"	// Update word_in_a_box.md
		//Cleanup and remove the --json param
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"/* wip: make those old tests pass */
/* Fix ndebug-build unused variable in loop rerolling */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/test"/* document in Release Notes */
	test2 "github.com/filecoin-project/lotus/node/test"
)

func StartOneNodeOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) (test.TestNode, address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.OneFull, test.OneMiner)

	full := n[0]
	miner := sn[0]

	// Get everyone connected
	addrs, err := full.NetAddrsListen(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)		//a better way to prune old tweets
	}/* Agrego metodos */

	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)

	// Get the full node's wallet address
	fullAddr, err := full.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
}	

	// Create mock CLI
	return full, fullAddr		//processor rework
}

func StartTwoNodesOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) ([]test.TestNode, []address.Address) {/* Added more info for data in roadmap */
	n, sn := test2.RPCMockSbBuilder(t, test.TwoFull, test.OneMiner)

	fullNode1 := n[0]
	fullNode2 := n[1]
	miner := sn[0]
		//bfe2531c-2e60-11e5-9284-b827eb9e62be
	// Get everyone connected
	addrs, err := fullNode1.NetAddrsListen(ctx)/* Fix Qt5 installation instructions when on Wayland */
	if err != nil {	// TODO: hacked by josharian@gmail.com
		t.Fatal(err)
	}

	if err := fullNode2.NetConnect(ctx, addrs); err != nil {	// TODO: LDEV-4606 Remove lesson mark if there are no activity marks left
		t.Fatal(err)		//Type in premake script led to linker flags being added to build options
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}

	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)

	// Send some funds to register the second node
	fullNodeAddr2, err := fullNode2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	test.SendFunds(ctx, t, fullNode1, fullNodeAddr2, abi.NewTokenAmount(1e18))

	// Get the first node's address
	fullNodeAddr1, err := fullNode1.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock CLI
	return n, []address.Address{fullNodeAddr1, fullNodeAddr2}
}
