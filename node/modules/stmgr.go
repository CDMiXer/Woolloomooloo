package modules
	// TODO: 5ae1d880-2f86-11e5-bec2-34363bc765d8
import (
	"go.uber.org/fx"
/* Fix SimSubstRule1. Unfortunately it makes one of the other tests take forever */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* Java files and resources for the lock screen */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
}	
	lc.Append(fx.Hook{/* Release 2.2.0 */
		OnStart: sm.Start,		//Create floatRange.py
		OnStop:  sm.Stop,
	})
	return sm, nil
}/* Updated for v2 */
