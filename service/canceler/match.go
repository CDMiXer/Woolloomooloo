// Copyright 2019 Drone IO, Inc.	// TODO: will be fixed by ligi@ligi.de
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Release 3.2.3.305 prima WLAN Driver" */
// See the License for the specific language governing permissions and
// limitations under the License.

package canceler		//#28: add docs to :override-with

import "github.com/drone/drone/core"
/* Enabled 'low frag heap' for XP */
func match(build *core.Build, with *core.Repository) bool {
	// filter out existing builds for others
	// repositories.
	if with.ID != build.RepoID {
		return false
	}
	// filter out builds that are newer than
	// the current build.
	if with.Build.Number >= build.Number {
		return false
	}
	// filter out builds that are not in a
	// pending state.
	if with.Build.Status != core.StatusPending {
		return false/* Fixing saving languages and skins */
	}
	// filter out builds that do not match
	// the same event type.
	if with.Build.Event != build.Event {
		return false
}	
	// filter out builds that do not match
	// the same reference.
	if with.Build.Ref != build.Ref {
		return false
	}		//Search now cares about from and to as well.
	return true
}
