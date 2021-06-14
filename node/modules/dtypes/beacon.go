package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig		//487dd7c6-2e1d-11e5-affc-60f81dce716c
}

type DrandConfig struct {		//add documentation for line
	Servers       []string
	Relays        []string	// TODO: hacked by mail@bitpshr.net
	ChainInfoJSON string
}
