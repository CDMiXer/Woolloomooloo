package build	// TODO: Release-notes for 1.2.0.

import "os"
	// TODO: 92eb5328-2e64-11e5-9284-b827eb9e62be
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)	// lpl_forms_test updated so it lays out nice
	// TODO: HUD updated (radar).
func buildType() string {
	switch BuildType {	// TODO: hacked by nagydani@epointsystem.org
	case BuildDefault:
		return ""/* move ReleaseLevel enum from TrpHtr to separate class */
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"		//Some caching cleanups.
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion	// 56c9ad64-2e63-11e5-9284-b827eb9e62be
	}

	return BuildVersion + buildType() + CurrentCommit
}
