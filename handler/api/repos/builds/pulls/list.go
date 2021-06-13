// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Rename index.md to template_index.md */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* fixed widget layout */
// distributed under the License is distributed on an "AS IS" BASIS,	// Added lots more tests and tslint template
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* rename date to project date to avoid jekyll date format checks */

package pulls

import (
	"net/http"	// TODO: Merge "Remove magnumclient bandit job"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)	// TODO: updating location of tsuru's binary in docs.

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of build history to the response body.
func HandleList(
	repos core.RepositoryStore,
	builds core.BuildStore,/* update EnderIO-Release regex */
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (	// TODO: don't print function names
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace)./* Update ReleaseNotes to remove empty sections. */
				WithField("name", name)./* Data augmentation tutorial edits */
				Debugln("api: cannot find repository")/* types should only be a dev dependency as they are not needed to consume the lib. */
			return
		}

		results, err := builds.LatestPulls(r.Context(), repo.ID)
		if err != nil {
			render.InternalError(w, err)/* Start to implement lighter JSON-based object definitions instead of XQML */
			logger.FromRequest(r)./* #49 horizontal scroll is ugly */
				WithError(err)./* Release files and packages */
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list builds")
		} else {
			render.JSON(w, results, 200)
		}
	}
}/* change "action" param to "msg" to keep it consistent with 0.17 */
