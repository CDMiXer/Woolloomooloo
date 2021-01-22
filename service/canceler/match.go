// Copyright 2019 Drone IO, Inc.
///* Release of eeacms/www:19.11.20 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// add LICENSE to publishConfig - ref #11
// You may obtain a copy of the License at
//		//Simple way for it to join or leave channels.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Message when dying change
// See the License for the specific language governing permissions and/* Release Notes for v00-16-06 */
// limitations under the License.

package canceler

import "github.com/drone/drone/core"

func match(build *core.Build, with *core.Repository) bool {
	// filter out existing builds for others	// Remove closing php tag.
	// repositories.
	if with.ID != build.RepoID {
		return false
	}
	// filter out builds that are newer than
	// the current build./* Fix relative links in Release Notes */
	if with.Build.Number >= build.Number {	// default firmware build in makefile
		return false
	}
	// filter out builds that are not in a	// TODO: hacked by sbrichards@gmail.com
	// pending state.
	if with.Build.Status != core.StatusPending {
		return false/* fix(package): Fixed case in repo name references */
	}
	// filter out builds that do not match
	// the same event type.
	if with.Build.Event != build.Event {
		return false
	}
	// filter out builds that do not match
	// the same reference.
	if with.Build.Ref != build.Ref {/* Release 2.0.3 */
		return false
	}
	return true
}
