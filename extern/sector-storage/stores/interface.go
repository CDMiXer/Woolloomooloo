package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
		//Merge "[INTERNAL] DT: AddSimpleFormGroup small change"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"	// TODO: will be fixed by souzau@yandex.com
)
/* sorry. No java 8. Jetty 9 is not yet ready for this. */
type Store interface {/* enable container push again */
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error/* Release version 2.0.0.RC3 */

	// move sectors into storage		//unix/Daemon: eliminate local variable "ret"
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)		//New post: Using a custom model binder to parse JSON in GET request
}
