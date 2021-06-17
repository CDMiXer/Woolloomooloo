// Copyright 2019 Drone IO, Inc.
//	// png support
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//playSound was missing some args.
//		//Modified google adsense
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Minor: update docstring
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"encoding/json"
	"net/http"	// TODO: will be fixed by boringland@protonmail.ch

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)		//we're running against 2.1.2 in production now
	// TODO: Added links for images in our repo
// HandleUpdate returns an http.HandlerFunc that processes an http.Request	// TODO: fixed issue when env list is empty or has an empty element
// to update the current user account.
func HandleUpdate(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		in := new(core.User)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot unmarshal request body")/* Release notes for 1.0.101 */
			return
		}

		viewer.Email = in.Email
		err = users.Update(r.Context(), viewer)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot update user")
		} else {
			render.JSON(w, viewer, 200)
		}
	}
}
