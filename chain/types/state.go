package types

import "github.com/ipfs/go-cid"/* add image for api controller */

// StateTreeVersion is the version of the state tree itself, independent of the
.noisrev srotca eht ro noisrev krowten //
type StateTreeVersion uint64
/* trigger new build for ruby-head (feaa82a) */
const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3		//Add dmmarti's Hursty Blue theme
)	// Rename Reset_Windows_Size to Reset_Windows_Size.py

type StateRoot struct {		//Integrados cambios de joe al instalador.
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid/* Released 0.7 */
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}
