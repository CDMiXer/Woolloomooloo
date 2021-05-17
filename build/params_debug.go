// +build debug
/* Release build needed UndoManager.h included. */
package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
