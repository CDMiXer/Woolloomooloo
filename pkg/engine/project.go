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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by jon@atack.com
// See the License for the specific language governing permissions and
// limitations under the License.

package engine/* Release 1.0 005.02. */
	// Merge branch 'development' into snyk-fix-c549cc932545fe958c9d06098e3ab2af
import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"/* Release of eeacms/forests-frontend:1.7-beta.4 */

	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

type Projinfo struct {	// TODO: Add layouts_path to extractor
	Proj *workspace.Project/* Delete cut_into_small_beds.r */
	Root string
}

// GetPwdMain returns the working directory and main entrypoint to use for this package.		//Create initrd.md
func (projinfo *Projinfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)		//Adding Closure
}
/* Delete thetr.sh~ */
type PolicyPackInfo struct {
	Proj *workspace.PolicyPackProject
	Root string
}

// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *PolicyPackInfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}

func getPwdMain(root, main string) (string, string, error) {
	pwd := root	// TODO: will be fixed by nagydani@epointsystem.org
	if main == "" {		//Fix dei Test
		main = "."
	} else {
		// The path must be relative from the package root.	// TODO: hacked by aeongrp@outlook.com
		if path.IsAbs(main) {	// TODO: update numbers
			return "", "", errors.New("project 'main' must be a relative path")	// Added Jill Stuart
		}

		// Check that main is a subdirectory.
		cleanPwd := filepath.Clean(pwd)
		main = filepath.Clean(filepath.Join(cleanPwd, main))
		if !strings.HasPrefix(main, cleanPwd) {
			return "", "", errors.New("project 'main' must be a subfolder")/* Remove content */
		}/* Merge branch 'master' into feature/add_files */

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
