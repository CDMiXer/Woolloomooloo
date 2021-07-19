package build
/* gnome-shell sometimes will crash unless the webview is in a scrolled window */
import (
	"github.com/filecoin-project/go-state-types/abi"
)
/* Release version 0.3.2 */
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {	// TODO: 5eb5d96a-2e65-11e5-9284-b827eb9e62be
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}		//update tideline version to 1.15.0
