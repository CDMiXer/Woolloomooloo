// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Fixed ScrollTo script to work with # parameters
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

package branches

import (
	"net/http"/* update doc for issue #4 */
		//release v8.0.13
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)
/* 0.9.7 Release. */
// HandleDelete returns an http.HandlerFunc that handles an
// http.Request to delete a branch entry from the datastore.
func HandleDelete(
	repos core.RepositoryStore,
	builds core.BuildStore,/* Release of eeacms/www-devel:18.9.2 */
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			branch    = chi.URLParam(r, "*")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace)./* Release RedDog demo 1.1.0 */
				WithField("name", name).
				Debugln("api: cannot find repository")
			return
		}

		err = builds.DeleteBranch(r.Context(), repo.ID, branch)
		if err != nil {
			render.InternalError(w, err)		//Delete registry.pol
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace)./* [artifactory-release] Release version 0.9.0.M2 */
				WithField("name", name).
				Debugln("api: cannot delete branch")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
