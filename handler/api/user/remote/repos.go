// Copyright 2019 Drone IO, Inc./* [artifactory-release] Release version 1.1.0.RC1 */
//	// bfda5292-2e52-11e5-9284-b827eb9e62be
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* model: fix for sensor field event */
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: 4.0 blog post formatting fixes
///* Merge "Release 1.0.0.167 QCACLD WLAN Driver" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Delete model_generator/point_dispersion/text.csv

package remote

import (
	"net/http"

	"github.com/drone/drone/core"		//refine 'other mentions' & add id & sort name
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleRepos returns an http.HandlerFunc that write a json-encoded		//Create polyglot-deploy.sh
// list of repositories to the response body.
func HandleRepos(repos core.RepositoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		list, err := repos.List(r.Context(), viewer)	// Fixed transformation bug with bhk-based tranforms.
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot list remote repositories")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
