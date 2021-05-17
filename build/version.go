package build

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0		//- finish model_factory - nothing major, just lots of little fixes
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"	// Update results table
	case Build2k:
		return "+2k"		//update xml loading to function with flat files as well as from classpath
	case BuildDebug:/* Merge "Release 4.0.10.003  QCACLD WLAN Driver" */
		return "+debug"/* Rename Why Mock HTTP?.md to why-mock-http?.md */
	case BuildCalibnet:
		return "+calibnet"/* Added multitouch support. Release 1.3.0 */
	default:/* Release version [10.8.1] - prepare */
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"/* CSS Typo (missing pound) */
/* Release 0.3.5 */
func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
