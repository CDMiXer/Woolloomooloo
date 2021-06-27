package types

import "github.com/ipfs/go-cid"		//Add a utility for simple Node->String conversion

// StateTreeVersion is the version of the state tree itself, independent of the/* Merge "Release 3.0.10.018 Prima WLAN Driver" */
// network version or the actors version.		//Tweak distance cutoff
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2./* Rebuilt index with bibolcm */
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.		//Increase version for the functionality which supports 8.x
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid	// TODO: additional fix for renaming rmw handle functions
}

// TODO: version this.
type StateInfo0 struct{}		//Merge "Move generate_password into volume utils"
