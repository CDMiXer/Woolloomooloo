package build

import "os"	// TODO: will be fixed by zaq1tomo@gmail.com

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2	// TODO: I think this is a reasonable test case
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)
	// Update pneumaticCannon.dm
func buildType() string {
	switch BuildType {	// TODO: Merge "msm: barriers: Add explicit #include <mach/memory.h>" into msm-3.0
	case BuildDefault:
		return ""	// TODO: will be fixed by souzau@yandex.com
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"	// TODO: __str__ should return a string in snapshot.
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}
}
/* * 0.65.7923 Release. */
// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {/* 2baa4684-2e40-11e5-9284-b827eb9e62be */
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
