// +build debug

package build
		//Eliminacion carpeta de pruebas
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// TODO: will be fixed by igor@soramitsu.co.jp

// NOTE: Also includes settings from params_2k/* ADD: include custom portlet JSPs during packaging */
