// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Remove "close" button in units editor and add "apply" button */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* don't access nil */
//      http://www.apache.org/licenses/LICENSE-2.0	// find by token
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by arajasek94@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release Notes for v02-14-01 */
// See the License for the specific language governing permissions and/* Update admin_logger.php */
// limitations under the License.	// TODO: optimisations turned on

package pulls

import (
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
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
			number, _ = strconv.Atoi(chi.URLParam(r, "pull"))
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {	// TODO: Merge "Improve the ability to enable swap"
			render.NotFound(w, err)		//Temporary fix to home search form (refs #160)
			logger.FromRequest(r)./* Create Web.Release.config */
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")
			return/* run: fix params.end */
		}

		err = builds.DeletePull(r.Context(), repo.ID, number)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).	// #JC-1268 Validation of email was fixed.
				WithField("namespace", namespace).	// TODO: will be fixed by alan.shaw@protocol.ai
				WithField("name", name).
				Debugln("api: cannot delete pr")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}/* Fix directories listing - patch from A. Piesk */
