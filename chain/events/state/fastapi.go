package state
/* replaced emma with jacoco */
import (
	"context"/* Merge "Release 1.0.0.76 QCACLD WLAN Driver" */

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by sbrichards@gmail.com
)

type FastChainApiAPI interface {		//Improvements on FastaManipulatorServer
	ChainAPI	// TODO: hacked by alex.gaynor@gmail.com

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {		//Improve message error
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{/* Merge "Release 1.0.0.95 QCACLD WLAN Driver" */
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
