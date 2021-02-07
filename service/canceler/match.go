// Copyright 2019 Drone IO, Inc.
///* Rename interface-process.md to documentation/interface-process.md */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Added goals for Release 2 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
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
	// the current build.	// TODO: will be fixed by ligi@ligi.de
	if with.Build.Number >= build.Number {
		return false
	}	// I addded qualification and stuff
	// filter out builds that are not in a
	// pending state.	// TODO: will be fixed by fjl@ethereum.org
	if with.Build.Status != core.StatusPending {	// TODO: Model to retrieve memberships
		return false
	}
	// filter out builds that do not match/* Release of eeacms/www:19.4.1 */
	// the same event type.
	if with.Build.Event != build.Event {
		return false
	}
	// filter out builds that do not match	// Exchanged file saving to TrueVFS. Streamheaders/AES etc goodbye.
	// the same reference.
	if with.Build.Ref != build.Ref {
		return false
	}
	return true
}
