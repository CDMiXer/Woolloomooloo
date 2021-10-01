// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
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
/* email updater spurce:local-branches/hawk-hhg/2.5 */
package deploys		//Fixed bug that only active timers could be deleted

import (/* prepared for both: NBM Release + Sonatype Release */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
	// disable nginx access logs for now
	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that handles an
// http.Request to delete a branch entry from the datastore.
func HandleDelete(
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
			target    = chi.URLParam(r, "*")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")
			return
		}
	// TODO: Added to manipulate query doc for how to use new aliased fields.
		err = builds.DeleteDeploy(r.Context(), repo.ID, target)	// Implemented LoC grammar for levels 0-2.
		if err != nil {
			render.InternalError(w, err)		//ggsn admin fields for sheffi and field translate function moved to Table model
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name)./* Merge "Sync models with migrations." */
				Debugln("api: cannot delete deployment")/* Clear error logs for wc 3.0 */
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}		//Update to new revel var names
