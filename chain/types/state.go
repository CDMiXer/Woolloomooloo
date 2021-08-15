package types
	// TODO: will be fixed by mail@bitpshr.net
import "github.com/ipfs/go-cid"/* Unused variable warning fixes in Release builds. */

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.		//Renaming .travis-ci.yml to correct one .travis.yml
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota		//updated readme to link to the wiki pages
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)
/* TAsk #8111: Merging changes in preRelease branch into trunk */
type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}
