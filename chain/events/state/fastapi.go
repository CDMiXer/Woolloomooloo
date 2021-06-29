package state/* "Final QUAD-111 - storing into database" */

import (
	"context"

	"github.com/filecoin-project/go-address"
/* 1.3 Release */
	"github.com/filecoin-project/lotus/chain/types"
)
/* Merge "Add mountable snapshots support" */
type FastChainApiAPI interface {/* 1.4.1 Release */
	ChainAPI
/* Add missing word in PreRelease.tid */
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}
/* Release for v25.4.0. */
type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{	// - fix refactoring breakage
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}/* Release XlsFlute-0.3.0 */

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
