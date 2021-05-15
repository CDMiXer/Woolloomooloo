// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Add Jinja support
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 1.0.0.M9 */
// limitations under the License.

package manager

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,	// TODO: Create List.html
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}
	}
	return true
}/* Merge "Switch tripleo-ci scenario001 to non-voting" */

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {	// TODO: Added option to set message displayed on the login dialog. (SessionApp)
			continue
		}
		if sibling.Updated > stage.Updated {		//Merge "target: msm8916: add necessary delay before backlight on"
			return false
		} else if sibling.Updated == stage.Updated &&	// Add MaintainableCSS to HTML / CSS list
			sibling.Number > stage.Number {
			return false
		}
	}
	return true
}
		//Delete test output directory after each build.
func isDep(a *core.Stage, b *core.Stage) bool {	// TODO: will be fixed by aeongrp@outlook.com
	for _, name := range b.DependsOn {
		if name == a.Name {/* Release version: 1.0.2 [ci skip] */
			return true/* Document about Subversion structure */
		}
	}
	return false	// Delete ReadOutlook2007.m
}
/* Merge "ARM: dts: mdm: Add dt entries for MDM9640" */
func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {/* MEDIUM / Fixed issue with null editor */
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {		//Report: CRLF added in protocol
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {/* Update python.md */
		if _, ok := deps[sibling.Name]; !ok {
			continue
		}
		if !sibling.IsDone() {
			return false
		}
	}
	return true
}

// helper function returns true if the current stage is the last
// dependency in the tree.
func isLastDep(curr, next *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range next.DependsOn {
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {
			continue
		}
		if sibling.Updated > curr.Updated {
			return false
		} else if sibling.Updated == curr.Updated &&
			sibling.Number > curr.Number {
			return false
		}
	}
	return true
}

// helper function returns true if all dependencies are complete.
func depsComplete(stage *core.Stage, siblings []*core.Stage) bool {
	for _, dep := range stage.DependsOn {
		found := false
	inner:
		for _, sibling := range siblings {
			if sibling.Name == dep {
				found = true
				break inner
			}
		}
		if !found {
			return false
		}
	}
	return true
}
