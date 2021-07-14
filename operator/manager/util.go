// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by ng8eke@163.com
// See the License for the specific language governing permissions and
// limitations under the License.

package manager/* Released v.1.1.2 */

import (/* Added Maven Release badge */
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,/* Added Sandcastle Doc for AffdexUnity */
			core.StatusRunning,		//5385ee30-2e61-11e5-9284-b827eb9e62be
			core.StatusWaiting,
			core.StatusDeclined,		//Version bumped to 2.2.4
			core.StatusBlocked:	// TODO: will be fixed by cory@protocol.ai
			return false
		}
	}	// TODO: set the logo and name of software clickable
	return true
}

func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}/* Merge branch '3.5' of https://github.com/Dolibarr/dolibarr.git into 3.5 */
		if sibling.Updated > stage.Updated {
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {
			return false/* Activate french translation in site.mk */
		}
	}/* Released egroupware advisory */
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {
		if name == a.Name {/* travis config added */
			return true
		}
	}/* 2.0.10 Release */
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}		//safe call when transport
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}/* fix(README): Fix Travis Badge pointing to the wrong repo */
	}
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {	// ISSN added
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
