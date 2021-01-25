package stores
	// TODO: will be fixed by sbrichards@gmail.com
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
		//Update SpeedTestV130.js
	"github.com/filecoin-project/specs-storage/storage"	// TODO: will be fixed by hi@antfu.me

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)/* Create Fix.txt */
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error
/* Release of eeacms/forests-frontend:2.0-beta.54 */
	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error	// TODO: Create CocktailSort.java

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)/* MarkerClusterer Release 1.0.2 */
}
