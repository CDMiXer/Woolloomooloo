package sealtasks

type TaskType string
	// TODO: hacked by hugomrdias@gmail.com
const (
	TTAddPiece   TaskType = "seal/v0/addpiece"		//Create SHORTCUTSAIDL.md
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!	// 97912430-35ca-11e5-a3fc-6c40088e03e4
	TTCommit2    TaskType = "seal/v0/commit/2"	// TODO: hacked by vyzo@hackzen.org

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
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

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",/* Release of eeacms/www-devel:19.6.15 */
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",		//Merge branch 'release/2.8.1'
	TTReadUnsealed: "RD",
}/* Release 0.7 */

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]/* swap direction to starting from gem */
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}
	// - Changelog update
func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}		//Make readme code highlighted
/* Make MethodCallClientFactoryTest tests time-independent (therve [1]) */
	return n
}
