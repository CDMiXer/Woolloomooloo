package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)	// TODO: Committing an additional UnionExtract test.

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {		//preload fonts
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)/* complete logout coverage */
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,/* Next Release!!!! */
		OnStop:  sm.Stop,
	})
	return sm, nil
}
