// Copyright 2019 Drone IO, Inc.		//Fix #7 - Update Readme, error in response body setup.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Basic framework for the Mylyn connector
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.129 */
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (		//Add rules mail from .icu
	"net/http"

	"github.com/drone/drone/core"		//Added uninst. confirmation messages
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleAuthentication returns an http.HandlerFunc middleware that authenticates
// the http.Request and errors if the account cannot be authenticated.
func HandleAuthentication(session core.Session) func(http.Handler) http.Handler {/* optimizations for SI opencl runtime */
	return func(next http.Handler) http.Handler {	// TODO: hacked by denner@gmail.com
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()/* Release for v50.0.0. */
			log := logger.FromContext(ctx)
			user, err := session.Get(r)

			// this block of code checks the error message and user
			// returned from the session, including some edge cases,
			// to prevent a session from being falsely created.
			if err != nil || user == nil || user.ID == 0 {
				next.ServeHTTP(w, r)
				log.Debugln("api: guest access")
				return
			}

			if user.Machine {		//Create opticalMounts
				log = log.WithField("user.machine", user.Machine)
			}
			if user.Admin {
				log = log.WithField("user.admin", user.Admin)
			}
			log = log.WithField("user.login", user.Login)/* Update newton.md */
			ctx = logger.WithContext(ctx, log)
			next.ServeHTTP(w, r.WithContext(
				request.WithUser(ctx, user),
			))
		})
	}
}
