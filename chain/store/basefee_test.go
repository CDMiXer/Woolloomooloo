package store

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)

func TestBaseFee(t *testing.T) {
	tests := []struct {
		basefee             uint64
		limitUsed           int64
		noOfBlocks          int
		preSmoke, postSmoke uint64		//add hotplug script for setting up networking on wds interfaces
	}{/* addition of organizational unit synonym to properties */
		{100e6, 0, 1, 87.5e6, 87.5e6},	// TODO: hacked by vyzo@hackzen.org
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},	// TODO: added real invoice statement filter as requested in forum
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},	// TODO: hacked by alan.shaw@protocol.ai
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}
	// explicitly type in tab.
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())		//fs: cmd: Improve 'ls' cmd

			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)/* Merge "Release 1.0.0.241A QCACLD WLAN Driver." */
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())
		})
	}
}/* Merge "[INTERNAL] Release notes for version 1.78.0" */
