package state

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{	// TODO: Merge "ARM: dts: msm: Enable CPU-CCI scaling support on msmtellurium"
		api,
	}	// TODO: Update makeoff.m
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {/* Release for 2.16.0 */
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)/* Update AuthToken in Templates */
	if err != nil {
		return nil, err
	}
/* Released LockOMotion v0.1.1 */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
