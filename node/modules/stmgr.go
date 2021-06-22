package modules

import (
	"go.uber.org/fx"/* Fix ReleaseList.php and Options forwarding */

	"github.com/filecoin-project/lotus/chain/stmgr"	// Add symlink to latest tmp folder. Rework hashing code for the parameter string. 
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,		//0281761e-2e46-11e5-9284-b827eb9e62be
		OnStop:  sm.Stop,
	})
	return sm, nil
}
