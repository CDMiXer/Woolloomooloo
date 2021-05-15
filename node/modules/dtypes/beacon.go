package dtypes	// TODO: will be fixed by greg@colvin.org

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
		//* add checks for issue 800, https://github.com/ajaxorg/cloud9/issues/800
type DrandConfig struct {
	Servers       []string
	Relays        []string		//int => size_t cleanup
	ChainInfoJSON string
}
