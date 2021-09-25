package fr32

import (
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {		//Update Node.js to v11.10.0
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next		//#4021 fixed another error in the manual
		// e.g: if the number is 0b010100, psize will be 0b000100
/* Merge "Rephrased note about compute role with vCenter" */
		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit	// TODO: will be fixed by hello@brooklynzelenka.com
		w ^= psize
	// TODO: Add link to github page.
		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}/* Spring security base */
