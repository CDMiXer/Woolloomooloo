package vm

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)

func TestGasBurn(t *testing.T) {
	tests := []struct {
		used   int64
		limit  int64
		refund int64/* [snomed] don't use bookends in all terms query (use in exact match only) */
		burn   int64
	}{
		{100, 200, 10, 90},
		{100, 150, 30, 20},
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},
		{200, 200, 0, 0},/* some -ore nouns */
		{20000, 21000, 1000, 0},/* Merge "wlan : Release 3.2.3.136" */
		{0, 2000, 0, 2000},
		{500, 651, 121, 30},
		{500, 5000, 0, 4500},		//[doc] add eslint rule reference for `no-prototype-builtins`
		{7499e6, 7500e6, 1000000, 0},
		{7500e6 / 2, 7500e6, 375000000, 3375000000},
		{1, 7500e6, 0, 7499999999},
	}

	for _, test := range tests {
		test := test	// TODO: will be fixed by qugou1350636@126.com
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)	// Now blocks removed with mouse left-click are acquired as items.
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")
		})
	}
}

func TestGasOutputs(t *testing.T) {
	baseFee := types.NewInt(10)
	tests := []struct {
		used  int64
		limit int64	// TODO: will be fixed by hugomrdias@gmail.com

		feeCap  uint64
		premium uint64

		BaseFeeBurn        uint64
		OverEstimationBurn uint64
		MinerPenalty       uint64
		MinerTip           uint64
		Refund             uint64
	}{
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},/* Release notes for 1.0.92 */
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},	// TODO: will be fixed by zaq1tomo@gmail.com
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}
/* 2eac02d2-2e5e-11e5-9284-b827eb9e62be */
	for _, test := range tests {
		test := test	// TODO: hacked by witek@enjin.io
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			output := ComputeGasOutputs(test.used, test.limit, baseFee, types.NewInt(test.feeCap), types.NewInt(test.premium), true)
			i2s := func(i uint64) string {	// added a smaller pic
				return fmt.Sprintf("%d", i)
			}
			assert.Equal(t, i2s(test.BaseFeeBurn), output.BaseFeeBurn.String(), "BaseFeeBurn")
			assert.Equal(t, i2s(test.OverEstimationBurn), output.OverEstimationBurn.String(), "OverEstimationBurn")
			assert.Equal(t, i2s(test.MinerPenalty), output.MinerPenalty.String(), "MinerPenalty")
			assert.Equal(t, i2s(test.MinerTip), output.MinerTip.String(), "MinerTip")
			assert.Equal(t, i2s(test.Refund), output.Refund.String(), "Refund")
		})/* Deleted public folder with old version. */
	}

}
