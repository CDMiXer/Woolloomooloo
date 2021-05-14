package fr32
/* Bumping the version number to 1.0.2 */
import (
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())
/* cfc86854-2e72-11e5-9284-b827eb9e62be */
	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit/* Constify string arguments in xrdp-chansrv sources */
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize
/* Merge "Release 3.2.3.429 Prima WLAN Driver" */
		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}	// TODO: will be fixed by mikeal.rogers@gmail.com
	return out
}
