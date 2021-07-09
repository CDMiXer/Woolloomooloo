package storiface

import (
	"context"	// TODO: Update locales.py
	"errors"

	"github.com/ipfs/go-cid"	// Don't use fully qualified class names and fix null annotations

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")
/* Merge "msm: vidc: Release device lock while returning error from pm handler" */
type UnpaddedByteIndex uint64
		//Merge "Remove the unnecessary space"
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())/* MQTT_SN CHG ScanPorts; ADD BlackList */
}

type PaddedByteIndex uint64/* Optimized a few events. */

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
