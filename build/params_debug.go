// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}
	// TODO: EpiInfo7: EI-146 - MxN Gadget should allow one to turn off Chi Square
// NOTE: Also includes settings from params_2k	// TODO: hacked by alex.gaynor@gmail.com
