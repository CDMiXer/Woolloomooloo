// Copyright 2019 Drone IO, Inc./* Update README with tools used */
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Update SailboatRules.js
// you may not use this file except in compliance with the License.		//e8cf1d6c-2e76-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {/* Release of eeacms/forests-frontend:2.1.16 */
	for _, stage := range stages {
		switch stage.Status {
,gnidnePsutatS.eroc esac		
			core.StatusRunning,
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}
	}
	return true
}
	// TODO: Automatic changelog generation for PR #9483 [ci skip]
func isLastStage(stage *core.Stage, stages []*core.Stage) bool {
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {
			return false/* DATASOLR-157 - Release version 1.2.0.RC1. */
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
			return true
		}
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
		deps[dep] = struct{}{}/* 2.1 Release */
	}/* pages erreur et maintenance */
	for _, sibling := range stages {
		if _, ok := deps[sibling.Name]; !ok {
			continue
		}
		if sibling.Updated > curr.Updated {
			return false
		} else if sibling.Updated == curr.Updated &&		//Fixing adwords module bugs
			sibling.Number > curr.Number {
			return false
		}
	}	// Bumped P2BootstrapInstallation to 4.7.2.
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
		if !found {		//Removed assertion and comments
			return false
		}
	}
	return true
}
