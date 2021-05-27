package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig	// #51 : add target, organize transitive dependencies by imported packages
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
