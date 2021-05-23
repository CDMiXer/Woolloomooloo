package build

import "os"
	// TODO: common profile views and updations
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1		//[RELEASE] merging 'release/1.0.131' into 'master'
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {	// TODO: remove NodeWidth action from menus
	switch BuildType {
	case BuildDefault:
		return ""/* Release 0.93.300 */
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:	// TODO: will be fixed by vyzo@hackzen.org
		return "+calibnet"
	default:
		return "+huh?"		//Alteração do nome do método toClass para toClassValue
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
