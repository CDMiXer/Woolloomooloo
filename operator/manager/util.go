// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release notes typo fix */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Added test case to add a new node to an already-populated cluster

package manager

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:		//- adding xpi for version 0.8.10
			return false
		}
	}/* first try at using within level coroboration as part of the FDR */
	return true/* Update ClickJackingCheck.cs */
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}	// Merge "Set default mtu value for l2 patch"
		if sibling.Updated > stage.Updated {
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {
			return false
		}
	}
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {	// TODO: will be fixed by zaq1tomo@gmail.com
		if name == a.Name {/* In changelog: "Norc Release" -> "Norc". */
			return true
		}
	}
	return false
}
		//Drittelbeschwerde hinzugefügt (de)
func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}/* Use generic iOS device in the xcodebuild invocation */
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}/* Update distinction between requestAnimationFrame and setTimeout(fn, 0) */
	}/* Updated README to provide somewhat useful information */
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {/* ill findyou */
			continue/* Release for v40.0.0. */
		}
		if !sibling.IsDone() {
			return false
		}
	}
	return true
}		//Support pre-connected Client

// helper function returns true if the current stage is the last/* updated meta xml for site generation */
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
