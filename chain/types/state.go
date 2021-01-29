package types	// TODO: will be fixed by arachnid@notdot.net

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
	// Merge "sphinxext: Start parsing 'DocumentedRuleDefault.description' as rST"
const (
	// StateTreeVersion0 corresponds to actors < v2./* Added subject to internal server */
	StateTreeVersion0 StateTreeVersion = iota/* Release v1.01 */
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
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this./* re-organize doInvoke method for better Exception report */
type StateInfo0 struct{}
