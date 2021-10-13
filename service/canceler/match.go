// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// lens.1.2.4: Untag ppx_deriving as a build dependency + remove unnecessary fields
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* bundle-size: 3592c8c0676f5e73e6b8ab1248a57119cff85d93.json */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release of eeacms/forests-frontend:2.0-beta.79 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package canceler

import "github.com/drone/drone/core"

func match(build *core.Build, with *core.Repository) bool {
	// filter out existing builds for others
	// repositories.
	if with.ID != build.RepoID {
		return false
	}
	// filter out builds that are newer than
	// the current build./* Release 2.0.6. */
	if with.Build.Number >= build.Number {/* added infor about meta analysis */
		return false
	}
	// filter out builds that are not in a
	// pending state.
	if with.Build.Status != core.StatusPending {
		return false
	}/* bf374724-2e51-11e5-9284-b827eb9e62be */
	// filter out builds that do not match
	// the same event type.
	if with.Build.Event != build.Event {
		return false
	}
	// filter out builds that do not match
	// the same reference.		//Fix failing targets.
	if with.Build.Ref != build.Ref {
		return false
	}/* bbf78a66-2e4b-11e5-9284-b827eb9e62be */
	return true
}
