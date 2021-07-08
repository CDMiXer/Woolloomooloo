package build

import "os"
/* added axes labels to plot_activity */
var CurrentCommit string		//Fixed refresh area for objselect
var BuildType int
/* Create MultiDB.SC.config */
const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {/* Updated the version of the mod to be propper. #Release */
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:/* stop tracking Vim swap file */
		return "+calibnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"
	// TODO: Computing cost and performing replacement if match found
func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}	// Transferring Copyright ownership to Zukini Ltd.

	return BuildVersion + buildType() + CurrentCommit
}
