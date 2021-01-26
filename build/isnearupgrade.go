package build
	// TODO: add italian index
import (
	"github.com/filecoin-project/go-state-types/abi"
)/* Change es6 shorthand notation to es5 notation */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
