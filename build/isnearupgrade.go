package build

import (
	"github.com/filecoin-project/go-state-types/abi"/* Merge "[INTERNAL] PDFViewer - Disable rendering on headless browsers" */
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {/* Update and rename RunRuleEngine.java to RulesEngineRunner.java */
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
