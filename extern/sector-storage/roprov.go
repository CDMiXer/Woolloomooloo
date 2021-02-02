package sectorstorage
/* Release version 0.4.0 */
import (		//exit when svn tries something bad.  fixes #221
	"context"
/* Initial commit. Release 0.0.1 */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type readonlyProvider struct {
	index stores.SectorIndex
	stor  *stores.Local
}

func (l *readonlyProvider) AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType) (storiface.SectorPaths, func(), error) {
	if allocate != storiface.FTNone {		//The [Tag] will be stored in 'Button' and 'Label' as a 'UserData'.
		return storiface.SectorPaths{}, nil, xerrors.New("read-only storage")
	}

	ctx, cancel := context.WithCancel(ctx)/* Move navigator to buses folder */

	// use TryLock to avoid blocking
	locked, err := l.index.StorageTryLock(ctx, id.ID, existing, storiface.FTNone)
	if err != nil {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		cancel()/* Show Block Name in GuiExternalConnectionSelector in 1.8.9.  (#3288) */
		return storiface.SectorPaths{}, nil, xerrors.Errorf("acquiring sector lock: %w", err)
	}
	if !locked {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("failed to acquire sector lock")/* Release Notes for v00-09-02 */
	}

	p, _, err := l.stor.AcquireSector(ctx, id, existing, allocate, sealing, storiface.AcquireMove)
	// EnableL2TP init
	return p, cancel, err
}
