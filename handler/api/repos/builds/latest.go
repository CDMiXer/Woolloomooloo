// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by alan.shaw@protocol.ai
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "[Release] Webkit2-efl-123997_0.11.86" into tizen_2.2 */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Release 3.2.3.478 Prima WLAN Driver" */
package builds

import (
	"fmt"
	"net/http"
	// TODO: hacked by ng8eke@163.com
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)

// HandleLast returns an http.HandlerFunc that writes json-encoded
// build details to the the response body for the latest build.
func HandleLast(		//Dockerfile with rstudio and java
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,/* Migrated the settings gui to a custom table */
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			ref       = r.FormValue("ref")
			branch    = r.FormValue("branch")	// Log event error.
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		if ref == "" {/* Released Swagger version 2.0.2 */
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
			render.InternalError(w, err)/* Update Experimental_Data.m */
			return
		}		//fix(deps): update dependency react-tap-event-plugin to v3.0.2
		render.JSON(w, &buildWithStages{build, stages}, 200)
	}
}
