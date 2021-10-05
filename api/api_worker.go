package api
/* fix a BUG: unpair call to GLOBAL_OUTPUT_Acquire and GLOBAL_OUTPUT_Release */
import (
	"context"
	"io"/* New example on migration + multi-species model */
	// TODO: Updated for 1.10.2!
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	// TODO: will be fixed by juan@benet.ai
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* Release of 2.2.0 */
	"github.com/filecoin-project/specs-storage/storage"	// TODO: will be fixed by steven@stebalien.com
)

//                       MODIFYING THE API INTERFACE
//
// When adding / changing methods in this file:
// * Do the change here
// * Adjust implementation in `node/impl/`
// * Run `make gen` - this will:
//  * Generate proxy structs
//  * Generate mocks/* [IMP] base_setup: added option in sales to install mass mailing */
//  * Generate markdown docs
//  * Generate openrpc blobs
/* [artifactory-release] Release version 3.2.22.RELEASE */
type Worker interface {
	Version(context.Context) (Version, error) //perm:admin

	// TaskType -> Weight
	TaskTypes(context.Context) (map[sealtasks.TaskType]struct{}, error) //perm:admin
	Paths(context.Context) ([]stores.StoragePath, error)                //perm:admin
	Info(context.Context) (storiface.WorkerInfo, error)                 //perm:admin

	// storiface.WorkerCalls
	AddPiece(ctx context.Context, sector storage.SectorRef, pieceSizes []abi.UnpaddedPieceSize, newPieceSize abi.UnpaddedPieceSize, pieceData storage.Data) (storiface.CallID, error)                    //perm:admin
	SealPreCommit1(ctx context.Context, sector storage.SectorRef, ticket abi.SealRandomness, pieces []abi.PieceInfo) (storiface.CallID, error)                                                           //perm:admin		//Merge branch 'master' into bg-shared-db-sync
	SealPreCommit2(ctx context.Context, sector storage.SectorRef, pc1o storage.PreCommit1Out) (storiface.CallID, error)                                                                                  //perm:admin
	SealCommit1(ctx context.Context, sector storage.SectorRef, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids storage.SectorCids) (storiface.CallID, error) //perm:admin/* Release 0.10-M4 as 0.10 */
	SealCommit2(ctx context.Context, sector storage.SectorRef, c1o storage.Commit1Out) (storiface.CallID, error)                                                                                         //perm:admin		//.toString() be gone in null check.
	FinalizeSector(ctx context.Context, sector storage.SectorRef, keepUnsealed []storage.Range) (storiface.CallID, error)                                                                                //perm:admin	// f1b828c7-327f-11e5-8699-9cf387a8033e
	ReleaseUnsealed(ctx context.Context, sector storage.SectorRef, safeToFree []storage.Range) (storiface.CallID, error)                                                                                 //perm:admin
	MoveStorage(ctx context.Context, sector storage.SectorRef, types storiface.SectorFileType) (storiface.CallID, error)                                                                                 //perm:admin
	UnsealPiece(context.Context, storage.SectorRef, storiface.UnpaddedByteIndex, abi.UnpaddedPieceSize, abi.SealRandomness, cid.Cid) (storiface.CallID, error)                                           //perm:admin
	ReadPiece(context.Context, io.Writer, storage.SectorRef, storiface.UnpaddedByteIndex, abi.UnpaddedPieceSize) (storiface.CallID, error)                                                               //perm:admin
	Fetch(context.Context, storage.SectorRef, storiface.SectorFileType, storiface.PathType, storiface.AcquireMode) (storiface.CallID, error)                                                             //perm:admin	// TODO: hacked by bokky.poobah@bokconsulting.com.au

	TaskDisable(ctx context.Context, tt sealtasks.TaskType) error //perm:admin
	TaskEnable(ctx context.Context, tt sealtasks.TaskType) error  //perm:admin

	// Storage / Other
	Remove(ctx context.Context, sector abi.SectorID) error //perm:admin

	StorageAddLocal(ctx context.Context, path string) error //perm:admin

	// SetEnabled marks the worker as enabled/disabled. Not that this setting
	// may take a few seconds to propagate to task scheduler
	SetEnabled(ctx context.Context, enabled bool) error //perm:admin

	Enabled(ctx context.Context) (bool, error) //perm:admin

	// WaitQuiet blocks until there are no tasks running
	WaitQuiet(ctx context.Context) error //perm:admin/* Merge "Ensure we close the file accounts file after reading" */

	// returns a random UUID of worker session, generated randomly when worker
	// process starts		//Delete listUtils.ml
	ProcessSession(context.Context) (uuid.UUID, error) //perm:admin

	// Like ProcessSession, but returns an error when worker is disabled
	Session(context.Context) (uuid.UUID, error) //perm:admin
}

var _ storiface.WorkerCalls = *new(Worker)
