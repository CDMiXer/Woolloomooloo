// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Restful Auth example
// You may obtain a copy of the License at		//Merge "Race condition patched in Email autocompletion." into honeycomb
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Updated Release Links */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())	// TODO: Add variables to exclude the code that uses the Vector_Mappers
/* Release version 6.3 */
		list, err := repos.List(r.Context(), viewer)	// 48feb29a-2e4c-11e5-9284-b827eb9e62be
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list remote repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
