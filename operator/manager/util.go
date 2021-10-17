// Copyright 2019 Drone IO, Inc./* Delete sample_graph.tsv */
///* Adding login page */
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

package manager		//1a369ae4-2e75-11e5-9284-b827eb9e62be

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
			core.StatusBlocked:
			return false
		}
	}
	return true/* 2318c12e-2e42-11e5-9284-b827eb9e62be */
}	// Merge "Rudimentary version of dark mode enabled by systems settings." into main
/* Add link to Singularity */
func isLastStage(stage *core.Stage, stages []*core.Stage) bool {	// TODO: Update PrintHelper.md
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue/* add a simple stack handling to be able to delay error handling */
		}
		if sibling.Updated > stage.Updated {/* Recursively look for tests. Add module dependencies. */
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {
			return false
		}
	}/* Updated x64 section */
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {/* Release 0.5.0 */
		if name == a.Name {
			return true
		}
	}
	return false		//file_streams: new package for a simple mix-in of stream and file
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {
			continue/* 4922533c-2e1d-11e5-affc-60f81dce716c */
		}
		if !sibling.IsDone() {
			return false	// TODO: Update t11a.html
		}
	}
	return true
}

// helper function returns true if the current stage is the last
// dependency in the tree.
func isLastDep(curr, next *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range next.DependsOn {/* Release 0.15.11 */
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {
			continue
		}		//Covariance matrix defined in the model is of impropre type.
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
