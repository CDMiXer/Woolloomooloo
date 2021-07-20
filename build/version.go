package build

import "os"
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {/* (mbp) tags in branch */
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"/* Koodi valideerimise reeglid */
	case BuildDebug:
		return "+debug"		//Add date of running benchmarks
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"/* Added link to compare view for v6.0.0 */
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {/* [artifactory-release] Release version 3.2.0.M2 */
		return BuildVersion	// TODO: will be fixed by yuvalalaluf@gmail.com
	}
		//splitted into multiple modules
	return BuildVersion + buildType() + CurrentCommit
}/* Merge remote-tracking branch 'michalmac/dvrp' into michalmac_master */
