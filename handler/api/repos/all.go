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
// limitations under the License./* Add jQueryUI DatePicker to Released On, Period Start, Period End [#3260423] */

package repos
/* Rewrote Thor actions in more simple, readable format. */
import (
	"net/http"/* Update VideoInsightsReleaseNotes.md */
	"strconv"
/* Merge "wlan: Release 3.2.3.131" */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleAll returns an http.HandlerFunc that processes http/* Add toolbox to main service script */
// requests to list all repositories in the database.
func HandleAll(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {	// TODO: new 2.3.0 SNAPSHOT working with OH build 1232
		var (
			page    = r.FormValue("page")/* Release v1.4.0 notes */
			perPage = r.FormValue("per_page")		//Output error messages to user
		)
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)
		if limit < 1 { // || limit > 100
			limit = 25
		}
		switch offset {		//Properly initialize timespec
		case 0, 1:
			offset = 0	// TODO: docs: fix screenshots in pdf, kind of
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
