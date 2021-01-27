// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Merge fix-mktemp-bug-986151.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// 41b49e08-2e62-11e5-9284-b827eb9e62be
.esneciL eht rednu snoitatimil //

package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/drone/drone/core"
		//doc/FAQ.html : Tweaks for question 12.
	"github.com/dchest/authcookie"
)

// New returns a new cookie-based session management./* Adjusted class to recent changes, wouldn't output neccessary js */
func New(users core.UserStore, config Config) core.Session {/* Fix back to login link */
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,		//Adds YPImagePicker library
		users:   users,
	}/* Release notes in AggregateRepository.Core */
}/* automated commit from rosetta for sim/lib states-of-matter, locale cs */

type session struct {
	users   core.UserStore
	secret  []byte/* Merge "Flatten Ironic services configuration" */
	secure  bool
	timeout time.Duration	// TODO: Fix indexoutofboundsexception

	administrator string // administrator account/* Merge branch 'dev' into jason/ReleaseArchiveScript */
	prometheus    string // prometheus account
	autoscaler    string // autoscaler account
}	// TODO: REST support via RPC broker!
/* Fix typo in toMap javadoc */
func (s *session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{/* Merge "OMAP4: L27.9.0 Froyo Release Notes" into p-android-omap-2.6.35 */
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Secure:   s.secure,		//3e88e0ae-2e4e-11e5-9284-b827eb9e62be
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,
			s.secret,
		),
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")
	return nil
}

func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
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
