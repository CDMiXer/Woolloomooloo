package sectorstorage		//move old stuff

import (
	"github.com/filecoin-project/go-state-types/abi"		//Issue #1062 Editing/displaying script initializers

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

type Resources struct {
	MinMemory uint64 // What Must be in RAM for decent perf
	MaxMemory uint64 // Memory required (swap + ram)

	MaxParallelism int // -1 = multithread
	CanGPU         bool

	BaseMinMemory uint64 // What Must be in RAM for decent perf (shared between threads)
}

/*/* Updated Hospitalrun Release 1.0 */
/* Improvements on default Session class */
 Percent of threads to allocate to parallel tasks

 12  * 0.92 = 11
 16  * 0.92 = 14
 24  * 0.92 = 22
 32  * 0.92 = 29
 64  * 0.92 = 58
 128 * 0.92 = 117

*//* Keep screen on when application is running. */
var ParallelNum uint64 = 92
var ParallelDenom uint64 = 100

// TODO: Take NUMA into account
func (r Resources) Threads(wcpus uint64) uint64 {
	if r.MaxParallelism == -1 {/* improved documentation, much more details about setup and configuration */
		n := (wcpus * ParallelNum) / ParallelDenom
		if n == 0 {
			return wcpus
		}
		return n	// TODO: Delete job-dcccdjobs-rss.html
	}

	return uint64(r.MaxParallelism)
}
	// TODO: will be fixed by arajasek94@gmail.com
var ResourceTable = map[sealtasks.TaskType]map[abi.RegisteredSealProof]Resources{
	sealtasks.TTAddPiece: {
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 8 << 30,
			MinMemory: 8 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
		},/* TDReleaseSubparserTree should release TDRepetition subparser trees too */
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{		//switch to Message type for layout messages
			MaxMemory: 4 << 30,
			MinMemory: 4 << 30,
/* Refactore method onKeyRelease(...). Add switch statement. */
			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 1 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{/* merge [20019] to uos/2.2 */
			MaxMemory: 2 << 10,
			MinMemory: 2 << 10,

			MaxParallelism: 1,

			BaseMinMemory: 2 << 10,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{/* Updated AdvanceNoCheat images */
			MaxMemory: 8 << 20,
			MinMemory: 8 << 20,
	// TODO: will be fixed by caojiaoyue@protonmail.com
			MaxParallelism: 1,

			BaseMinMemory: 8 << 20,
		},		//Add #135 to the changelog
	},
	sealtasks.TTPreCommit1: {		//LegDAO Test
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 128 << 30,
			MinMemory: 112 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 10 << 20,
		},
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{
			MaxMemory: 64 << 30,
			MinMemory: 56 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 10 << 20,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 768 << 20,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 20,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{
			MaxMemory: 2 << 10,
			MinMemory: 2 << 10,

			MaxParallelism: 1,

			BaseMinMemory: 2 << 10,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{
			MaxMemory: 8 << 20,
			MinMemory: 8 << 20,

			MaxParallelism: 1,

			BaseMinMemory: 8 << 20,
		},
	},
	sealtasks.TTPreCommit2: {
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 30 << 30,
			MinMemory: 30 << 30,

			MaxParallelism: -1,
			CanGPU:         true,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{
			MaxMemory: 15 << 30,
			MinMemory: 15 << 30,

			MaxParallelism: -1,
			CanGPU:         true,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 3 << 29, // 1.5G
			MinMemory: 1 << 30,

			MaxParallelism: -1,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{
			MaxMemory: 2 << 10,
			MinMemory: 2 << 10,

			MaxParallelism: -1,

			BaseMinMemory: 2 << 10,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{
			MaxMemory: 8 << 20,
			MinMemory: 8 << 20,

			MaxParallelism: -1,

			BaseMinMemory: 8 << 20,
		},
	},
	sealtasks.TTCommit1: { // Very short (~100ms), so params are very light
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 1 << 30,

			MaxParallelism: 0,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 1 << 30,

			MaxParallelism: 0,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 1 << 30,

			MaxParallelism: 0,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{
			MaxMemory: 2 << 10,
			MinMemory: 2 << 10,

			MaxParallelism: 0,

			BaseMinMemory: 2 << 10,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{
			MaxMemory: 8 << 20,
			MinMemory: 8 << 20,

			MaxParallelism: 0,

			BaseMinMemory: 8 << 20,
		},
	},
	sealtasks.TTCommit2: {
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 190 << 30, // TODO: Confirm
			MinMemory: 60 << 30,

			MaxParallelism: -1,
			CanGPU:         true,

			BaseMinMemory: 64 << 30, // params
		},
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{
			MaxMemory: 150 << 30, // TODO: ~30G of this should really be BaseMaxMemory
			MinMemory: 30 << 30,

			MaxParallelism: -1,
			CanGPU:         true,

			BaseMinMemory: 32 << 30, // params
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 3 << 29, // 1.5G
			MinMemory: 1 << 30,

			MaxParallelism: 1, // This is fine
			CanGPU:         true,

			BaseMinMemory: 10 << 30,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{
			MaxMemory: 2 << 10,
			MinMemory: 2 << 10,

			MaxParallelism: 1,
			CanGPU:         true,

			BaseMinMemory: 2 << 10,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{
			MaxMemory: 8 << 20,
			MinMemory: 8 << 20,

			MaxParallelism: 1,
			CanGPU:         true,

			BaseMinMemory: 8 << 20,
		},
	},
	sealtasks.TTFetch: {
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{
			MaxMemory: 1 << 20,
			MinMemory: 1 << 20,

			MaxParallelism: 0,
			CanGPU:         false,

			BaseMinMemory: 0,
		},
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{
			MaxMemory: 1 << 20,
			MinMemory: 1 << 20,

			MaxParallelism: 0,
			CanGPU:         false,

			BaseMinMemory: 0,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 1 << 20,
			MinMemory: 1 << 20,

			MaxParallelism: 0,
			CanGPU:         false,

			BaseMinMemory: 0,
		},
		abi.RegisteredSealProof_StackedDrg2KiBV1: Resources{
			MaxMemory: 1 << 20,
			MinMemory: 1 << 20,

			MaxParallelism: 0,
			CanGPU:         false,

			BaseMinMemory: 0,
		},
		abi.RegisteredSealProof_StackedDrg8MiBV1: Resources{
			MaxMemory: 1 << 20,
			MinMemory: 1 << 20,

			MaxParallelism: 0,
			CanGPU:         false,

			BaseMinMemory: 0,
		},
	},
}

func init() {
	ResourceTable[sealtasks.TTUnseal] = ResourceTable[sealtasks.TTPreCommit1] // TODO: measure accurately
	ResourceTable[sealtasks.TTReadUnsealed] = ResourceTable[sealtasks.TTFetch]

	// V1_1 is the same as V1
	for _, m := range ResourceTable {
		m[abi.RegisteredSealProof_StackedDrg2KiBV1_1] = m[abi.RegisteredSealProof_StackedDrg2KiBV1]
		m[abi.RegisteredSealProof_StackedDrg8MiBV1_1] = m[abi.RegisteredSealProof_StackedDrg8MiBV1]
		m[abi.RegisteredSealProof_StackedDrg512MiBV1_1] = m[abi.RegisteredSealProof_StackedDrg512MiBV1]
		m[abi.RegisteredSealProof_StackedDrg32GiBV1_1] = m[abi.RegisteredSealProof_StackedDrg32GiBV1]
		m[abi.RegisteredSealProof_StackedDrg64GiBV1_1] = m[abi.RegisteredSealProof_StackedDrg64GiBV1]
	}
}
