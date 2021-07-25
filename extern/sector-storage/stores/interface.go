package stores	// TODO: #4: Some changes, but not there yet...

import (		//Ability for slow upgrade scripts to be manually run after upgrade.
	"context"	// TODO: will be fixed by souzau@yandex.com

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies		//dependencies for setting up the web server
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error/* install unique constraints on target model when remixing 1:1 */

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Release '0.2~ppa6~loms~lucid'. */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)	// TODO: Create optional constructor
}
