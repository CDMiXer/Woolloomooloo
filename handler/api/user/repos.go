// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Added My Releases section */
//
// Unless required by applicable law or agreed to in writing, software	// [ powerpoint ] added badges to the README.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Test Heading Formatting
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)
	// Hard lower bound on shape parameter of GammaSite*Model
// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryStore) http.HandlerFunc {	// TODO: Update baseBot.py
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		var list []*core.Repository
		var err error	// TODO: SO-1710: converted SnomedRelationshipReadRequest to GetRequest.
		if r.FormValue("latest") != "true" {
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}
		if err != nil {/* Release 2.1.3 - Calendar response content type */
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
