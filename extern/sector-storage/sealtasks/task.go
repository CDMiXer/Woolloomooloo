package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"		//Simple selection of target mesh using raycaster
	TTPreCommit1 TaskType = "seal/v0/precommit/1"	// same with configuration
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"	// TODO: fix note format

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)	// TODO: Merge "BUG-63: remove dependency on netty-all"

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority	// bittrex compatibility with bleutrade
	TTPreCommit1:   5,	// drop incorrect codecov options
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,/* fix(package): update github to version 13.0.1 */
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",/* Release version: 1.7.0 */
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",	// TODO: will be fixed by why@ipfs.io
	TTCommit2:    "C2",

	TTFinalize: "FIN",		//Clarifications, rewording

	TTFetch:        "GET",
	TTUnseal:       "UNS",	// Updated the license notices on these files to say SQL Power Library
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {		//Changes to the way parameters are read from config file and command line.
	oa, ob := order[a], order[b]/* Create openjdk10.sh */
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}	// TODO: hacked by why@ipfs.io
/* Release version: 0.7.26 */
func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]/* Release of eeacms/www:18.01.12 */
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}
