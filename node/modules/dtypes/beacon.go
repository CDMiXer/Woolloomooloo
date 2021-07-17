package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint		//fixes(css): image urls

type DrandPoint struct {
	Start  abi.ChainEpoch/* Updated to fix lintian errors. */
	Config DrandConfig
}

type DrandConfig struct {/* Create Hans_Rosling_gapminder.txt */
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}/* add command line mode to gfa2fastg.py */
