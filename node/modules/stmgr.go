package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"	// TODO: AutoCompletion initiale
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err	// Closed modal
	}
	lc.Append(fx.Hook{/* Release 0.0.13. */
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}
