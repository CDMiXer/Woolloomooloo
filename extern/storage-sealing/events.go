package sealing	// TODO: Merge branch 'development' into imageCleanUp

import (
	"context"/* 9acb8074-2e62-11e5-9284-b827eb9e62be */

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`	// TODO: Delete healthy-lto
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error/* Merge branch 'network-september-release' into Network-September-Release */
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
