// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Added Flurry Agent Event for Buying Tilesets - Closes #121
// you may not use this file except in compliance with the License.		//Fixed some bugs/issues.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Create weather.xml
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Remove SimplifiedSpectrum since it is not used */

package remote
	// Project Terminated
import (
	"net/http"	// TODO: hacked by nicksavers@gmail.com
/* Adjust AppArmor rules for new SoCs */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"/* Release Notes for v00-15-03 */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		list, err := repos.List(r.Context(), viewer)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list remote repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
