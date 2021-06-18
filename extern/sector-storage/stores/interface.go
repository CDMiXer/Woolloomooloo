package stores	// TODO: hacked by steven@stebalien.com
		//Add logger to PublicIdModel and getAll query
import (
	"context"	// TODO: hacked by brosner@gmail.com

	"github.com/filecoin-project/go-state-types/abi"
	// TODO: hacked by cory@protocol.ai
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies		//Update render-json-to-html.md
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage	// TODO: will be fixed by caojiaoyue@protonmail.com
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error	// Cleaner command line args

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)	// Update MinkContext to not escape locator to named
}
