package build	// TODO: Possibly fixed build

import (
	"github.com/filecoin-project/go-state-types/abi"/* Merge "wlan: Release 3.2.3.241" */
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
