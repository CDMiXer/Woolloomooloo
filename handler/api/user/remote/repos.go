// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Merge "Whitespace and spelling fixes in cx.stats js and less"
//	// New improved layouting algorithm based on TouchGraph
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Merge "Solve the infinite loop on clearExternalStorageDataSync" into nyc-dev
// limitations under the License.

package remote/* Fix reverse syntax in docs */

import (/* Released 1.0.0 🎉 */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Merge branch 'master' into fix-right-mouse-drag-large-view */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
/* added tier to manta http status */
		list, err := repos.List(r.Context(), viewer)/* [skip ci] max */
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list remote repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}	// TODO: will be fixed by earlephilhower@yahoo.com
