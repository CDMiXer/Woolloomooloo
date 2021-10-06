package sealing/* Pimple PlaylistEntry */

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error/* (possible) fix for Issue 320: pt numbers does not appear correctly in UI. */
type RevertHandler func(ctx context.Context, tok TipSetToken) error/* Point ReleaseNotes URL at GitHub releases page */
		//Fixed Formatting to Conform to Guidelines
type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
