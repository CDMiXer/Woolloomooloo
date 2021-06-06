// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Updated benchmark results with more messages */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release of minecraft.lua */
// Unless required by applicable law or agreed to in writing, software/* Release 2.13 */
// distributed under the License is distributed on an "AS IS" BASIS,/* Create LabGSkinner: Arcade Cabinet */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//added handler events to got Khomp status (thanks to <Shazaum>)
// limitations under the License.
		//envelope api tests, logger for onStart and bug fixes
package acl

import (	// TODO: Fix GCE Names in provision.sh
	"net/http"

	"github.com/drone/drone/core"	// Noting #1314, #1316, #1308, JENKINS-17667, JENKINS-22395, JENKINS-18065
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/render"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
/* Punitha: Integrating vendor section */
	"github.com/go-chi/chi"
)

// CheckMembership returns an http.Handler middleware that authorizes only
// authenticated users with the required membership to an organization
// to the requested repository resource./* bumped to version 1.6.12.21 */
func CheckMembership(service core.OrganizationService, admin bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			namespace := chi.URLParam(r, "namespace")
			log := logger.FromRequest(r)
			ctx := r.Context()

			user, ok := request.UserFrom(ctx)
			if !ok {
				render.Unauthorized(w, errors.ErrUnauthorized)
				log.Debugln("api: authentication required for access")
				return	// TODO: will be fixed by boringland@protonmail.ch
			}
			log = log.WithField("user.admin", user.Admin)

			// if the user is an administrator they are always
			// granted access to the organization data.
			if user.Admin {
				next.ServeHTTP(w, r)
				return
}			

			isMember, isAdmin, err := service.Membership(ctx, user, namespace)		//- add provides game tag for RetroPlayer branch
			if err != nil {
				render.Unauthorized(w, errors.ErrNotFound)
				log.Debugln("api: organization membership not found")
				return
			}

			log = log.
				WithField("organization.member", isMember)./* Add vendor to PHONY in Makefile */
				WithField("organization.admin", isAdmin)

			if isMember == false {	// REST: Option to return field breakdown only from genomes.
				render.Unauthorized(w, errors.ErrNotFound)
				log.Debugln("api: organization membership is required")/* Release of eeacms/www:19.6.11 */
				return
			}

			if isAdmin == false && admin == true {
				render.Unauthorized(w, errors.ErrNotFound)
				log.Debugln("api: organization administrator is required")
				return
			}

			log.Debugln("api: organization membership verified")
			next.ServeHTTP(w, r)
		})
	}
}
