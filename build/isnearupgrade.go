package build

import (/* Fix &quot; */
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {		//#89 AssociationMemberEnd has transient features now.
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
