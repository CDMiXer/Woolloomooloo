// Package blockstore and subpackages contain most of the blockstore
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this	// adds functionality to add/delete/remove all dynos for a specific app
// package may not work properly.
///* Clarifies explanations around Data Interface code */
// This package re-exports parts of the go-ipfs-blockstore package such that
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore
