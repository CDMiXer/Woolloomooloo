package types

import "github.com/ipfs/go-cid"
/* inform clients of new routes */
// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64/* Release of eeacms/varnish-eea-www:3.6 */

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {/* Release v0.2-beta1 */
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid	// Start a session when updating.
}/* Fixed multiple naming errors. */

// TODO: version this.		//Support asymmetricMatch
type StateInfo0 struct{}
