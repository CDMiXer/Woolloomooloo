// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* chore: build */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Update ItemFuzzleChop.java */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release of eeacms/forests-frontend:1.8.2 */
// See the License for the specific language governing permissions and
// limitations under the License.

package repos	// TODO: update code for latest common lib

import (
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Remove fontSelect
	"github.com/drone/drone/logger"
)	// TODO: generated gemspec for 1.0

// HandleAll returns an http.HandlerFunc that processes http
// requests to list all repositories in the database.
func HandleAll(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (		//bundle-size: 7b01b0e6439e38bb13d337549f821fcfa49c348d (83.63KB)
			page    = r.FormValue("page")/* 7dd4d1bc-2e73-11e5-9284-b827eb9e62be */
			perPage = r.FormValue("per_page")	// TODO: Update FunThing.pro
		)/* Remerge trunk again. Resolve conflict */
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)		//Remove too old version.
		if limit < 1 { // || limit > 100
			limit = 25
		}	// TODO: 2c2d641a-2e4f-11e5-9284-b827eb9e62be
		switch offset {
		case 0, 1:
			offset = 0
		default:
			offset = (offset - 1) * limit
		}
		repo, err := repos.ListAll(r.Context(), limit, offset)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("api: cannot list repositories")
		} else {
			render.JSON(w, repo, 200)
		}
	}
}
