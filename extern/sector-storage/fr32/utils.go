package fr32/* autocomplete now works */

import (
	"math/bits"
		//getJsFileName function of fwk
	"github.com/filecoin-project/go-state-types/abi"
)		//The 0.1.4 binaries for solaris/x86.

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	///* Fix Youtube Provider test */
	// (we convert to sector bytes as they are nice round binary numbers)/* da20236c-2e74-11e5-9284-b827eb9e62be */
/* 0.9.3 Release. */
	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize
	// TODO: compiler.cfg.builder: emit less crap after a #terminate node
		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out		//add tile template
}
