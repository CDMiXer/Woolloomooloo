package types/* 2516e350-2e67-11e5-9284-b827eb9e62be */

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2		//Delete insSpace.png
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {		//Cleaned up Rakefile file includes
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid		//Transaktionen aufgeteilt - loeschen wieder im before
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.	// Slimming the css down
type StateInfo0 struct{}
