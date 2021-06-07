package dtypes

import "github.com/filecoin-project/go-state-types/abi"/* 074df4a4-2e75-11e5-9284-b827eb9e62be */

type DrandSchedule []DrandPoint

type DrandPoint struct {/* commenting in various renders */
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string		//removed hint
}
