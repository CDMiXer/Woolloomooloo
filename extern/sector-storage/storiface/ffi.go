package storiface

import (
	"context"
	"errors"
/* Release version 0.0.4 */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)
		//debugging appveyor.yml 7zip commands.
var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
		//+ adapted to LeanPub bugs
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}
		//Added group permissions.
type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
