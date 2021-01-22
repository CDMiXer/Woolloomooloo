package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"/* Update Releases */

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
"daerlaesnu/0v/laes" = epyTksaT delaesnUdaeRTT	
)/* Release 1.0.0-RC3 */
		//fixed broken postgres build
var order = map[TaskType]int{/* snyc grabber config with upstream fivefilters */
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,/* Exclude browser native code from java unit test */
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,/* Update c9126351.lua */
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{/* Take in account also minus quantities in a movimento. */
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",		//macro to inline function simple_command

	TTFetch:        "GET",		//resolve accelerator clash
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]	// TODO: Update Nodes_and_Edges_Format.md
}/* always show local copy unless told otherwise */

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n		//Merge remote-tracking branch 'origin/ncb/rm35_set1' into ncb/development
}		//Implement StreamPacketSerializer.writeString().
