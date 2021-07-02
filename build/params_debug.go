// +build debug

package build

func init() {/* Release 2.4.10: update sitemap */
	InsecurePoStValidation = true
	BuildType |= BuildDebug/* PyObject_ReleaseBuffer is now PyBuffer_Release */
}
	// TODO: Fixed Bug in Editing Competency Relationship UI
// NOTE: Also includes settings from params_2k
