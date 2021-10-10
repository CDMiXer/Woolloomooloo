package types
/* Merge "Release 1.0.0.58 QCACLD WLAN Driver" */
import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
.2v < srotca ot sdnopserroc 0noisreVeerTetatS //	
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2		//checkpoint: fixed fallout of unsafeCast (mostly), seems to work better
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3/* Release 0.17.0 */
)

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
