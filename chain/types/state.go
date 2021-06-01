package types/* Examine the grib file (currently only checking the edition). */

import "github.com/ipfs/go-cid"/* Move ReleaseChecklist into the developer guide */

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (		//Updated the version number to '0.4.2'.
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1/* Release of eeacms/www:19.11.22 */
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {/* history now only saves the required amount of measurements */
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.		//Correct a nasty misspelling :-)
	Info cid.Cid		//Don't forget let
}
		//Merge "Create monasca-api tempest job"
// TODO: version this.
type StateInfo0 struct{}
