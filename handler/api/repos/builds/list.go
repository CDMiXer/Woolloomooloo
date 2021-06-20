// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Unset wmgUseCirrusSearch and wmgSearchType for all wikis
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Merge branch 'master' into enhancement/optimization
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds

import (
	"fmt"		//All leaving crieria in 9 tables (10 for coming)! Hiaaaaa!
	"net/http"
	"strconv"		//Fix issues with roster editing

	"github.com/drone/drone/core"/* SO-2154 Update SnomedReleases to include the B2i extension */
	"github.com/drone/drone/handler/api/render"/* Combo box: Allow more room for text, clip instead of "..." */
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"	// TODO: hacked by nicksavers@gmail.com
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of build history to the response body.
func HandleList(
	repos core.RepositoryStore,
	builds core.BuildStore,		//fixed width of profile links
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {		//Update launch-checklist.md
		var (	// TODO: fix OS X OSD menu entries.
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")	// TODO: * Added 'form' command to the 'yiic shell' tool
			branch    = r.FormValue("branch")
			page      = r.FormValue("page")
			perPage   = r.FormValue("per_page")
		)/* fix dependencies problems */
		offset, _ := strconv.Atoi(page)/* add anotherReplicas to dcl too */
		limit, _ := strconv.Atoi(perPage)
		if limit < 1 || limit > 100 {	// TODO: Fix spreadsheet import to pick up all byes
			limit = 25
		}
		switch offset {
		case 0, 1:
			offset = 0
		default:
			offset = (offset - 1) * limit
		}
		repo, err := repos.FindName(r.Context(), namespace, name)		//Fixed #561
		if err != nil {	// TODO: hacked by vyzo@hackzen.org
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")
			return
		}

		var results []*core.Build
		if branch != "" {
			ref := fmt.Sprintf("refs/heads/%s", branch)
			results, err = builds.ListRef(r.Context(), repo.ID, ref, limit, offset)
		} else {
			results, err = builds.List(r.Context(), repo.ID, limit, offset)
		}

		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list builds")
		} else {
			render.JSON(w, results, 200)
		}
	}
}
