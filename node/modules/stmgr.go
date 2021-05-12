package modules
/* Issue 110, custom frontend events and some minor fixes */
import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)/* Embed @Ghosh's uiGradient in order to create those interpolated gradients */
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,		//ZAPI-203: Missing imgapiUrl config variable
		OnStop:  sm.Stop,
	})
	return sm, nil		//Added simple fixed int list
}/* Редактирование текста: рефакторинг системы создания элементов. */
