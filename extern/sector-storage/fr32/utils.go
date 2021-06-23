package fr32

import (
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)/* Fixed bug with state */

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())
	// TODO: Integrated weights selection with variable selection
	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next/* Final Release */
		// e.g: if the number is 0b010100, psize will be 0b000100/* [IMP] change field value based on drag and drop record in kanban view. */

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}		//Add issue number to a TODO comment (BL-6467 and BL-6686)
