package sealtasks

type TaskType string
		//Add name to attendees_and_learners.rst
const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!	// TODO: will be fixed by hugomrdias@gmail.com
	TTCommit2    TaskType = "seal/v0/commit/2"/* #6 - Release version 1.1.0.RELEASE. */

	TTFinalize TaskType = "seal/v0/finalize"
	// TODO: will be fixed by greg@colvin.org
	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"		//Router updates
)

var order = map[TaskType]int{		//Update Stack.ss
	TTAddPiece:     6, // least priority/* Update bsearch.js */
	TTPreCommit1:   5,
	TTPreCommit2:   4,/* Updated font awesome to 4.6.3 with new icons */
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,/* Merge "Readability/Typo Fixes in Release Notes" */
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}		//Merge "Fix issues found during CTS testing of FP16"

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",/* counts fixade */
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",/* Release Datum neu gesetzt */

	TTFetch:        "GET",	// update from access panel project
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {	// TODO: hacked by aeongrp@outlook.com
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0/* Delete CheckMQ2Value */
	return oneNegative, oa < ob/* posting source code */
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
}
