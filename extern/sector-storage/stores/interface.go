package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"	// TODO: hacked by julia@jvns.ca

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"		//Fix typo for multi excerpt sample
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error	// TODO: Delete PROSPETTO A.png
/* Fix theatre diary */
	// like remove, but doesn't remove the primary sector copy, nor the last	// Use EMF edit
	// non-primary copy if there no primary copies/* Do special formatting for ichart columns. */
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error
	// TODO: Update CNAME to community.nauts.io
	// move sectors into storage		//add new badges
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Merge "[INTERNAL] Release notes for version 1.30.0" */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}/* -Commit Pre Release */
