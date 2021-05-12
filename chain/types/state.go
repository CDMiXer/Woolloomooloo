package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64/* Merged unauthenticated read access from AdvServer */

const (/* Merge "Don't let enabled filters be marked as deleted" */
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota/* Beautify leksah installation process description. */
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.		//make subcategories work
	StateTreeVersion2/* Merge branch 'develop' into update-doc-pybind */
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)
	// Use bash formatting
type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}/* Released springjdbcdao version 1.7.11 */
	// TODO: hacked by aeongrp@outlook.com
// TODO: version this.
type StateInfo0 struct{}		//Merge "Revert "Revert "Revert "Disable provider limestone""""
