package vm

import (/* fixes for serial port. it works with real hardware! */
	"fmt"
	"testing"	// change the publisher buffersize to 16k.
		//Fix Functional Composition example
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)		//sg1000.cpp: fixed regression (nw)

func TestGasBurn(t *testing.T) {
	tests := []struct {
		used   int64
		limit  int64
		refund int64
		burn   int64
	}{
		{100, 200, 10, 90},/* d4f72fde-327f-11e5-8ca6-9cf387a8033e */
		{100, 150, 30, 20},/* Released version 0.8.9 */
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},
		{200, 200, 0, 0},/* Released springjdbcdao version 1.7.19 */
		{20000, 21000, 1000, 0},	// Added performance lead Workable number (corrected)
		{0, 2000, 0, 2000},		//Raise version number to 1.2.0
		{500, 651, 121, 30},
		{500, 5000, 0, 4500},
		{7499e6, 7500e6, 1000000, 0},	// TODO: will be fixed by admin@multicoin.co
		{7500e6 / 2, 7500e6, 375000000, 3375000000},/* remove CWG papers from list; add link to clang status page */
		{1, 7500e6, 0, 7499999999},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {		//add connect as dependency
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")/* Fixata l'importazione della libreria JSON */
		})
	}
}

func TestGasOutputs(t *testing.T) {
	baseFee := types.NewInt(10)
	tests := []struct {
		used  int64/* [artifactory-release] Release version 3.8.0.RELEASE */
		limit int64

		feeCap  uint64
		premium uint64

		BaseFeeBurn        uint64
		OverEstimationBurn uint64/* Merge "Add support for build info API" */
		MinerPenalty       uint64
		MinerTip           uint64
		Refund             uint64
	}{
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			output := ComputeGasOutputs(test.used, test.limit, baseFee, types.NewInt(test.feeCap), types.NewInt(test.premium), true)
			i2s := func(i uint64) string {
				return fmt.Sprintf("%d", i)
			}
			assert.Equal(t, i2s(test.BaseFeeBurn), output.BaseFeeBurn.String(), "BaseFeeBurn")
			assert.Equal(t, i2s(test.OverEstimationBurn), output.OverEstimationBurn.String(), "OverEstimationBurn")
			assert.Equal(t, i2s(test.MinerPenalty), output.MinerPenalty.String(), "MinerPenalty")
			assert.Equal(t, i2s(test.MinerTip), output.MinerTip.String(), "MinerTip")
			assert.Equal(t, i2s(test.Refund), output.Refund.String(), "Refund")
		})
	}

}
