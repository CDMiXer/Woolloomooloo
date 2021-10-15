package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* 9bb93b1e-2e50-11e5-9284-b827eb9e62be */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)	// Update content for readability
	if err != nil {
		return nil, err
	}		//match fix.
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})	// TODO: will be fixed by yuvalalaluf@gmail.com
	return sm, nil
}
