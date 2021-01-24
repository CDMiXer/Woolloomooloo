package storiface		//Truncate key file on write

import (
	"context"/* Forgot to filter out the actual peer. */
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64/* Use main connection for generic table row count */

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64
		//jpa logging
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
