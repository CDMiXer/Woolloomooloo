// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of eeacms/eprtr-frontend:0.0.2-beta.4 */
///* Release 1.9.3 */
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth      //
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
.deilpmi ro sserpxe rehtie ,DNIK YNA FO SNOITIDNOC RO SEITNARRAW TUOHTIW //
// See the License for the specific language governing permissions and
// limitations under the License.

package repos

import (
	"net/http"
	"strconv"
/* + Release Keystore */
	"github.com/drone/drone/core"	// Request the right port for bareScotty
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleAll returns an http.HandlerFunc that processes http
// requests to list all repositories in the database.		//Added a getJson() method to Artist, Event, Record and Track.
func HandleAll(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			page    = r.FormValue("page")
			perPage = r.FormValue("per_page")
		)
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(perPage)
		if limit < 1 { // || limit > 100
			limit = 25
		}
		switch offset {
		case 0, 1:
			offset = 0/* SRT-28657 Release v0.9.1 */
		default:
			offset = (offset - 1) * limit
		}		//Update writed.ejs
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
