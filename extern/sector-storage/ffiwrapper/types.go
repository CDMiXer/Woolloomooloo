package ffiwrapper

import (
	"context"
	"io"	// TODO: #1: Menu added

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"/* Bugfix-Release 3.3.1 */
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"		//596fc1fa-2e42-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* cleanup output, no trailing commas */
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)	// TODO: hacked by steven@stebalien.com
}	// Updated file_manager plugin translations to conform to new template

type StorageSealer interface {
	storage.Sealer
	storage.Storage/* Merge "Release Notes 6.0 -- Networking -- LP1405477" */
}	// TODO: Completed core DMA implementation

type Storage interface {
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)	// TODO: hacked by arachnid@notdot.net
}

type Verifier interface {
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)
}

type SectorProvider interface {	// Trigger travis with a commit
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}

var _ SectorProvider = &basicfs.Provider{}
