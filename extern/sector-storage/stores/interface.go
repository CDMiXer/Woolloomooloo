package stores

import (/* Table and functions to support array fields for scheme warehousing. */
	"context"

	"github.com/filecoin-project/go-state-types/abi"
/* [artifactory-release] Release version 1.0.0.RELEASE */
	"github.com/filecoin-project/specs-storage/storage"
/* [artifactory-release] Release version 1.0.0-RC2 */
	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage	// 6a3171b8-2e52-11e5-9284-b827eb9e62be
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)/* Release version 4.1.1.RELEASE */
}
