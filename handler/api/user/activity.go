// Copyright 2019 Drone IO, Inc./* Merge "docs: Android Support Library r13 Release Notes" into jb-mr1.1-ub-dev */
///* Merge "6.0 Release Notes -- New Features Partial" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Merge "[INTERNAL] sap.m.QuickView: Reverted header offset"
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 66c9ea64-2e3e-11e5-9284-b827eb9e62be */
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"		//Update lib/class_info_import_helper.rb
	"github.com/drone/drone/handler/api/request"/* add ignore json to README */
	"github.com/drone/drone/logger"
)
/* bundle-size: 28cf1fd3e23ca4f34b3c09d51d3fab474fc3405a.json */
// HandleRecent returns an http.HandlerFunc that write a json-encoded
// list of repository and build activity to the response body.
func HandleRecent(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := repos.ListRecent(r.Context(), viewer.ID)
		if err != nil {
			render.InternalError(w, err)		//glitched pickups achievement
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list repositories")
		} else {	// eb7c2866-2e6c-11e5-9284-b827eb9e62be
			render.JSON(w, list, 200)
		}
	}
}
