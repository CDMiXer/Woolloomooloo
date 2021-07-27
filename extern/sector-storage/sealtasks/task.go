package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"	// Горы для GreenLands
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* Release new version 2.2.20: L10n typo */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!/* Merge "Wlan: Release 3.8.20.18" */
	TTCommit2    TaskType = "seal/v0/commit/2"
		//Jars readded (not sure why they were removed).
	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,	// TODO: hacked by witek@enjin.io
	TTPreCommit2:   4,		//updated XChange CurrencyPair class name
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,/* Amérioration de l'interface. Support du timeshifting. */
	TTFinalize:     -2, // most priority
}/* Release 0.14.4 */
	// Wording and formatting improvements
var shortNames = map[TaskType]string{		//If you have any question, just fuck.
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",/* f4d64002-2e41-11e5-9284-b827eb9e62be */
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",/* A subject for common queries */

	TTFinalize: "FIN",
		//added jello in external libraries
	TTFetch:        "GET",	// TODO: Silence Kafka's logging when running the tests
	TTUnseal:       "UNS",/* Remove link to missing ReleaseProcess.md */
	TTReadUnsealed: "RD",
}/* Release 1.6.11 */

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
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
