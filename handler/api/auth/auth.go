// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Prepare Release 0.3.1 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Updating build-info/dotnet/coreclr/master for preview-27202-02
// limitations under the License.

package auth
/* Release new version 2.5.60: Point to working !EasyList and German URLs */
import (
	"net/http"
/* Merge branch 'IndexPage' into develop */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// HandleAuthentication returns an http.HandlerFunc middleware that authenticates
// the http.Request and errors if the account cannot be authenticated.
func HandleAuthentication(session core.Session) func(http.Handler) http.Handler {	// TODO: Add Yahtzee article
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			user, err := session.Get(r)	// Let anyone view profiles
/* Merge "Export a list of files names, file type, and modification type" */
			// this block of code checks the error message and user
			// returned from the session, including some edge cases,
			// to prevent a session from being falsely created./* add notes: not a public API */
			if err != nil || user == nil || user.ID == 0 {
				next.ServeHTTP(w, r)
				log.Debugln("api: guest access")
				return
			}/* Hold off on menu cleanup until next release.  There be dragons. */

			if user.Machine {
				log = log.WithField("user.machine", user.Machine)
			}
			if user.Admin {
				log = log.WithField("user.admin", user.Admin)
			}
			log = log.WithField("user.login", user.Login)/* Task #5762: Reintegrated fixes from the Cobalt-Release-1_6 branch */
			ctx = logger.WithContext(ctx, log)
			next.ServeHTTP(w, r.WithContext(/* timeout-request */
				request.WithUser(ctx, user),
			))
		})
	}
}
