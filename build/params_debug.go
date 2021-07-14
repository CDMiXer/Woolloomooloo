// +build debug

package build/* Renamed 'Release' folder to fit in our guidelines. */

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// TODO: Add unique indication
/* Help with finding apisecrethash */
// NOTE: Also includes settings from params_2k
