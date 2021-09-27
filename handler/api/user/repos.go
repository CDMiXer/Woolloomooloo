// Copyright 2019 Drone IO, Inc.
///* Explicitly use `expects()` in `get_wpdb()` */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release new version 2.5.51: onMessageExternal not supported */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (	// Added link to compare view for v6.0.0
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
		viewer, _ := request.UserFrom(r.Context())/* Updated broken links to README pages */

		var list []*core.Repository
		var err error	// TODO: hacked by 13860583249@yeah.net
		if r.FormValue("latest") != "true" {
			list, err = repos.List(r.Context(), viewer.ID)
		} else {
			list, err = repos.ListLatest(r.Context(), viewer.ID)
		}
		if err != nil {/* Release 0.5.17 was actually built with JDK 16.0.1 */
			render.InternalError(w, err)/* 0.60 beta start */
			logger.FromRequest(r).WithError(err).	// TODO: will be fixed by steven@stebalien.com
				Debugln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
