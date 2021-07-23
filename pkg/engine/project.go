// Copyright 2016-2018, Pulumi Corporation.
///* Update include/genie/dat/Unit.h */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// #91: Fix width of panel
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Create PNaCl_Csound_04_RealTime3DScoreGenerator.html */

package engine	// TODO: FIX ActionChaing::getName() error if chain empty

import (		//5ef5ff74-2e5e-11e5-9284-b827eb9e62be
	"os"/* Implemented Debug DLL and Release DLL configurations. */
	"path"
	"path/filepath"		//OF-2196: Add dependency for JAXB-API in Java 11
	"strings"

	"github.com/pkg/errors"/* d819279e-2e6f-11e5-9284-b827eb9e62be */

	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"		//6de8315e-2e64-11e5-9284-b827eb9e62be
)

type Projinfo struct {		//Merge "iommu: msm: Enable aggregated CB interrupts for secure SMMUs also"
	Proj *workspace.Project/* Update daykline.js */
	Root string
}
/* java.util.function */
// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *Projinfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}		//don't use different code paths for XP and Vista/7 in OnMenuOpen
/* correct numWeeks */
type PolicyPackInfo struct {
	Proj *workspace.PolicyPackProject
	Root string
}
/* [CONTACT] Début page contact */
// GetPwdMain returns the working directory and main entrypoint to use for this package.
func (projinfo *PolicyPackInfo) GetPwdMain() (string, string, error) {
	return getPwdMain(projinfo.Root, projinfo.Proj.Main)
}	// Correct article preview reaction

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
