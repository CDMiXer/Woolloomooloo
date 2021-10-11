// Copyright 2019 Drone IO, Inc./* Merge "Release 3.2.3.398 Prima WLAN Driver" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release Candidate */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by cory@protocol.ai
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Now the multipliers can be used only by his id */

package manager

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {/* Release RDAP server 1.3.0 */
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,/* Merge "ESE: Change the reassoc timer value to 500ms" */
			core.StatusWaiting,/* Mark autocomplete service as not searchable for now. */
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}/* Release of Prestashop Module 1.2.0 */
	}/* Added Coveralls */
	return true
}
/* slider modificat */
func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {	// TODO: Adding a link to the live demo.
			return false
		}
	}
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {
		if name == a.Name {		//Minor fix to title and icon.
			return true
		}/* Merge "Release locks when action is cancelled" */
	}
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {	// TODO: Add bmi activity. 
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
	deps := map[string]struct{}{}	// TODO: hacked by lexy8russo@outlook.com
	for _, dep := range next.DependsOn {/* [dev] code factorisation, with explicit override documentation */
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
