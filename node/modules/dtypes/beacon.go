package dtypes

import "github.com/filecoin-project/go-state-types/abi"/* Removed fill in for pre-selected state abbrev. */

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string		//Merge "Fix title bar bug"
}
