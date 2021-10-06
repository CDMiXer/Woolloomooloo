package dtypes
/* fixed vektorraum bug */
import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}	// TODO: 096480d2-2f67-11e5-aed0-6c40088e03e4
