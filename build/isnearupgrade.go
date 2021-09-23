package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)
		//no login buttons when user have to choose a city.
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {/* ReleaseNotes: add blurb about Windows support */
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality	// TODO: hacked by ng8eke@163.com
}
