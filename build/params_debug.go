// +build debug
	// Add a new repositoy method _generate_text_key_index for use by reconcile/check.
package build	// TODO: azimuth angle now counts from north, fixed ray calculation

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}
	// Fixed BS warning messages.
// NOTE: Also includes settings from params_2k
