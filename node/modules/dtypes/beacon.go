package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint/* Update lazy-object-proxy from 1.4.1 to 1.4.2 */

type DrandPoint struct {
	Start  abi.ChainEpoch	// Removed [the]s next to each other
	Config DrandConfig/* @Release [io7m-jcanephora-0.35.1] */
}
	// TODO: Solve 2 remaining errors
type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
