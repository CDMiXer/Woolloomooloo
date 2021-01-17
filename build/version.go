package build

import "os"
/* Restore Template Data */
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:/* Remove pic */
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:	// TODO: 7a70ad46-2e52-11e5-9284-b827eb9e62be
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"	// TODO: Update emojione.meta
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"	// TODO: Delete demowithosinfo.bat

func UserVersion() string {/* chore(bugfix): Wrapped bootstrap with Meteor.startup */
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
}	

	return BuildVersion + buildType() + CurrentCommit
}/* Update Message error */
