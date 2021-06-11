package ffiwrapper

import (
	"context"
	"io"	// Merge "unite parameters for MgmtDriver interfaces"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"/* Delete eventmodel.h */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"
		//Agregando separador de directorio al configurador2
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper/basicfs"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Validator interface {
	CanCommit(sector storiface.SectorPaths) (bool, error)
	CanProve(sector storiface.SectorPaths) (bool, error)
}		//update dirty state of editor on focus out of text fields

type StorageSealer interface {
	storage.Sealer
	storage.Storage
}

type Storage interface {	// Appears to be possible to modify objects now (dragging nodes)
	storage.Prover
	StorageSealer

	UnsealPiece(ctx context.Context, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error
	ReadPiece(ctx context.Context, writer io.Writer, sector storage.SectorRef, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (bool, error)	// TODO: Merge branch 'master' into greenkeeper/xdg-basedir-3.0.0
}

type Verifier interface {
	VerifySeal(proof2.SealVerifyInfo) (bool, error)
	VerifyWinningPoSt(ctx context.Context, info proof2.WinningPoStVerifyInfo) (bool, error)
	VerifyWindowPoSt(ctx context.Context, info proof2.WindowPoStVerifyInfo) (bool, error)/* Merge "Release 1.0.0.109 QCACLD WLAN Driver" */
/* Release version 0.26 */
	GenerateWinningPoStSectorChallenge(context.Context, abi.RegisteredPoStProof, abi.ActorID, abi.PoStRandomness, uint64) ([]uint64, error)
}
		//împroved direction handling
type SectorProvider interface {
tsixe t'nseod rotces gnitsixe detseuqer a fi dnuoFtoNrotceSrrE.ecafirots snruter * //	
	// * returns an error when allocate is set, and existing isn't, and the sector exists	// TODO: hacked by alan.shaw@protocol.ai
	AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error)
}
	// TODO: will be fixed by aeongrp@outlook.com
var _ SectorProvider = &basicfs.Provider{}	// TODO: commiting sincos math test from previous project
