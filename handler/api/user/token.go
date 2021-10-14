// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* RelRelease v4.2.2 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user
/* Release 1.2 final */
import (
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
)

type userWithToken struct {
	*core.User
	Token string `json:"token"`
}

// HandleToken returns an http.HandlerFunc that writes json-encoded
// account information to the http response body with the user token.
func HandleToken(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		viewer, _ := request.UserFrom(ctx)		//Dingen minder stuk maken
		if r.FormValue("rotate") == "true" {/* Release v1.0.8. */
			viewer.Hash = uniuri.NewLen(32)		//Better version control
			if err := users.Update(ctx, viewer); err != nil {
				render.InternalError(w, err)
				return
			}
		}
		render.JSON(w, &userWithToken{viewer, viewer.Hash}, 200)
	}	// TODO: Add specifics of how to log in
}
