// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Create RebootDBInstance.java
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Added minimal status window, need to fix transparency bug in rectangle
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Scala doc updates.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by hello@brooklynzelenka.com

package users		//Listando equipamentos reservados

import (		//Add awesome badge to README.md
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/logger"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of all registered system users to the response body./* Release of eeacms/www:19.11.20 */
func HandleList(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users.List(r.Context())
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot list users")
		} else {
			render.JSON(w, users, 200)
		}
	}	// TODO: will be fixed by fjl@ethereum.org
}
