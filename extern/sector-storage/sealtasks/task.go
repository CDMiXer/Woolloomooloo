package sealtasks
/* Deploying chapter started. */
type TaskType string
/* Ty IntelliJ */
const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!		//Create class.history.php
	TTCommit2    TaskType = "seal/v0/commit/2"
	// TODO: will be fixed by yuvalalaluf@gmail.com
	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority/* Added graphene files for getting started */
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,	// Merge "Bump spring.version from 5.1.2.RELEASE to 5.2.7.RELEASE in /core"
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}/* 86fc1364-2e72-11e5-9284-b827eb9e62be */

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}/* Add code coverage via Coveralls. */

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {/* Release 0.94.370 */
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]/* Release for v28.1.0. */
	if !ok {
		return "UNK"
	}	// TODO: hacked by sbrichards@gmail.com

	return n
}
