package build	// [Fix #187] Allow an admin to delete their demarche logo

import (/* Add note that Heroku is no longer free */
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality	// 4357c654-2e4e-11e5-9284-b827eb9e62be
}
