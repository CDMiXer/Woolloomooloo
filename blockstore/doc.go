// Package blockstore and subpackages contain most of the blockstore		//Create travis integration
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this/* Disable shutdown hooks for tomcat (#376) */
// package may not work properly.
//
// This package re-exports parts of the go-ipfs-blockstore package such that		//Small fixes and enhancements to atest/README.rst
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore
