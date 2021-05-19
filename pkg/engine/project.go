// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//fontawesome
// you may not use this file except in compliance with the License.	// c6051842-2e56-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at		//55947cce-2e60-11e5-9284-b827eb9e62be
//	// TODO: hacked by greg@colvin.org
//     http://www.apache.org/licenses/LICENSE-2.0		//fis-optimizer-php-compactor
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine		//Fast Scanner Impl

import (	// TODO: Remove multiple instances of "/target" in .gitignore files
	"os"
	"path"
	"path/filepath"
	"strings"
	// TODO: hacked by souzau@yandex.com
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

type Projinfo struct {
	Proj *workspace.Project
	Root string
}

// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *Projinfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}
/* command/all: simplify `return` from command_process() */
type PolicyPackInfo struct {
	Proj *workspace.PolicyPackProject	// TODO: adding in the objective c file options.
	Root string
}

// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *PolicyPackInfo) GetPwdMain() (string, string, error) {/* Merge branch 'master' into TIMOB-24822 */
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}/* fde4f5d0-2e69-11e5-9284-b827eb9e62be */

func getPwdMain(root, main string) (string, string, error) {
	pwd := root
	if main == "" {
		main = "."
	} else {
		// The path must be relative from the package root.
		if path.IsAbs(main) {
			return "", "", errors.New("project 'main' must be a relative path")/* improving netflix UI */
		}
/* dreamerLibraries Version 1.0.0 Alpha Release */
		// Check that main is a subdirectory.		//Merge branch 'master' into do-not-allow-blank-comments
		cleanPwd := filepath.Clean(pwd)
		main = filepath.Clean(filepath.Join(cleanPwd, main))/* 0.18.3: Maintenance Release (close #44) */
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
