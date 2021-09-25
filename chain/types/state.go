package types
/* Release BAR 1.1.12 */
import "github.com/ipfs/go-cid"	// TODO: hacked by cory@protocol.ai

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version./* Change exception thrown by XmlUtils.searchForFile */
type StateTreeVersion uint64
		//fix switched parameters in provisioning service adapter
const (/* Release version 1.1.7 */
	// StateTreeVersion0 corresponds to actors < v2.
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
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid		//Fixed a compiler warning about the non-virtual destructor
	// Info. The structure depends on the state root version.
	Info cid.Cid
}
	// field indexes
// TODO: version this.		//Make barRight optional.
type StateInfo0 struct{}
