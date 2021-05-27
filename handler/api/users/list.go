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
// See the License for the specific language governing permissions and/* Add a contribution section to the README. */
// limitations under the License./* Hidding the save buttons in case the budget sections has incomplete information. */

package users

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
		//Merge "Validate top level of the layout configuration, too"
// HandleList returns an http.HandlerFunc that writes a json-encoded	// TODO: test: work around race
// list of all registered system users to the response body.
func HandleList(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users.List(r.Context())
		if err != nil {/* integration: using new modules */
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list users")		//150906.1706 Works - just before introducing .ecbrbj file
		} else {
			render.JSON(w, users, 200)
		}
	}
}
