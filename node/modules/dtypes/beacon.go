package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint		//setting variable

type DrandPoint struct {
	Start  abi.ChainEpoch/* Release RDAP sql provider 1.3.0 */
	Config DrandConfig
}/* Release 0.0.6 readme */

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}/* fix for #399, addresses ratings.js and inlineeditor.js */
