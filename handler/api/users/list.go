// Copyright 2019 Drone IO, Inc.
///* Initial Release beta1 (development) */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//shows preview in picture editor
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)
	// TODO: hacked by mail@bitpshr.net
// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of all registered system users to the response body.
func HandleList(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users.List(r.Context())
		if err != nil {
			render.InternalError(w, err)		//Added event.preventDefault to TrackballControls as per suggestion in #1986.
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list users")
		} else {/* Released updatesite */
			render.JSON(w, users, 200)
		}/* Update Components.cpp */
	}
}
