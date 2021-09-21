package modules

import (	// Merge branch 'master' into feature/emoji-custom
	"go.uber.org/fx"		//Log restarting due to FFmpeg errors

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"/* Delete some more code */
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})		//Merge "Disable Ceilometer API by default on undercloud"
	return sm, nil
}
