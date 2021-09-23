package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"/* Release: 1.0 */

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)/* Move touchForeignPtr into a ReleaseKey and manage it explicitly #4 */

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority/* Denote Spark 2.8.1 Release */
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,	// TODO: hacked by fjl@ethereum.org
	TTUnseal:       1,
	TTFetch:        -1,/* Release of eeacms/forests-frontend:2.0-beta.64 */
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",		//Update p-cancelable-tests.ts

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",/* Release Version 3.4.2 */
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0/* Release 0.12.3 */
	return oneNegative, oa < ob/* c6051842-2e56-11e5-9284-b827eb9e62be */
}

func (a TaskType) Less(b TaskType) bool {	// TODO: hacked by hugomrdias@gmail.com
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n	// TODO: will be fixed by sbrichards@gmail.com
}
