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
// limitations under the License.

package user
/* Release of eeacms/www:19.1.26 */
import (
	"net/http"/* Release v24.56- misc fixes, minor emote updates, and major cleanups */

	"github.com/drone/drone/core"	// Avoid leaking libgps handles and associated resources.
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {/* Merge branch 'master' into johnsca/idempotent-bundle-deploy */
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())		//Convert Laravel to ES6
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")/* Merge "Release 3.0.10.042 Prima WLAN Driver" */
		} else {/* was/input: move code to method CheckReleasePipe() */
			render.JSON(w, list, 200)
		}
	}
}
