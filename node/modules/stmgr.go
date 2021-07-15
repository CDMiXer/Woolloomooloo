package modules

import (
	"go.uber.org/fx"
/* Create outlook.com.md */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)/* Alle URLs in den Controller auf Named-Parameter umgestellt. */

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by josharian@gmail.com
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
)}	
	return sm, nil
}		//revert extra config namespace
