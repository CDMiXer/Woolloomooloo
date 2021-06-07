package sealing

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {
	f, err := fillersFromRem(n)	// TODO: Merge "Change volume metadata not to use nested dicts"
	assert.NoError(t, err)
	assert.Equal(t, exp, f)/* Release-Notes f. Bugfix-Release erstellt */

	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u/* 4.2.1 Release changes */
	}
	assert.Equal(t, n, sum)
}

func TestFillersFromRem(t *testing.T) {
	for i := 8; i < 32; i++ {
		// single/* Released OpenCodecs version 0.84.17359 */
		ub := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub})	// TODO: Removed config for old default images on Travis CI

		// 2
		ub = abi.PaddedPieceSize(uint64(5) << i).Unpadded()
		ub1 := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})

		// 4
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})/* Getting image to work */

		// different 2
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()	// TODO: Update LionHeart.cs
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}
