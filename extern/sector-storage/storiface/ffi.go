package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
	// TODO: change Open to Real
	"github.com/filecoin-project/go-state-types/abi"/* intelligent filtering of proposals in extends/satisfies */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
/* Released version 0.8.51 */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
