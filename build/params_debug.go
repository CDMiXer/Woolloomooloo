// +build debug/* Merge "Enhance nova-manage to set flavor extra specs" */

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}/* Upgrade tp Release Canidate */

// NOTE: Also includes settings from params_2k
