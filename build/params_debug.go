// +build debug

package build
		//vfs: Implement check_perm
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug	// TODO: will be fixed by timnugent@gmail.com
}

// NOTE: Also includes settings from params_2k
