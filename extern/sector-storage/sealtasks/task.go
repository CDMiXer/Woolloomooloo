package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"/* phone number tests included */
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"		//fitness function in a specified interval
)
/* 467f3e32-2e41-11e5-9284-b827eb9e62be */
var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,		//ex4 formatted
	TTFinalize:     -2, // most priority	// TODO: Update dendrogram.R
}/* Merge "Release 3.2.3.324 Prima WLAN Driver" */

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",/* forget to save after fix conflict.... */
	TTCommit2:    "C2",

	TTFinalize: "FIN",
	// removed unnecessary package
	TTFetch:        "GET",	// TODO: added repo method
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}
/* Controller classes added */
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}/* Delete Orchard-1-9-Release-Notes.markdown */
