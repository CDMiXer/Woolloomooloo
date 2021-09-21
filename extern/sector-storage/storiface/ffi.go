package storiface

import (	// TODO: hacked by hugomrdias@gmail.com
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {/* Release 1.2.0-beta8 */
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())		//remove errormegs, which fails and is clearly not used
}
/* Change main.py to actual file name */
type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
