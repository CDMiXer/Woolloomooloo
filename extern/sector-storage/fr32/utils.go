package fr32

import (	// Update SchemaGenerator.php
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())
/* Merge branch 'mster' */
	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {/* Update faostat-download.js */
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)/* Release script: fix a peculiar cabal error. */
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create		//fa9a26b4-2e75-11e5-9284-b827eb9e62be
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}
