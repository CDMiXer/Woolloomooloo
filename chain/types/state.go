package types/* 9ae1eaf2-2e68-11e5-9284-b827eb9e62be */

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.	// modified delete icon. Fixed a problem with Delete from the menu.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.		//75cdcbc0-2e4d-11e5-9284-b827eb9e62be
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.		//merge url identifies branch
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.	// Reverting changes to scanAllRequest
type StateInfo0 struct{}
