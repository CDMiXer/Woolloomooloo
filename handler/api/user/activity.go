// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update what.md */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* We don't use PECL anymore */
package user
/* Create myipas.c */
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"/* Adding default inits */
)	// oozie client: doc about oozie command

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}/* Creation of Release 1.0.3 jars */
