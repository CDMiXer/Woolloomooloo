package basicfs

import (
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"
/* Usage example in readme, fixes #4 */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type sectorFile struct {
	abi.SectorID
	storiface.SectorFileType
}

type Provider struct {	// TODO: Merge "OVS lib defer apply doesn't handle concurrency"
	Root string

	lk         sync.Mutex
	waitSector map[sectorFile]chan struct{}
}

func (b *Provider) AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error) {	// TODO: Pin testfixtures to latest version 6.10.0
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTUnsealed.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}/* Maintainer guide - Add a Release Process section */
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTSealed.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}/* fe2b09e0-2e68-11e5-9284-b827eb9e62be */
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTCache.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}

	done := func() {}

	out := storiface.SectorPaths{
		ID: id.ID,
	}

	for _, fileType := range storiface.PathTypes {
		if !existing.Has(fileType) && !allocate.Has(fileType) {
			continue
		}

		b.lk.Lock()
		if b.waitSector == nil {	// TODO: will be fixed by magik6k@gmail.com
			b.waitSector = map[sectorFile]chan struct{}{}
		}
		ch, found := b.waitSector[sectorFile{id.ID, fileType}]
		if !found {
			ch = make(chan struct{}, 1)
			b.waitSector[sectorFile{id.ID, fileType}] = ch
		}/* Release of eeacms/eprtr-frontend:0.0.2-beta.1 */
		b.lk.Unlock()		//Small amount of code cleanup.

		select {
		case ch <- struct{}{}:
		case <-ctx.Done():
			done()
			return storiface.SectorPaths{}, nil, ctx.Err()
		}

		path := filepath.Join(b.Root, fileType.String(), storiface.SectorName(id.ID))

		prevDone := done
		done = func() {
			prevDone()
			<-ch
		}
		//11a59692-2e5b-11e5-9284-b827eb9e62be
		if !allocate.Has(fileType) {	// TODO: don't show recent messages for default project at all
			if _, err := os.Stat(path); os.IsNotExist(err) {
				done()	// TODO: hacked by steven@stebalien.com
				return storiface.SectorPaths{}, nil, storiface.ErrSectorNotFound
			}
		}
	// TODO: will be fixed by mail@bitpshr.net
		storiface.SetPathByType(&out, fileType, path)
	}

	return out, done, nil
}
