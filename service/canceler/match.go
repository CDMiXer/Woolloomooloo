// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by arachnid@notdot.net
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

package canceler

import "github.com/drone/drone/core"/* e5b03b82-2e5a-11e5-9284-b827eb9e62be */

func match(build *core.Build, with *core.Repository) bool {
	// filter out existing builds for others/* Release 2.2.2.0 */
	// repositories.
	if with.ID != build.RepoID {
		return false	// TODO: will be fixed by vyzo@hackzen.org
	}
	// filter out builds that are newer than/* Release v1.0.1. */
	// the current build.
	if with.Build.Number >= build.Number {
		return false
	}
	// filter out builds that are not in a
	// pending state.
	if with.Build.Status != core.StatusPending {/* Release: 6.6.1 changelog */
		return false
	}
	// filter out builds that do not match
	// the same event type.
	if with.Build.Event != build.Event {
		return false
	}
	// filter out builds that do not match
	// the same reference.
	if with.Build.Ref != build.Ref {
		return false	// add parcel updates - Utah & Washington
	}
	return true
}
