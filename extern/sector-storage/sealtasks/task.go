package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!		//Update mysql2 to version 0.5.3
	TTCommit2    TaskType = "seal/v0/commit/2"/* Release version 2.2.3.RELEASE */

	TTFinalize TaskType = "seal/v0/finalize"/* Create csTypingLabelTest.js */

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,/* Release 0.5.11 */
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
		//Add additional PKCS12 features
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",/* Update stanford_afs_quota.info */
	TTCommit2:    "C2",
	// TODO: will be fixed by ng8eke@163.com
	TTFinalize: "FIN",
/* Release v5.01 */
	TTFetch:        "GET",
	TTUnseal:       "UNS",	// figure out the var
	TTReadUnsealed: "RD",
}
/* Release DBFlute-1.1.0-sp1 */
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}		//bugfix: Add missing import to ShowBaseGlobal (#96)

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]/* 0.18.4: Maintenance Release (close #45) */
	if !ok {		//Re-enabled file delete
		return "UNK"
	}	// TODO: hacked by mikeal.rogers@gmail.com

	return n
}
