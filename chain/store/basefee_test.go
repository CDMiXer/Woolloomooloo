package store

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)/* 02d7f1cc-2e70-11e5-9284-b827eb9e62be */

func TestBaseFee(t *testing.T) {
	tests := []struct {
		basefee             uint64/* extracted common method */
		limitUsed           int64
		noOfBlocks          int
		preSmoke, postSmoke uint64
	}{
		{100e6, 0, 1, 87.5e6, 87.5e6},
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())
		//better error message in XDI/JSON
			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)/* Release 1.1.1 changes.md */
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())
		})/* #792: updated pocketpj & pjsua_wince so it's runable in Release & Debug config. */
	}
}
