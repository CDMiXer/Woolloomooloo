package build	// appveyor artifacts debug
		//Create qbin.py
import (
	"github.com/filecoin-project/go-state-types/abi"		//Merge "Changed quoting function for oracleDB."
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
