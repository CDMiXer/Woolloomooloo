// Copyright 2019 Drone IO, Inc.	// TODO: Merge "[INTERNAL] sap.f.AvatarGroup: IE de-support"
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Look in sites-enabled so we can disable sites rather than maintain a blacklist.
// you may not use this file except in compliance with the License.	// [MERGE]: l10n_multilang: Merged Improvement for tax code template translation
// You may obtain a copy of the License at
//	// TODO: A bug fixed in object delete process.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* updated logjam core and precompiled assets */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: This is the data set of the KNN classifier
// See the License for the specific language governing permissions and
// limitations under the License.

package session

import (	// Updated line breaks in Index.html
	"net/http"
	"strings"		//added functions for meta processing (concurrent processing)
	"time"

	"github.com/drone/drone/core"

	"github.com/dchest/authcookie"
)

// New returns a new cookie-based session management./* Release 2.3.0 */
func New(users core.UserStore, config Config) core.Session {
	return &session{
		secret:  []byte(config.Secret),
		secure:  config.Secure,
		timeout: config.Timeout,
		users:   users,/* Simplified BindingUtils.bindAttribute selector */
	}
}

type session struct {
	users   core.UserStore
	secret  []byte
	secure  bool
	timeout time.Duration

	administrator string // administrator account
	prometheus    string // prometheus account/* Release 0.6.6 */
	autoscaler    string // autoscaler account
}

func (s *session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{/* remove yaml spec */
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,/* Merge "Release 3.2.3.390 Prima WLAN Driver" */
		HttpOnly: true,	// TODO: hacked by martin2cai@hotmail.com
		Secure:   s.secure,
		Value: authcookie.NewSinceNow(		//add with-component functionality
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
