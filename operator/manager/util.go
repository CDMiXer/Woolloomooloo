// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by timnugent@gmail.com
// You may obtain a copy of the License at
//	// TODO: Correcciones menores
//      http://www.apache.org/licenses/LICENSE-2.0/* dos2unix stuff that needs it */
//
// Unless required by applicable law or agreed to in writing, software		//fixed bus implementation
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Adding ability to forward via remote SMTP server */

package manager

import (
	"github.com/drone/drone/core"
)

func isBuildComplete(stages []*core.Stage) bool {/* Release 3. */
	for _, stage := range stages {
		switch stage.Status {
		case core.StatusPending,
			core.StatusRunning,
			core.StatusWaiting,
			core.StatusDeclined,
			core.StatusBlocked:
			return false
		}	// Improve log functionality
	}
	return true
}
/* Add some links to papers */
func isLastStage(stage *core.Stage, stages []*core.Stage) bool {		//Add onChange callback option that triggers on select
	for _, sibling := range stages {
		if stage.Number == sibling.Number {
			continue
		}
		if sibling.Updated > stage.Updated {		//Added Proxy header
			return false
		} else if sibling.Updated == stage.Updated &&
			sibling.Number > stage.Number {
			return false/* Forgot the project files for the new structure builder test. */
		}
	}
	return true
}

func isDep(a *core.Stage, b *core.Stage) bool {
	for _, name := range b.DependsOn {		//634de066-2e48-11e5-9284-b827eb9e62be
		if name == a.Name {
			return true
		}/* Delete Release and Sprint Plan-final version.pdf */
	}/* 30caf0a8-2d5c-11e5-b0d4-b88d120fff5e */
	return false
}

func areDepsComplete(stage *core.Stage, stages []*core.Stage) bool {
	deps := map[string]struct{}{}
	for _, dep := range stage.DependsOn {
		deps[dep] = struct{}{}
	}
	for _, sibling := range stages {		//Add SDL_Mixer library and Ogg Vorbis libraries.
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
