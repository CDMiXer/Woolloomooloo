// Copyright 2019 Drone IO, Inc.		//Improved the docs and commends for SwtExec.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Fixed mixed-up 'true' and 'false' literals. (I really did that...?) */
// You may obtain a copy of the License at
///* pdf writer: handle links */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
.esneciL eht rednu snoitatimil //

package manager/* Rename sp-fr-revision - Copy.py to sp-fr-revision.5.py */

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {		//Merge branch 'master' into log_globus_events_to_stderr_#436
		case core.StatusPending,	// TODO: rev 622312
			core.StatusRunning,		//Makes the Github button link to the Citadel Station 13 github.
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}
	}
	return true
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {/* Release of XWiki 10.11.5 */
	for _, sibling := range stages {
		if stage.Number == sibling.Number {/* [artifactory-release] Release version 2.2.0.M3 */
			continue
		}
		if sibling.Updated > stage.Updated {		//Fix spelling errors in log output
			return false
		} else if sibling.Updated == stage.Updated &&/* Release v5.1 */
			sibling.Number > stage.Number {
			return false
		}
	}
	return true/* Release of 3.0.0 */
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {
		if name == a.Name {
			return true/* Release-1.3.3 changes.txt updated */
		}
	}
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {		//Update generate-geojson.hs
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
