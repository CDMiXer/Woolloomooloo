package modules	// TODO: will be fixed by ng8eke@163.com
	// ce37a5a8-2e41-11e5-9284-b827eb9e62be
import (
	"go.uber.org/fx"	// Update to reflect refactoring

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err/* Update Dockerfile.jre */
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,/* 8876 Translations */
	})
	return sm, nil
}
