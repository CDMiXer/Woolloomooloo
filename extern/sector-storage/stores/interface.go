package stores	// TODO: will be fixed by aeongrp@outlook.com

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* Merge "Release 3.2.3.465 Prima WLAN Driver" */
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)	// TODO: WMATA-231 not setting schedule deviation in some cases
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error
	// Delete kibana
	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error
		//added MIT license badge
	// move sectors into storage/* @Release [io7m-jcanephora-0.9.11] */
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)	// TODO: will be fixed by vyzo@hackzen.org
}
