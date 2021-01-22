package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
/* ass setReleaseDOM to false so spring doesnt change the message  */
	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")
/*  Arch: import ifc, fix typo in material colors */
type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())/* Release note updated */
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
