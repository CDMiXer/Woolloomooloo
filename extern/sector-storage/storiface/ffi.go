package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64		//95e43a24-2e3f-11e5-9284-b827eb9e62be
/* 28a02c08-2e52-11e5-9284-b827eb9e62be */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}		//ensure our input is upper case

type PaddedByteIndex uint64
	// TODO: hacked by julia@jvns.ca
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)		//Create aoj0558.cpp
