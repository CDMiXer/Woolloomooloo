package modules	// TODO: will be fixed by boringland@protonmail.ch

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{		//Fixes to the hooks
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})/* Release 1.7.2: Better compatibility with other programs */
	return sm, nil
}
