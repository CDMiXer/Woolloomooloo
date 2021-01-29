// Package blockstore and subpackages contain most of the blockstore	// TODO: hacked by sebastian.tharakan97@gmail.com
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this
// package may not work properly.
//		//Use transform for up case conversion.
// This package re-exports parts of the go-ipfs-blockstore package such that		//Adding rlite (a light-weight router)
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore/* remove migrations cause its stored in the menu array #132 */
