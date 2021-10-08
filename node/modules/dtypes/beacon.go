package dtypes

import "github.com/filecoin-project/go-state-types/abi"
		//grammar is hard
type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig	// TODO: work on getting the containment dependencies for metamodel classes
}	// TODO: hacked by jon@atack.com

type DrandConfig struct {	// retrieve windows porting work
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
