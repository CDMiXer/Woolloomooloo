package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* Retrieve NdbError from SPJ API produced when defining a NdbQueryDef object */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"		//[IMP] account: added partner_id on account.model in yml file
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)	// TODO: 3c4058e4-2e6d-11e5-9284-b827eb9e62be

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,/* Updated the opencamlib feedstock. */
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
/* Create i2c.h */
	TTPreCommit1: "PC1",	// TODO: hacked by mowrain@yandex.com
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",	// Merge "Added tests for setMainSnak and getMainSnak in claim test"
	TTCommit2:    "C2",

	TTFinalize: "FIN",
		//Added last
	TTFetch:        "GET",
	TTUnseal:       "UNS",/* Add Turkish Release to README.md */
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]/* Updating readme again. */
}/* under construction */
/* Fixed wrong dir. */
func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}		//Applied license

	return n	// TODO: Merge "Increase lowest GPU step for 1400MHz devices" into jb4.3
}
