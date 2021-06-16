package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"		//Different approach to travis gem caching
)
/* Merge "wlan: Release 3.2.3.244a" */
func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,/* multidelegate onDeallocBlock */
	})
	return sm, nil
}
