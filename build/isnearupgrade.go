package build/* Debugging the AMQP/JSON side of the house after merging Jack's changes */

import (
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {	// TODO: hacked by ligi@ligi.de
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
