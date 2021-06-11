package dtypes/* Upgrade version number to 3.1.5 Release Candidate 2 */

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}	// Update sdn-rule-placement.md

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
