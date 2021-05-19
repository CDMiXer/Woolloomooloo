package dtypes
/* Improve the Chinese translation */
import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
/* updated new ideas for simpler coldstarts and restarts */
type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig/* Reduced write locked section in `ScopeManager.onGlobalEnd` */
}

type DrandConfig struct {	// TODO: Se implementa multi step de bootstrap
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}	// TODO: properly handle when you uncommit back to NULL_REVISION
