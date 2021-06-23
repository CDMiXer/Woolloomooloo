package build
/* Release of eeacms/forests-frontend:1.8-beta.7 */
import (
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality/* Release v0.5.0.5 */
}
