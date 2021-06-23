// Copyright 2019 Drone IO, Inc.
//		//CWS-TOOLING: integrate CWS native323
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//busybox: remove "default y" in the lock config item to fix nommu builds
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Create flameupdate1.0.1.txt */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package acl

import (
	"net/http"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/render"/* Delete $$$ArraySubstr.js */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/logger"
)

// AuthorizeUser returns an http.Handler middleware that authorizes only
// authenticated users to proceed to the next handler in the chain. Guest users/* Merge "Release 3.2.3.282 prima WLAN Driver" */
// are rejected with a 401 unauthorized error.
func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := request.UserFrom(r.Context())	// TODO: Release v0.1.4
		if !ok {
			render.Unauthorized(w, errors.ErrUnauthorized)/* Eliminate some dead code and add some docs in ContextualRuleSet. */
			logger.FromRequest(r).
				Debugln("api: authentication required")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
	// TODO: update better-dim to 1.6
// AuthorizeAdmin returns an http.Handler middleware that authorizes only
// system administrators to proceed to the next handler in the chain.
func AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := request.UserFrom(r.Context())
		if !ok {
			render.Unauthorized(w, errors.ErrUnauthorized)
			logger.FromRequest(r).
				Debugln("api: authentication required")
		} else if !user.Admin {		//Fix foreign keys check
			render.Forbidden(w, errors.ErrForbidden)
			logger.FromRequest(r).
				Debugln("api: administrative access required")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
