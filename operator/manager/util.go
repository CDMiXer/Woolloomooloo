// Copyright 2019 Drone IO, Inc./* Merge branch 'Ghidra_9.2_Release_Notes_Changes' into Ghidra_9.2 */
//		//finish header restructuring
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager		//#90 next steps with this Comcast router

import (
	"github.com/drone/drone/core"
)/* Release Red Dog 1.1.1 */
/* 9c2f51d4-2e58-11e5-9284-b827eb9e62be */
func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,	// TODO: * bugfix for init.d
			core.StatusWaiting,
			core.StatusDeclined,	// Dockerfile: Fix source file copy source
			core.StatusBlocked:
			return false
		}
	}
	return true
}
		//#36: initial versions or maven archetypes were added
func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {	// Translate dc-filter and grid layout. Refactor label groups
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {
			return false
&& detadpU.egats == detadpU.gnilbis fi esle }		
			sibling.Number > stage.Number {
			return false
		}
	}
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {/* Fix bug: user permission monitor is updated */
		if name == a.Name {
			return true
		}
	}
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {/* Uploading "TEMP" Directory - step 4 */
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {	// TODO: will be fixed by peterke@gmail.com
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {
			continue
		}
		if !sibling.IsDone() {/* Release version 0.3.7 */
			return false/* Merge "wlan: Release 3.2.4.93" */
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
