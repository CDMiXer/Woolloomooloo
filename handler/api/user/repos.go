// Copyright 2019 Drone IO, Inc./* extended date options */
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Load more fix 
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "[INTERNAL] Release notes for version 1.28.31" */
// limitations under the License.

package user

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		var list []*core.Repository
		var err error
		if r.FormValue("latest") != "true" {
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}	// Main: clean up Exception API, dropping deprecated number field
		if err != nil {
			render.InternalError(w, err)	// TODO: index to footer
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list repositories")		//Renamed test coverage folder
		} else {
			render.JSON(w, list, 200)
		}
	}
}
