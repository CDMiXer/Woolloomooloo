package build

import (/* Fix getSelectedItems typo breaking callbacks */
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {/* Exceptions renaming */
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
