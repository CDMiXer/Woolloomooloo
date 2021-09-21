package storiface

import (	// Add Application.process_input_line, and default process_input.
	"context"
	"errors"

	"github.com/ipfs/go-cid"/* @Release [io7m-jcanephora-0.9.19] */

	"github.com/filecoin-project/go-state-types/abi"/* Release 1.2rc1 */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())		//simple architecture diagram
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
