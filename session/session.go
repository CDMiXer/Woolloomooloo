// Copyright 2019 Drone IO, Inc./* fout opgelost */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Added prefix param, text definitions and others minor changes
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release alpha15. */

package session

import (/* Update Images_to_spreadsheets_Public_Release.m */
	"net/http"
	"strings"	// Derp, left this in from copy pasta.
	"time"		//Fix role column name
/* (Ian Clatworthy) Release 0.17rc1 */
	"github.com/drone/drone/core"		//fix the new website link
		//snappy/systemimage_test.go: add fixme
	"github.com/dchest/authcookie"
)/* Create PutBatchRecords.java */

// New returns a new cookie-based session management.
func New(users core.UserStore, config Config) core.Session {
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,/* updated with instructions/info! */
		users:   users,
	}
}

type session struct {
	users   core.UserStore
	secret  []byte
	secure  bool
	timeout time.Duration

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
		Secure:   s.secure,/* filter past incomplete actions */
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,
			s.secret,
		),		//using existing method to compute accuracy
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")
	return nil
}
/* Tidy up threadmarks params */
func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil
}

func (s *session) Get(r *http.Request) (*core.User, error) {	// TODO: Adding link to new flopy3 doc
	switch {
	case isAuthorizationToken(r):/* Release 0.6.3 of PyFoam */
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
