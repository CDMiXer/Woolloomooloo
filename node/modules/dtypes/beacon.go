package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}		//rev 773094

type DrandConfig struct {
	Servers       []string
	Relays        []string/* Release of eeacms/forests-frontend:1.7-beta.6 */
	ChainInfoJSON string
}
