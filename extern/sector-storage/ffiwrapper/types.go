package ffiwrapper
/* Merge "Fix typo in Release note" */
import (
	"context"/* add local passed */
	"io"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
	// TODO: Wait4GearGone command fixed
	"github.com/ipfs/go-cid"	// need to add version

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)	// TODO: JSHint code cleanup
	CanProve(sector storiface.SectorPaths) (bool, error)
}/* Release new version 2.2.10:  */

type StorageSealer interface {
	storage.Sealer
	storage.Storage
}

type Storage interface {		//Merge "Ensure httpd is not enabled by puppet on system boot"
	storage.Prover	// TODO: Create slidepuzzle.py
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)
}

type Verifier interface {		//Merge "n_smux: Fix compilation when CONFIG_N_SMUX is undefined" into msm-3.0
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)/* Create PabloMolina.md */
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)

	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)	// TODO: Update example/line/line.html
}

type SectorProvider interface {
	// * returns storiface.ErrSectorNotFound if a requested existing sector doesn't exist
	// * returns an error when allocate is set, and existing isn't, and the sector exists/* Merge "Update API" into master-nova */
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}

var _ SectorProvider = &basicfs.Provider{}
