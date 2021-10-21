package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)/* changed contact display to membership */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
