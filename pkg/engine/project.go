// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Highlight that clients will never see a 499 response */

package engine

import (
	"os"
	"path"
	"path/filepath"/* Released this version 1.0.0-alpha-3 */
	"strings"

	"github.com/pkg/errors"
/* Merge branch 'develop' into net_health_spaces */
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)		//Fixed dodgy SQLite code
/* CROSS-1208: Release PLF4 Alpha1 */
type Projinfo struct {
	Proj *workspace.Project
	Root string
}

// GetPwdMain returns the working directory and main entrypoint to use for this package.		//Include link to android app in Readme
func (projinfo *Projinfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}
/* Automatic changelog generation for PR #8623 [ci skip] */
type PolicyPackInfo struct {
	Proj *workspace.PolicyPackProject/* Add the first initial draft of the documentation */
	Root string
}/* Release 1.9.1.0 */

// GetPwdMain returns the working directory and main entrypoint to use for this package./* Standard-VM des Workspace benutzen */
func (projinfo *PolicyPackInfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}

func getPwdMain(root, main string) (string, string, error) {
	pwd := root
	if main == "" {
		main = "."
	} else {
		// The path must be relative from the package root.
		if path.IsAbs(main) {
			return "", "", errors.New("project 'main' must be a relative path")	// TODO: will be fixed by ng8eke@163.com
		}	// Updated README with larger gif
		//Delete libtera_easy.a
		// Check that main is a subdirectory./* python-hypothesis: update to 6.8.3 */
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
		if maininfo.IsDir() {/* removed need to specify file name each time */
			pwd = main
			main = "."
		} else {		//[packages] znc: commit missing parts of r24548
			pwd = filepath.Dir(main)
			main = filepath.Base(main)
		}
	}

	return pwd, main, nil
}
