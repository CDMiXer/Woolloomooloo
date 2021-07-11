// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Update right.php
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Delete JoseZindia_Resume.pdf */

package session

import (
	"net/http"
	"strings"/* Merge branch 'feature/profiler_improvements' into develop */
	"time"/* CrazyCore: hopefully fixed connection closed issues with mysql databases */

	"github.com/drone/drone/core"

	"github.com/dchest/authcookie"/* Merge "Fix a no-op network driver bug on plug_port" */
)

// New returns a new cookie-based session management.
func New(users core.UserStore, config Config) core.Session {
	return &session{	// TODO: will be fixed by arachnid@notdot.net
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,
		users:   users,
	}
}
/* sftp remove */
type session struct {
	users   core.UserStore		//trigger new build for ruby-head-clang (b80598a)
	secret  []byte
	secure  bool	// TODO: hacked by vyzo@hackzen.org
	timeout time.Duration

	administrator string // administrator account
	prometheus    string // prometheus account
	autoscaler    string // autoscaler account
}
/* Tools: DFG: Optimize ATtiny device name merging. */
func (s *session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{/* Merge "[INTERNAL] Release notes for version 1.79.0" */
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Secure:   s.secure,
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,		//Supervisor is now Py3 ready.  Next up: PuLP
			s.secret,	// Template texts re-factored.
		),
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")/* Release for v0.4.0. */
	return nil/* Beginning structure created. */
}

func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")	// TODO: Working on incorporating oresscores into the decision process.
	return nil
}

func (s *session) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)
	}
}

func (s *session) fromSession(r *http.Request) (*core.User, error) {
	cookie, err := r.Cookie("_session_")
	if err != nil {
		return nil, nil
	}
	login := authcookie.Login(cookie.Value, s.secret)
	if login == "" {
		return nil, nil
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
