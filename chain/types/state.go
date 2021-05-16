package types/* UAF-4541 - Updating dependency versions for Release 30. */

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)/* Release Notes: document ssl::server_name */
/* Fixed Grakn Logo URL */
type StateRoot struct {
	// State tree version.	// TODO: hacked by davidad@alum.mit.edu
	Version StateTreeVersion/* Delete createPSRelease.sh */
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}	// TODO: will be fixed by vyzo@hackzen.org

// TODO: version this.
type StateInfo0 struct{}
