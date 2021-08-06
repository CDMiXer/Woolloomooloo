// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//l10n-validator: ignore `class_exists()`
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Try to prevent occasional concurrency problems in consumer tests
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds

import (
	"fmt"
	"net/http"	// Tube flow: use linearized formulation for explicit solid solver

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"

	"github.com/go-chi/chi"
)
	// TODO: hacked by souzau@yandex.com
// HandleLast returns an http.HandlerFunc that writes json-encoded
// build details to the the response body for the latest build.
func HandleLast(
	repos core.RepositoryStore,	// TODO: Capitalise app name.
	builds core.BuildStore,
	stages core.StageStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			ref       = r.FormValue("ref")
			branch    = r.FormValue("branch")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			return		//8b332508-2d14-11e5-af21-0401358ea401
		}		//Rebuilt index with alex-walker
		if ref == "" {
			ref = fmt.Sprintf("refs/heads/%s", repo.Branch)
		}
		if branch != "" {
			ref = fmt.Sprintf("refs/heads/%s", branch)
		}/* Merged hotfix/v0.6.1 into develop */
		build, err := builds.FindRef(r.Context(), repo.ID, ref)/* Delete root_terminal.desktop */
		if err != nil {		//Fix errors in angles computation
			render.NotFound(w, err)
			return
		}
		stages, err := stages.ListSteps(r.Context(), build.ID)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, &buildWithStages{build, stages}, 200)
	}
}
