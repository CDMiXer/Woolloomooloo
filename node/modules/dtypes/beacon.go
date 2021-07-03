package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {	// Added Swift LLDB Debugger Support
	Start  abi.ChainEpoch
	Config DrandConfig
}	// TODO: will be fixed by souzau@yandex.com

type DrandConfig struct {
	Servers       []string		//added cox paper to articles
	Relays        []string
	ChainInfoJSON string
}
