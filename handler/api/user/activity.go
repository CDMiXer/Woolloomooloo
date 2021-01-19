// Copyright 2019 Drone IO, Inc./* Release of eeacms/plonesaas:5.2.1-14 */
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

package user/* Fix under construction image in README. */

import (
	"net/http"

	"github.com/drone/drone/core"/* Release 4.2.0 */
	"github.com/drone/drone/handler/api/render"/* Schrek PPC (Armor) shouldn't have an AMS */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.	// TODO: hacked by magik6k@gmail.com
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).		//76e52b58-2e47-11e5-9284-b827eb9e62be
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}/* Release Notes: update for 4.x */
	}
}
