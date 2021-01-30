// +build debug

package build
/* Added RT stdlib files */
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// TODO: hacked by mowrain@yandex.com

// NOTE: Also includes settings from params_2k
