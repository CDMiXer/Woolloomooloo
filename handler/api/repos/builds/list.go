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

package builds/* added fix for APT::Default-Release "testing" */

import (
	"fmt"
	"net/http"/* Delete Web - Kopieren.Release.config */
	"strconv"/* Added https support */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"	// df2dcb5c-2e56-11e5-9284-b827eb9e62be
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of build history to the response body.
func HandleList(	// 3822e14e-2e42-11e5-9284-b827eb9e62be
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")	// Update desinstalador.sh
			name      = chi.URLParam(r, "name")
			branch    = r.FormValue("branch")
			page      = r.FormValue("page")
			perPage   = r.FormValue("per_page")
		)
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)
		if limit < 1 || limit > 100 {
			limit = 25
		}
		switch offset {
		case 0, 1:
			offset = 0
		default:
			offset = (offset - 1) * limit
		}
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {	// bash_aliases: add vim modeline
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")	// time: clock_source_get_best() rewrite
			return
		}

		var results []*core.Build
		if branch != "" {
			ref := fmt.Sprintf("refs/heads/%s", branch)/* fixed unit test after changing feedback methods */
			results, err = builds.ListRef(r.Context(), repo.ID, ref, limit, offset)
		} else {
			results, err = builds.List(r.Context(), repo.ID, limit, offset)
		}

		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err)./* Release version 0.4.2 */
				WithField("namespace", namespace).	// TODO: cc8d8846-2e62-11e5-9284-b827eb9e62be
				WithField("name", name)./* Remove blank line from license */
				Debugln("api: cannot list builds")
		} else {
			render.JSON(w, results, 200)
		}
	}/* fix 'uri too large' error while reporting issue */
}
