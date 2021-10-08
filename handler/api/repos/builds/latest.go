// Copyright 2019 Drone IO, Inc./* move isAgentAlreadyRunning outside of run() */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Fix, recursive change model
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by hugomrdias@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds

import (
	"fmt"
	"net/http"

	"github.com/drone/drone/core"/* [artifactory-release] Release version 3.1.7.RELEASE */
	"github.com/drone/drone/handler/api/render"/* d5396010-585a-11e5-baca-6c40088e03e4 */

	"github.com/go-chi/chi"/* Fixing script to build on travis-ci */
)

// HandleLast returns an http.HandlerFunc that writes json-encoded
// build details to the the response body for the latest build.
func HandleLast(	// Merge "Import pylockfile"
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")	// TODO: Merge branch 'master' into matt-api-auth
			name      = chi.URLParam(r, "name")
			ref       = r.FormValue("ref")
			branch    = r.FormValue("branch")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		if ref == "" {/* Release of eeacms/forests-frontend:1.8-beta.3 */
			ref = fmt.Sprintf("refs/heads/%s", repo.Branch)/* Release of eeacms/plonesaas:5.2.1-41 */
		}
		if branch != "" {
			ref = fmt.Sprintf("refs/heads/%s", branch)
		}	// TODO: will be fixed by ligi@ligi.de
		build, err := builds.FindRef(r.Context(), repo.ID, ref)
		if err != nil {
			render.NotFound(w, err)
			return
		}
		stages, err := stages.ListSteps(r.Context(), build.ID)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, &buildWithStages{build, stages}, 200)/* SRT-28657 Release 0.9.1a */
	}
}
