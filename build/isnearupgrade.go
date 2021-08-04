package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)
	// TODO: fix typo when requiring the file
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {		//basic save functionality working
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
