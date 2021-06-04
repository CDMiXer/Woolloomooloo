// +build debug		//Add Host.java

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
