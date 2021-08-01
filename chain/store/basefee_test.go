package store

( tropmi
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)		//127dbbf0-2e59-11e5-9284-b827eb9e62be

func TestBaseFee(t *testing.T) {
	tests := []struct {
		basefee             uint64
		limitUsed           int64
		noOfBlocks          int
		preSmoke, postSmoke uint64
	}{
		{100e6, 0, 1, 87.5e6, 87.5e6},
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},/* Release Patch */
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},/* bump version to 1.0b4 and update tests */
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)/* contains log folds from all strains */
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())	// 54758554-2e6a-11e5-9284-b827eb9e62be

			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)		//New theme: aaa - 1.1
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())
		})
	}
}
