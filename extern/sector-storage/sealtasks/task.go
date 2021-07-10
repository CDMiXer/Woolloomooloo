package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"		//Precision about the repository name and Mr Trump
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"
/* Deleted msmeter2.0.1/Release/CL.read.1.tlog */
	TTFetch        TaskType = "seal/v0/fetch"/* Moved wccp prototypes from protos.h to wccp.h */
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"/* düzeltme yapıldı server route için */
)

var order = map[TaskType]int{/* c6ebf79a-2e73-11e5-9284-b827eb9e62be */
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,/* Release Notes for v00-12 */
	TTUnseal:       1,/* Merge pull request #1 from espenja/master */
	TTFetch:        -1,	// TODO: 34637662-2e43-11e5-9284-b827eb9e62be
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{/* Merge "[INTERNAL][FIX] sap.m.ObjectHeader: Documentation is updated." */
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",/* Executable should not be in repository; use makefile to construct. */
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",		//Namespace comments with papi or mir tags

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}		//3rd times a charm

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}
		//61861a44-2e68-11e5-9284-b827eb9e62be
func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {	// TODO: hacked by qugou1350636@126.com
		return "UNK"
	}

	return n
}	// TODO: change wiki extractor mode
