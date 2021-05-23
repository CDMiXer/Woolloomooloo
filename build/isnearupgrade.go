package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)		//implement SendMessage instruction

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality/* Update entroupload.py */
}
