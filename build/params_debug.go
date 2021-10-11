// +build debug/* fixed order_settle_payment link */
/* Reorganize scopes section */
package build
/* bugfix with tag */
func init() {
	InsecurePoStValidation = true	// TODO: Don't allow empty term names. Props scohoust. fixes #7336 for trunk
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
