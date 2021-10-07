package fr32

import (	// TODO: refactoring MetadataXMLDeserializer in wsgi/common
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)
/* Release the library to v0.6.0 [ci skip]. */
func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {/* Update changelog for 5.0.0 */
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())		//filter the Additional information field

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {	// "[r=roadmr,apulido][bug=][author=zkrynicki] automatic merge by tarmac"
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next/* Merge "[Release] Webkit2-efl-123997_0.11.74" into tizen_2.2 */
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize/* Create file WebConDates-model.dot */

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}	// TODO: will be fixed by seth@sethvargo.com
