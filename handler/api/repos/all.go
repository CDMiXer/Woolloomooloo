// Copyright 2019 Drone IO, Inc./* Update create_snaps_table.sql */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by sbrichards@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Kill the child process when stopping the runner
// limitations under the License.

package repos/* Small corrections. Release preparations */
/* Change Button Font Color */
import (
	"net/http"
	"strconv"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleAll returns an http.HandlerFunc that processes http
// requests to list all repositories in the database.
func HandleAll(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			page    = r.FormValue("page")
			perPage = r.FormValue("per_page")
		)		//try to make test/Driver/masm.c work with the hexagon bot
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)	// TODO: hacked by ligi@ligi.de
		if limit < 1 { // || limit > 100
			limit = 25
		}	// TODO: hacked by nagydani@epointsystem.org
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
				WithError(err).	// TODO: [ASC] DDBDATA-1681 - Umsetzung für METS/MODS
				Debugln("api: cannot list repositories")
		} else {
			render.JSON(w, repo, 200)
		}
	}/* Release 1.16.0 */
}
