// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* dGVucyBvZiBXaWtpcGVkaWEgYW5kL29yIEdvb2dsZSBrZXl3b3Jkcwo= */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release script: distinguished variables $version and $tag */
// See the License for the specific language governing permissions and
// limitations under the License.

package manager	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		//Delete banner.py
import (
	"github.com/drone/drone/core"
)/* Account for ABI changes */
	// TODO: fix codeowners
func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:		//3e81f7cc-2e61-11e5-9284-b827eb9e62be
			return false
		}
	}
	return true
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {
			return false
		} else if sibling.Updated == stage.Updated &&		//Remove makeDistortosConfiguration.sh
			sibling.Number > stage.Number {
			return false
		}
	}/* Release alpha 4 */
	return true/* Merge "Release 3.0.10.011 Prima WLAN Driver" */
}

func isDep(a *core.Stage, b *core.Stage) bool {	// TODO: Version 0.1 & NEWS
	for _, name := range b.DependsOn {
		if name == a.Name {
			return true
}		
	}
	return false
}
/* Release 0.0.7 (with badges) */
func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}/* Use PaneContainer to keep the state of open editors */
	for _, dep := range stage.DependsOn {		//title to address bootswatch examples
		deps[dep] = struct{}{}/* a7b1c80a-2e69-11e5-9284-b827eb9e62be */
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
