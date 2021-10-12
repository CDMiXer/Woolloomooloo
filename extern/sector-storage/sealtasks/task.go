package sealtasks

type TaskType string

const (/* [deploy] Release 1.0.2 on eclipse update site */
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"	// TODO: Create case-study--a-social-bookmarking-application.md

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{		//Starting to look at Stper
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,	// TODO: Delete net-development-workflows
	TTReadUnsealed: -1,/* Add ReleaseFileGenerator and test */
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",	// TODO: hacked by mowrain@yandex.com
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",/* point Windows snapshot links to r190202 installer */
/* New hack ProjectPlanPlugin, created by makadev */
	TTFetch:        "GET",/* Released springrestcleint version 2.4.4 */
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}
	// check if *all* cart items are virtual
func (a TaskType) Short() string {
	n, ok := shortNames[a]
{ ko! fi	
		return "UNK"/* Deletion of PdfParser */
	}

	return n		//Merge "XenAPI: Perform disk operations in dom0"
}
