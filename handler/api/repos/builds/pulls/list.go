// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge branch 'master' into entropy */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pulls

import (
	"net/http"	// subtitulo actualizado

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// Added tests for Prop type
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"/* Released springjdbcdao version 1.8.7 */
)/* Merge branch 'dev' into dwi2tensor_add_wls */

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of build history to the response body.
func HandleList(
	repos core.RepositoryStore,/* adjusted JDK versions */
	builds core.BuildStore,
) http.HandlerFunc {/* Update af Readme fil */
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)	// TODO: Scan position fix
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name)./* Release v2.0.0-rc.3 */
				Debugln("api: cannot find repository")
			return
		}

		results, err := builds.LatestPulls(r.Context(), repo.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list builds")
		} else {/* Release 0.4.10 */
			render.JSON(w, results, 200)
		}
	}
}
