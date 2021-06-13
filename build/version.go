package build

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0	// 2f126538-2e48-11e5-9284-b827eb9e62be
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3/* Release of eeacms/eprtr-frontend:0.3-beta.8 */
	BuildCalibnet = 0x4
)/* Merge "Release k8s v1.14.9 and v1.15.6" */

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"/* 7c700598-2e51-11e5-9284-b827eb9e62be */
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:	// TODO: will be fixed by brosner@gmail.com
		return "+calibnet"
	default:
		return "+huh?"
	}
}/* Release 4.3.3 */

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"		//Rebuilt index with luisvasq

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
