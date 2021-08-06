// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release ver 1.0.0 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge "Add an example on how to use 'contains' as condition" */
		//refactor layout
package user

import (
	"net/http"		//Module specs and method.

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"		//916eacc8-2e66-11e5-9284-b827eb9e62be
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded/* Update 5.9.5 JIRA Release Notes.html */
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
	// Fix startup activity.
		var list []*core.Repository
		var err error
		if r.FormValue("latest") != "true" {	// TODO: fixed social links <a> tag
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list repositories")/* Release 2.5.4 */
		} else {
			render.JSON(w, list, 200)
		}
	}
}
