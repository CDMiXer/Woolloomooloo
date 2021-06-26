// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release jedipus-2.6.29 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Singularize Millionen, Billionen */
// Unless required by applicable law or agreed to in writing, software		//Install video update
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//allow methods with duplicate class names in separate packages
// limitations under the License.

package manager
/* fixed bug in SetCoords. All unit tests are ok now */
import (
	"github.com/drone/drone/core"
)
/* Release of eeacms/www-devel:18.5.8 */
func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,/* Improving performance of remote upload. */
			core.StatusDeclined,
			core.StatusBlocked:
			return false/* Small stringfix for production program, needed for translations */
		}
	}
	return true
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {/* try without quotes */
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {/* Release 0.9.8 */
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {
			return false
		}
	}
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {
		if name == a.Name {
			return true	// Fixed trailing slashes with log files path
		}
	}		//Create AirBox-SiteName-Hsinchu20.txt
	return false/* (jam) Release 2.2b4 */
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {/* Update fl.R */
	deps := map[string]struct{}{}	// TODO: hacked by witek@enjin.io
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

// helper function returns true if the current stage is the last	// TODO: hacked by davidad@alum.mit.edu
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
