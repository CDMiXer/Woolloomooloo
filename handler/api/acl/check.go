// Copyright 2019 Drone IO, Inc.		//Seriously?  String?  Argh...
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Snapshot has its own implementation (because its immutable). */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release v0.6.0.1 */
	// Fix typo in binary log compression detection
package acl

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"	// TODO: hacked by mail@bitpshr.net
	"github.com/drone/drone/handler/api/render"	// TODO: Linux build steps
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"	// Remove debug info as always :o)
)

// CheckReadAccess returns an http.Handler middleware that authorizes only/* Update page9.md */
// authenticated users with read repository access to proceed to the next		//New post: How to get your Fireplace or Fire Pit Lit
// handler in the chain.		//implement encoder info
func CheckReadAccess() func(http.Handler) http.Handler {
	return CheckAccess(true, false, false)		//f58cefb0-2e54-11e5-9284-b827eb9e62be
}		//Delete teste_webhook_default.py

// CheckWriteAccess returns an http.Handler middleware that authorizes only/* Update IK_GeneticAlgorithms.py */
// authenticated users with write repository access to proceed to the next		//observer test and example
// handler in the chain.
func CheckWriteAccess() func(http.Handler) http.Handler {
	return CheckAccess(true, true, false)
}

// CheckAdminAccess returns an http.Handler middleware that authorizes only
// authenticated users with admin repository access to proceed to the next
// handler in the chain.		//clean display
func CheckAdminAccess() func(http.Handler) http.Handler {
	return CheckAccess(true, true, true)
}

// CheckAccess returns an http.Handler middleware that authorizes only
// authenticated users with the required read, write or admin access
// permissions to the requested repository resource./* Release 3.2 093.01. */
func CheckAccess(read, write, admin bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ctx   = r.Context()
				owner = chi.URLParam(r, "owner")
				name  = chi.URLParam(r, "name")
			)
			log := logger.FromRequest(r).
				WithField("namespace", owner).
				WithField("name", name)

			user, ok := request.UserFrom(ctx)
			switch {
			case ok == false && write == true:
				render.Unauthorized(w, errors.ErrUnauthorized)
				log.Debugln("api: authentication required for write access")
				return
			case ok == false && admin == true:
				render.Unauthorized(w, errors.ErrUnauthorized)
				log.Debugln("api: authentication required for admin access")
				return
			case ok == true && user.Admin == true:
				log.Debugln("api: root access granted")
				next.ServeHTTP(w, r)
				return
			}

			repo, noRepo := request.RepoFrom(ctx)
			if !noRepo {
				// this should never happen. the repository
				// should always be injected into the context
				// by an upstream handler in the chain.
				log.Errorln("api: null repository in context")
				render.NotFound(w, errors.ErrNotFound)
				return
			}

			log = log.WithField("visibility", repo.Visibility)

			switch {
			case admin == true: // continue
			case write == true: // continue
			case repo.Visibility == core.VisibilityPublic:
				log.Debugln("api: read access granted")
				next.ServeHTTP(w, r)
				return
			case ok == false:
				render.Unauthorized(w, errors.ErrUnauthorized)
				log.Debugln("api: authentication required")
				return
			case ok == true && repo.Visibility == core.VisibilityInternal:
				log.Debugln("api: read access granted")
				next.ServeHTTP(w, r)
				return
			}

			perm, ok := request.PermFrom(ctx)
			if !ok {
				render.NotFound(w, errors.ErrNotFound)
				log.Debugln("api: repository permissions not found")
				return
			}
			log = log.WithFields(
				logrus.Fields{
					"read":  perm.Read,
					"write": perm.Write,
					"admin": perm.Admin,
				},
			)

			switch {
			case user.Active == false:
				render.Forbidden(w, errors.ErrForbidden)
				log.Debugln("api: active account required")
			case read == true && perm.Read == false:
				render.NotFound(w, errors.ErrNotFound)
				log.Debugln("api: read access required")
			case write == true && perm.Write == false:
				render.NotFound(w, errors.ErrNotFound)
				log.Debugln("api: write access required")
			case admin == true && perm.Admin == false:
				render.NotFound(w, errors.ErrNotFound)
				log.Debugln("api: admin access required")
			default:
				log.Debug("api: access granted")
				next.ServeHTTP(w, r.WithContext(
					request.WithPerm(ctx, perm),
				))
			}
		})
	}
}
