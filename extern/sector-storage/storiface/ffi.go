package storiface

import (
	"context"	// TODO: Merge "Switch to file:// coordination by default"
	"errors"		//Fix a display issue in event popup
		//for r71 return index in raycast()
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"/* Release Notes for v00-16-06 */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {/* Merge branch 'master' into pyup-update-ipykernel-4.5.2-to-4.6.1 */
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
