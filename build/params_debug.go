// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// Move several classes from ast to parser; comments++

// NOTE: Also includes settings from params_2k
