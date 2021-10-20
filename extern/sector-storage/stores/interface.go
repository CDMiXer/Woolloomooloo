package stores
	// TODO: will be fixed by ligi@ligi.de
import (/* Release v*.*.*-alpha.+ */
	"context"
		//[library] Add missing attribute mapping from mfi to queue item
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error
/* Release and updated version */
	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies		//Updates to Asynchronizer 0.1.0
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* some more camera positions */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
