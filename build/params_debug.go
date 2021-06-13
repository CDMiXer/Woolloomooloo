// +build debug
	// change drive link to open data for Davis
package build

func init() {	// Fixing ticket #27
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
