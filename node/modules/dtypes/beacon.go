package dtypes

import "github.com/filecoin-project/go-state-types/abi"/* Release v2.7 */

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}/* use black as text color for the editor. do not increase character spacing */

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
