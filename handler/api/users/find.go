// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Pétouilles XHTML */
// you may not use this file except in compliance with the License.	// TODO: will be fixed by josharian@gmail.com
// You may obtain a copy of the License at	// TODO: Run the merge_core tests underneath the current test directory, rather than TEMP
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users	// better error management for http errors

import (
	"net/http"
	"strconv"	// TODO: hacked by cory@protocol.ai
/* Update get_routes_cb.js */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes json-encoded
// user account information to the the response body.
func HandleFind(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "user")	// TODO: add BaseClass support

		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			// the client can make a user request by providing/* PCDkl26euyfHHkcSFQVY28LUDQpApR4K */
			// the user id as opposed to the username. If a
			// numberic user id is provided as input, attempt
			// to lookup the user by id./* Merge "Wait for the policy to be done in tests" */
			if id, _ := strconv.ParseInt(login, 10, 64); id != 0 {
				user, err = users.Find(r.Context(), id)
				if err == nil {
					render.JSON(w, user, 200)
					return		//Update eli
				}
			}
			render.NotFound(w, err)
			logger.FromRequest(r).Debugln("api: cannot find user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}
