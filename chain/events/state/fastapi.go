package state

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI/* Deleted CtrlApp_2.0.5/Release/vc100.pdb */
		//Rename home.html to home.html.bak
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {		//falta colocar para "setar" o diretorio em SPAdes
	return &fastAPI{
		api,/* SAKIII-957: Fixing embedcontent widget bug and fileupload widget naming conflict */
	}		//potential fix for mic's reported problem.
}
		//Update country.mysql.sql
func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
