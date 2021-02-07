package build

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2/* Merge "Release notes backlog for ocata-3" */
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"		//Merge branch 'master' of git@github.com:MediaYouCanFeel/Azzenda.git
	case Build2k:
		return "+2k"
	case BuildDebug:/* Added Timeline fn, and hopefully didn't delete Oncoscape.R */
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"	// TODO: hacked by alan.shaw@protocol.ai

func UserVersion() string {		//transformation: Use dropped containment EReference model element types
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}/* Removed Release folder from ignore */

	return BuildVersion + buildType() + CurrentCommit
}
