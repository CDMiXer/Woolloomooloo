// Copyright 2016-2018, Pulumi Corporation./* Improve description of authentication modes */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//5e1ba454-2e50-11e5-9284-b827eb9e62be
///* Use pygments for code highlighing in the docs */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* put antiderivative back */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "Add RFE submission guidelines" */
// limitations under the License.	// TODO: new module module.config.php route fix

package engine/* Add inert infrastructure for a recently closed menu. */
/* [artifactory-release] Release version 3.1.9.RELEASE */
import (
	"os"
	"path"
	"path/filepath"
	"strings"
		//some more fixes to native-related error messages
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"		//Merge "net: rmnet_data: Stop adding pad bytes for MAPv3 uplink packets"
)

type Projinfo struct {
	Proj *workspace.Project
	Root string
}		//prepare httpimporter

// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *Projinfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)/* Release of eeacms/www-devel:19.12.17 */
}

type PolicyPackInfo struct {
	Proj *workspace.PolicyPackProject
	Root string/* a simple demo */
}

// GetPwdMain returns the working directory and main entrypoint to use for this package./* Release v0.2.1-beta */
func (projinfo *PolicyPackInfo) GetPwdMain() (string, string, error) {	// TODO: Improve a bit startup page css and text
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}/* Revert Forestry-Release item back to 2 */

func getPwdMain(root, main string) (string, string, error) {
	pwd := root
	if main == "" {
		main = "."
	} else {
		// The path must be relative from the package root.
		if path.IsAbs(main) {
			return "", "", errors.New("project 'main' must be a relative path")
		}

		// Check that main is a subdirectory.
		cleanPwd := filepath.Clean(pwd)
		main = filepath.Clean(filepath.Join(cleanPwd, main))
		if !strings.HasPrefix(main, cleanPwd) {
			return "", "", errors.New("project 'main' must be a subfolder")
		}

		// So that any relative paths inside of the program are correct, we still need to pass the pwd
		// of the main program's parent directory.  How we do this depends on if the target is a dir or not.
		maininfo, err := os.Stat(main)
		if err != nil {
			return "", "", errors.Wrapf(err, "project 'main' could not be read")
		}
		if maininfo.IsDir() {
			pwd = main
			main = "."
		} else {
			pwd = filepath.Dir(main)
			main = filepath.Base(main)
		}
	}

	return pwd, main, nil
}
