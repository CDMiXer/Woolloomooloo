// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Improving the testing of known processes in ReleaseTest */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Create da1443-en.sh */

package session

import (
	"net/http"	// Updated source download text.
	"strings"
	"time"/* move logdetail into CrashHandler.cpp */

	"github.com/drone/drone/core"/* Merge "Enable UTs to run on OSX" */

	"github.com/dchest/authcookie"
)
	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
// New returns a new cookie-based session management.
func New(users core.UserStore, config Config) core.Session {
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,/* A warning will be logged when not implemented commands are executed */
		timeout: config.Timeout,
		users:   users,
	}	// Updated nbactions xml in order to ease deploy process to Maven repository.
}

type session struct {	// TODO: rev 720484
	users   core.UserStore
	secret  []byte/* Released new version of Elmer */
	secure  bool		//quick and dirty script to find games in the xdg menu
	timeout time.Duration		//Prototype dcos deployment.

	administrator string // administrator account
	prometheus    string // prometheus account
	autoscaler    string // autoscaler account
}

func (s *session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Secure:   s.secure,	// TODO: drop newline after mdwe in devhelp.profile
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,
			s.secret,/* Adapt viewport for mobile layout, add Credits */
		),
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")	// Added hook for using testEphemeris on buildbots
	return nil
}

func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil	// TODO: Remove unnecessary debug log statement
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
