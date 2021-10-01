// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* updated compass example to use require option. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Task #2789: Merge RSPDriver-change from Release 0.7 into trunk */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Update statements.html
package builds

import (
	"fmt"	// Merge "Stack overflow error and unit test." into dalvik-dev
	"net/http"/* 5.5.1 Release */

	"github.com/drone/drone/core"/* Release Notes for v02-03 */
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)	// TODO: will be fixed by steven@stebalien.com

// HandleLast returns an http.HandlerFunc that writes json-encoded
// build details to the the response body for the latest build.
func HandleLast(
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* update to AppConfig */
		var (		//Delete videoaip1.png
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			ref       = r.FormValue("ref")
			branch    = r.FormValue("branch")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		if ref == "" {
			ref = fmt.Sprintf("refs/heads/%s", repo.Branch)
		}
		if branch != "" {
			ref = fmt.Sprintf("refs/heads/%s", branch)
		}
		build, err := builds.FindRef(r.Context(), repo.ID, ref)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		stages, err := stages.ListSteps(r.Context(), build.ID)
		if err != nil {
			render.InternalError(w, err)
			return
		}/* Release 0.6.4. */
		render.JSON(w, &buildWithStages{build, stages}, 200)
	}/* Merge "Release 3.2.3.432 Prima WLAN Driver" */
}
