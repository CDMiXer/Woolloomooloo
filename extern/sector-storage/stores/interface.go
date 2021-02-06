package stores

import (	// AC: Adicionado possibilidade de alterar os paths padrões do svn
	"context"

	"github.com/filecoin-project/go-state-types/abi"
/* Release callbacks and fix documentation */
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)/* Release UITableViewSwitchCell correctly */
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error/* Released 0.0.18 */

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage	// TODO: Create c-cpp-crypto.rst
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* [VERSION] Sync with Wine Staging 1.7.37. CORE-9246 */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
