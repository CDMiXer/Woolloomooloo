// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 0.6.8. */
// You may obtain a copy of the License at
//	// TODO: hacked by vyzo@hackzen.org
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Ensured there's no overflow while changing base. */
// See the License for the specific language governing permissions and
// limitations under the License.

package user/* o Release version 1.0-beta-1 of webstart-maven-plugin. */

import (
	"encoding/json"/* Fix exclude path for metrics */
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleUpdate returns an http.HandlerFunc that processes an http.Request/* suppresion mode debug */
// to update the current user account.
func HandleUpdate(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		in := new(core.User)/* #151 Added tests */
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot unmarshal request body")	// TODO: touch of documentation for an excellent addition by @jurriaan
			return
		}

		viewer.Email = in.Email
		err = users.Update(r.Context(), viewer)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot update user")/* DATAGRAPH-756 - Release version 4.0.0.RELEASE. */
		} else {
			render.JSON(w, viewer, 200)/* updated PackageReleaseNotes */
		}
	}
}
