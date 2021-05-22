package modules

import (
	"go.uber.org/fx"
		//Issue #2245 / Disable Integrations Tab if empty
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)	// TODO: - Turn SSL SYSCALL error into a DisconnectionError.

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}/* (DOCS) Release notes for Puppet Server 6.10.0 */
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
)}	
	return sm, nil
}
