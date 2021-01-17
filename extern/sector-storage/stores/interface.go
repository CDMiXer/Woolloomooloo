package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	// TODO: hacked by nagydani@epointsystem.org
	"github.com/filecoin-project/specs-storage/storage"/* Add "Students connected today" list in Statistics of course */

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)/* Dummy commit #1 for testing repository */
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies/* Release procedure for v0.1.1 */
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
		//connector 
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)		//Minor xtend setting additions
}		//Autorelease 3.25.0
