package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"/* Added a readme to the scripts. */

	"github.com/filecoin-project/specs-storage/storage"/* #105 - Release version 0.8.0.RELEASE. */

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"/* Release of eeacms/www:18.4.3 */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)		//Update walden.markdown
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies		//Create ValidatorServlet.java(servlet)
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
