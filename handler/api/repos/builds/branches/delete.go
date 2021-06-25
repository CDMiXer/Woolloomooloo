// Copyright 2019 Drone IO, Inc.	// TODO: Delete Spark_Machine_Learning_Pipeline_v4.ipynb
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Adding explicit HOME env var. */
// You may obtain a copy of the License at/* Release 4.5.3 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// support multiple urls, --pack and --dry-run
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package branches

import (	// TODO: will be fixed by mail@overlisted.net
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* build and display APA citation from configured metadata; refs #16152 */
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
			namespace = chi.URLParam(r, "owner")		//Work around KARAF-4771; improve performance; add Wildfly
			name      = chi.URLParam(r, "name")
			branch    = chi.URLParam(r, "*")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")/* Delete ATFWD-daemon */
			return
		}

		err = builds.DeleteBranch(r.Context(), repo.ID, branch)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot delete branch")/* Creación del archivo index.html */
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
