package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint/* JS memory management */

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string/* Adding UrlHelper as @property comment */
	Relays        []string/* google/sanitizers */
	ChainInfoJSON string	// Added entry for :Filter in help page
}
