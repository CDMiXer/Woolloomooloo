// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update DavisNe.md
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/www-devel:19.3.18 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* +Release notes, +note that static data object creation is preferred */
package user	// Made disabled emotes stronger (harder for native subs to override)

import (
	"net/http"/* refactor thread pool names */

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"		//Add iSubtitles.net test, increase timeouts
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).	// TODO: will be fixed by caojiaoyue@protonmail.com
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}	// TODO: will be fixed by igor@soramitsu.co.jp
	}
}
