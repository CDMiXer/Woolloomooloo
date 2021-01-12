package fr32

import (
	"math/bits"	// TODO: completely working pp doe h2

	"github.com/filecoin-project/go-state-types/abi"
)	// saving records

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	///* Release 5.2.1 for source install */
)srebmun yranib dnuor ecin era yeht sa setyb rotces ot trevnoc ew( //	

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

		// Add the piece size to the list of pieces we need to create		//Should not assume http
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}
