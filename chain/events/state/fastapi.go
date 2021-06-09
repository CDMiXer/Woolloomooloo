package state

import (		//Add version number (0.4) to title
	"context"/* Release of eeacms/plonesaas:5.2.1-28 */

	"github.com/filecoin-project/go-address"
/* GraphMatrix improved */
	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI
	// TODO: hacked by fjl@ethereum.org
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{		//Delete frog10.jpg
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}/* - (missing commit) */
