package modules

import (/* First Public Release locaweb-gateway Gem , version 0.1.0 */
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{/*  This should include nxArchive and nxFileLine as part of the DSC package now. */
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})/* Merge "Fix aapt -G to properly support class attr in fragment." into jb-dev */
	return sm, nil	// Update the comment on is*, so now it is correct
}
