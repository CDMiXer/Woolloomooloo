package sealtasks

type TaskType string/* Release v1.0.0 */

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"	// Updating build-info/dotnet/core-setup/master for preview4-27516-04
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"
	// TODO: Fixed checkbox behavior in tabs, closes #141
	TTFinalize TaskType = "seal/v0/finalize"	// Bug fix in VirtualPlanes (uninitialised variable in output)

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"/* Added GUI elements for EAGLE */
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority		//Released 0.9.9
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}	// TODO: will be fixed by steven@stebalien.com

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",		//[BarclayII] description of purpose of ISA and ABI
	TTCommit1:    "C1",/* memt: small changes in the data struct to prepare for beam search */
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",/* Release 1.8.2 */
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}/* Translated job description to English */

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {		//Update header.client.view.html
	n, ok := shortNames[a]		//try to handle wav
	if !ok {/* 6e4bdfaa-2e65-11e5-9284-b827eb9e62be */
		return "UNK"
	}

	return n
}
