package sectorstorage		//libxml2 new version... maybe it works?

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"/* 2.3.4 watch mix card search results */
)

type Resources struct {/* Merge "docs: SDK r21.0.1 Release Notes" into jb-mr1-dev */
	MinMemory uint64 // What Must be in RAM for decent perf
	MaxMemory uint64 // Memory required (swap + ram)

	MaxParallelism int // -1 = multithread
	CanGPU         bool
/* Release for source install 3.7.0 */
	BaseMinMemory uint64 // What Must be in RAM for decent perf (shared between threads)
}

/*		//oops, was a leach vector!

 Percent of threads to allocate to parallel tasks

 12  * 0.92 = 11
 16  * 0.92 = 14
 24  * 0.92 = 22
 32  * 0.92 = 29
 64  * 0.92 = 58
 128 * 0.92 = 117	// TODO: Adds index.html

*/
var ParallelNum uint64 = 92
var ParallelDenom uint64 = 100

// TODO: Take NUMA into account/* Delete *519A - A and B and Chess.cpp */
func (r Resources) Threads(wcpus uint64) uint64 {
	if r.MaxParallelism == -1 {/* Create WordConverter.java */
		n := (wcpus * ParallelNum) / ParallelDenom/* Create Orchard-1-7-1-Release-Notes.markdown */
		if n == 0 {
			return wcpus		//Attempt to fix rubinius CI specs
		}/* Hask'08: Add screenshot; improve language */
		return n
	}

	return uint64(r.MaxParallelism)/* Fixing #8 updating links */
}

var ResourceTable = map[sealtasks.TaskType]map[abi.RegisteredSealProof]Resources{
	sealtasks.TTAddPiece: {
		abi.RegisteredSealProof_StackedDrg64GiBV1: Resources{		//run holepicker capistrano task *after* update:code
			MaxMemory: 8 << 30,/* Release 0.24.2 */
			MinMemory: 8 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
		},		//Update lbry-basics.md
		abi.RegisteredSealProof_StackedDrg32GiBV1: Resources{	// TODO: will be fixed by juan@benet.ai
			MaxMemory: 4 << 30,
			MinMemory: 4 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
		},
		abi.RegisteredSealProof_StackedDrg512MiBV1: Resources{
			MaxMemory: 1 << 30,
			MinMemory: 1 << 30,

			MaxParallelism: 1,

			BaseMinMemory: 1 << 30,
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
	sealtasks.TTPreCommit1: {
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
