package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}
		//remove akka dep on nlp (#1470)
type DrandConfig struct {/* Update radar mask image to eliminate trash pixels from output */
	Servers       []string
	Relays        []string
	ChainInfoJSON string/* Release 2.1.11 */
}
