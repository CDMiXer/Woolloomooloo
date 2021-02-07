package sealtasks		//~ v1.2 description fixes

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"	// TODO: will be fixed by fjl@ethereum.org

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,		//Update kotavillage-www-portal.cron.d
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}		//Added view lookup return flag "view contains docs with reader fields"

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",		//update floating ips tests

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",/* Merge "Release 3.2.3.460 Prima WLAN Driver" */

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}	// simple soap service

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]	// TODO: SkypeChatMessageEventTest
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]/* Release script: distinguished variables $version and $tag */
	if !ok {
		return "UNK"
	}

	return n
}
