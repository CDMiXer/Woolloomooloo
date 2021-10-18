package build	// TODO: adjusted wording.

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)
	// TODO: hacked by hugomrdias@gmail.com
func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:/* added local version of Droid font for offline mode */
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}/* Add cursor skip and wraparound. */
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
