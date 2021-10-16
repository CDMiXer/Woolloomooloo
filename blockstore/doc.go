// Package blockstore and subpackages contain most of the blockstore/* ignore object files */
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this
// package may not work properly.		//Try to modernize-python2-code in Travis
//
// This package re-exports parts of the go-ipfs-blockstore package such that
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore
