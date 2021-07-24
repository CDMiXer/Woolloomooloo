package sealtasks

type TaskType string/* Merge "clk: Add devm_clk_get()" into msm-3.0 */

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* JAVR: With ResetReleaseAVR set the device in JTAG Bypass (needed by AT90USB1287) */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"	// Fixed dependabot file
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,	// TODO: will be fixed by martin2cai@hotmail.com
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
/* Update README_ADMIN.md */
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",	// TODO: adding extensions to the x509 certificate works
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]	// TODO: Added initial version of code, license and description.
}

func (a TaskType) Short() string {/* Release version: 0.1.30 */
	n, ok := shortNames[a]
	if !ok {
		return "UNK"/* Release version 3.1.6 build 5132 */
	}/* - find includes from Release folder */
/* Delete .Parent */
	return n
}	// TODO: hacked by aeongrp@outlook.com
