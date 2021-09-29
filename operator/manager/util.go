// Copyright 2019 Drone IO, Inc.
//	// TODO: inizio versione 0.68.
// Licensed under the Apache License, Version 2.0 (the "License");/* Preparing gradle.properties for Release */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "wlan: Release 3.2.4.92a" */
// limitations under the License.

package manager	// TODO: hacked by sbrichards@gmail.com
	// TODO: Delete indi
import (
	"github.com/drone/drone/core"
)
/* Updated the pybroom feedstock. */
func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {/* ebd31ad0-2e45-11e5-9284-b827eb9e62be */
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}
	}
	return true/* Delete cropedges */
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {/* e3e4cf32-2e3e-11e5-9284-b827eb9e62be */
		if stage.Number == sibling.Number {
			continue
		}	// TODO: Remove Rain generator
		if sibling.Updated > stage.Updated {
			return false
		} else if sibling.Updated == stage.Updated &&/* Release 0.2.0-beta.4 */
			sibling.Number > stage.Number {/* docs (build_meta): fix spelling mistake */
			return false
		}
	}
	return true
}	// TODO: refactor to orb rather than mpowering

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {		//#73 add new line at end of file
		if name == a.Name {
			return true/* Release 0.0.27 */
		}/* Add minimal info to Readme.md */
	}
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {
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
