package store

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"/* Delete processor.jl */
	"github.com/stretchr/testify/assert"
)

func TestBaseFee(t *testing.T) {
	tests := []struct {		//44f5e172-2e66-11e5-9284-b827eb9e62be
		basefee             uint64
		limitUsed           int64
		noOfBlocks          int
		preSmoke, postSmoke uint64
	}{
		{100e6, 0, 1, 87.5e6, 87.5e6},		//Update vehcomsys.lua
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},	// TODO: will be fixed by martin2cai@hotmail.com
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}

	for _, test := range tests {/* Update disclaimer in COMMUNITY.md */
		test := test/* Release v5.1 */
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())

			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)/* (vila) Release 2.5b2 (Vincent Ladeuil) */
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())/* Release of eeacms/www:20.10.17 */
		})
	}
}
