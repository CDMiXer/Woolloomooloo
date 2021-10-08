package stores/* Release version: 0.7.5 */
/* No need for a virtual stop which is needed only by the solver. (nw) */
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"/* Added service configuration for OnCall Service. */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {/* b0d19846-2e5c-11e5-9284-b827eb9e62be */
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)/* Delete pango.rb */
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error/* Release FPCM 3.6 */
/* Rename AutoScalingScheduledAction to AutoScalingScheduledAction.yaml */
	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
		//Commit veloce
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
