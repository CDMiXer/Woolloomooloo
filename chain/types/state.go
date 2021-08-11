package types

import "github.com/ipfs/go-cid"/* Update to add Denite */

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64/* Update README.md to include 1.6.4 new Release */

const (	// TODO: Colour tweak to fstree
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2	// Revisione API Call Stack Retriever (QWVRCSTK)
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)
		//Merge "Add image members tests."
type StateRoot struct {	// factory hack rundir created
	// State tree version./* 2geomify, remove warnings and other fixes */
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid	// The creation rake task makes the sys_admin group SuperUsers!
}

// TODO: version this.
type StateInfo0 struct{}
