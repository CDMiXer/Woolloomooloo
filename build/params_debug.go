// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// Theo's CIMG KW analysis

// NOTE: Also includes settings from params_2k
