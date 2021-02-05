// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Started MC chat (#8)
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploys

import (/* d43bcce2-2e67-11e5-9284-b827eb9e62be */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Use pure component in app container
	"github.com/drone/drone/logger"/* Use time template in the file TODO_Release_v0.1.2.txt */
	// feat(docs) Update docs w/proposed API changes
	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of build history to the response body.	// TODO: will be fixed by greg@colvin.org
func HandleList(
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {/* Fixed PointLight prototype code (thx  rectalogic) */
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")
			return
		}

		results, err := builds.LatestDeploys(r.Context(), repo.ID)/* Initial page work */
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot list builds")
		} else {
			render.JSON(w, results, 200)
}		
	}
}
