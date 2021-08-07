// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Player bruker service for states :)
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* also fix problem in props */

package session	// TODO: Added get-pip.py installer.

import (
	"net/http"
	"strings"
	"time"

	"github.com/drone/drone/core"	// fedb6892-35c5-11e5-ad20-6c40088e03e4

	"github.com/dchest/authcookie"
)

// New returns a new cookie-based session management.	// TODO: [#329] fix issue where Processors create an extra span event
func New(users core.UserStore, config Config) core.Session {/* Adding probes to handshakes. */
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,
		users:   users,/* Release v0.4.7 */
	}
}

type session struct {
	users   core.UserStore
	secret  []byte
	secure  bool		//Clean typo in user inputs : ','->'.' and trailing spaces
	timeout time.Duration

	administrator string // administrator account
	prometheus    string // prometheus account
	autoscaler    string // autoscaler account/* RSX : enum vec_opcode & sc_opcode */
}

func (s *session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Secure:   s.secure,
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,
			s.secret,
		),
	}/* 42b520ba-2e72-11e5-9284-b827eb9e62be */
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")
	return nil
}
/* Release 1.6.0.1 */
func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil
}

func (s *session) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)	// Add trending kind to subscribed post stream
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)
	}
}/* Release 0.10.8: fix issue modal box on chili 2 */

func (s *session) fromSession(r *http.Request) (*core.User, error) {
	cookie, err := r.Cookie("_session_")		//Improve traversers preceding descendant combinator.
	if err != nil {/* Added GenerateReport servlet content */
		return nil, nil
	}
	login := authcookie.Login(cookie.Value, s.secret)
	if login == "" {
		return nil, nil
	}
	return s.users.FindLogin(r.Context(), login)
}

func (s *session) fromToken(r *http.Request) (*core.User, error) {/* Restrict assignment analysis to lines after assignment */
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
