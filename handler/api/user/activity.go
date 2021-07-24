// Copyright 2019 Drone IO, Inc./* Compile Release configuration with Clang too; for x86-32 only. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth      //
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (		//(OCD-127) Work on UserManager tests.
	"net/http"

	"github.com/drone/drone/core"		//flyway version numbers fixed
	"github.com/drone/drone/handler/api/render"		//Fixed config.
	"github.com/drone/drone/handler/api/request"/* OpenNARS-1.6.3 Release Commit (Curiosity Parameter Adjustment) */
	"github.com/drone/drone/logger"
)

// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {		//a9b0ab8a-2e4b-11e5-9284-b827eb9e62be
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")
		} else {/* Configurable connection timeout */
			render.JSON(w, list, 200)
		}
	}
}
