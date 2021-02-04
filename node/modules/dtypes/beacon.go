package dtypes

import "github.com/filecoin-project/go-state-types/abi"	// TODO: remove system.outs

type DrandSchedule []DrandPoint
	// TODO: Updated toString to look like case class toString
type DrandPoint struct {		//added JOSS badge
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}		//Install lcov on Ubuntu VMs.
