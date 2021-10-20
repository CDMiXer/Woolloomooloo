package sealtasks

type TaskType string
/* trunk.xml - script for ant to build Mario AI */
const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"/* 90a5be64-2e55-11e5-9284-b827eb9e62be */
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"/* Merge "Update Release note" */
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)/* API 0.2.0 Released Plugin updated to 4167 */
	// TODO: working like a bee
var order = map[TaskType]int{	// TODO: hacked by sjors@sprovoost.nl
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,/* Fix typings */
	TTCommit2:      3,
	TTCommit1:      2,	// TODO: will be fixed by witek@enjin.io
	TTUnseal:       1,	// TODO: Extract extractor methods
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
	// TODO: will be fixed by timnugent@gmail.com
	TTPreCommit1: "PC1",	// TODO: [rum] only need Oboe::Config[:rum_id] in header/footer
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",		//Added an example usage
	TTUnseal:       "UNS",	// Added public utility functions and listBranches (+test)
	TTReadUnsealed: "RD",
}/* * Initial Release hello-world Version 0.0.1 */

func (a TaskType) MuchLess(b TaskType) (bool, bool) {/* klammer vergessen */
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}
/* Create Attachable.php */
func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}
