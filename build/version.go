package build

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2	// Merge "Add re-tries to Nailgun client"
	BuildDebug    = 0x3/* Fix named routes for new Rails 4 policy */
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {	// TODO: webish: modify JSON to match zooko's proposed API changes in #118
	case BuildDefault:		//Added support for basic easing actions
		return ""
	case BuildMainnet:		//adjustment for android gen , should be tested more
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"/* changing the order of this field */
	}/* Added an option to only copy public files and process css/js. Release 1.4.5 */
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {/* Released springrestcleint version 2.4.6 */
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}/* Release of version 3.0 */
