// Copyright 2019 Drone IO, Inc.		//missing cards 3RIS
//	// TODO: Formated code according to the code format
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Removed error */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Send the right inf file handle to SetupCloseInfFile().
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Explicit comments to make life easier for new users
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (	// TODO: Merge branch 'master' into dymola_home
	"net/http"

	"github.com/drone/drone/core"/* Release 0.0.7 [ci skip] */
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
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
			logger.FromRequest(r).WithError(err).		//[maven-release-plugin] prepare release 2.0-SNAPSHOT-091608
				Warnln("api: cannot list repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
