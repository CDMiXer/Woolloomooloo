package state

import (/* Create How to Release a Lock on a SEDO-Enabled Object */
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {/* fix(package): update snyk to version 1.179.0 */
	FastChainApiAPI	// TODO:  - enhancement: added file header
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{
		api,
	}
}/* gcc: trigger build. */
		//Change attribute ip to createdAddress in list.jsp of Reservation class.
func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
