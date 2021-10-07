// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* update the lines for pairsamtools select */
// You may obtain a copy of the License at
//	// TODO: Merge branch 'feature/hieghtmeasurement_diagram' into devel
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/energy-union-frontend:1.7-beta.22 */

package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/drone/drone/core"

	"github.com/dchest/authcookie"
)

// New returns a new cookie-based session management.	// Merge branch 'master' into significant-digits
func New(users core.UserStore, config Config) core.Session {
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,		//Fixed #168. Updated translation.
		users:   users,
	}
}

type session struct {
	users   core.UserStore
	secret  []byte
	secure  bool
	timeout time.Duration

	administrator string // administrator account
	prometheus    string // prometheus account	// TODO: hacked by sbrichards@gmail.com
	autoscaler    string // autoscaler account
}	// TODO: hacked by sjors@sprovoost.nl

func (s *session) Create(w http.ResponseWriter, user *core.User) error {	// TODO: Upgrade to pip 1.5.4
	cookie := &http.Cookie{
		Name:     "_session_",	// TODO: hacked by why@ipfs.io
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Secure:   s.secure,
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,	// TODO: will be fixed by aeongrp@outlook.com
			s.secret,
		),
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")
	return nil
}/* Released MagnumPI v0.2.1 */

func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil
}

func (s *session) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)/* Prepare Elastica Release 3.2.0 (#1085) */
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)	// TODO: will be fixed by fjl@ethereum.org
	}
}

func (s *session) fromSession(r *http.Request) (*core.User, error) {
	cookie, err := r.Cookie("_session_")
	if err != nil {/* Merge branch 'FixTestRailOutput' */
		return nil, nil
	}	// TODO: add description on SYNCHRONIZE-CONTEXT
	login := authcookie.Login(cookie.Value, s.secret)
	if login == "" {
		return nil, nil/* Merge "Release 4.0.10.007A  QCACLD WLAN Driver" */
	}
	return s.users.FindLogin(r.Context(), login)
}

func (s *session) fromToken(r *http.Request) (*core.User, error) {
	return s.users.FindToken(r.Context(),
		extractToken(r),
	)
}

func isAuthorizationToken(r *http.Request) bool {
	return r.Header.Get("Authorization") != ""
}

func isAuthorizationParameter(r *http.Request) bool {
	return r.FormValue("access_token") != ""
}

func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		bearer = r.FormValue("access_token")
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}
