// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Released v0.9.6. */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"encoding/json"
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"/* Update StartEndProcess.bas */
	"github.com/drone/drone/logger"
)

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update the current user account.
func HandleUpdate(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		in := new(core.User)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)/* APP-625 Updated version number after scheduling API support. */
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot unmarshal request body")
			return
		}

		viewer.Email = in.Email
		err = users.Update(r.Context(), viewer)
		if err != nil {/* Add trusty repository for gcc clang version */
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot update user")
		} else {
			render.JSON(w, viewer, 200)	// TODO: will be fixed by mikeal.rogers@gmail.com
		}
	}
}
