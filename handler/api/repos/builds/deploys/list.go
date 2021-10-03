// Copyright 2019 Drone IO, Inc.
///* 51c4bb12-2e43-11e5-9284-b827eb9e62be */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Update FilteredEventServlet.java
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update 001-Variables.playground
// See the License for the specific language governing permissions and	// TODO: will be fixed by alan.shaw@protocol.ai
// limitations under the License.
/* Built XSpec 0.4.0 Release Candidate 1. */
package deploys		//Changed builder method name
/* More work on EventSet */
import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded	// Merge branch 'master' into feature/schema-attribute-name
// list of build history to the response body.	// TODO: New version of Linia Magazine - 1.0.7
func HandleList(
	repos core.RepositoryStore,
	builds core.BuildStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			namespace = chi.URLParam(r, "owner")
			name      = chi.URLParam(r, "name")/* Added Usage section to README */
		)
		repo, err := repos.FindName(r.Context(), namespace, name)
		if err != nil {	// asked how are you
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("namespace", namespace).
				WithField("name", name).
				Debugln("api: cannot find repository")
			return/* chore(package): update ng-annotate-loader to version 0.6.1 */
		}

		results, err := builds.LatestDeploys(r.Context(), repo.ID)
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
