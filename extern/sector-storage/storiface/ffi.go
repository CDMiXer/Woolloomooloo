package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")
	// TODO: will be fixed by steven@stebalien.com
type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {/* Check that current stack frame exists in stepForward. */
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
