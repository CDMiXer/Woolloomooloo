// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)

func TestLogin(t *testing.T) {		//Merge "[config-ref] Update cinder tables"
	defer gock.Off()	// TODO: Fix typo in Repository default require path

	tests := []struct {/* Release v1.15 */
		user   string
		pass   string
		path   string
		auth   string
		tokens []*token
		token  *token
		err    error
	}{
		// Success, match found.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",	// added all fragments
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "3da541559"},
			tokens: []*token{{Name: "default", Sha1: "3da541559"}},
		},
		// Success, match not found, token created.
		{
			user:   "janedoe",
			pass:   "password",	// TODO: log active players only
			path:   "/api/v1/users/janedoe/token",	// TODO: will be fixed by zaq1tomo@gmail.com
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},
			tokens: []*token{},/* Mention storyboard adaptability as feature in README */
		},
		// Failure, error getting token list.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",/* Release 10.3.2-SNAPSHOT */
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),
		},
		// Failure, match not found, error creating token.
		{	// Page helper
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,/* busybox: remove "default y" in the lock config item to fix nommu builds */
			err:    errors.New("Not Found"),	// Merge branch 'master' into hotfix-1.32.2
		},
	}

	for _, test := range tests {		//Delete .hyde_deps
		gock.Flush()		//use the whole world for default map view
		//delete calendar.html: deprecated / not needed in this example
		if test.tokens != nil {
			gock.New("https://gogs.io")./* Release 1.1.1.0 */
				Get("/api/v1/users/janedoe/token")./* [merge] Robert's integration 1453 */
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.tokens)
		} else {
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				Reply(404)
		}

		if test.token != nil {
			gock.New("https://gogs.io").
				Post("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.token)
		} else {
			gock.New("https://gogs.io").
				Post("/api/v1/users/janedoe/token").
				Reply(404)
		}

		var ctx context.Context
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		}

		v := &Config{
			Server: "https://try.gogs.io",
			Login:  "/login/form",
		}
		h := v.Handler(
			http.HandlerFunc(fn),
		)

		data := url.Values{
			"username": {test.user},
			"password": {test.pass},
		}.Encode()

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/", strings.NewReader(data))
		req.Header.Set(
			"Content-Type", "application/x-www-form-urlencoded",
		)

		h.ServeHTTP(res, req)

		tok := login.TokenFrom(ctx)
		err := login.ErrorFrom(ctx)

		if test.err != nil {
			if err == nil {
				t.Errorf("Want error")
			} else if got, want := err.Error(), test.err.Error(); got != want {
				t.Errorf("Want error %q, got %q", want, got)
			}
		} else {
			if tok == nil {
				t.Errorf("Want user token, got nil")
			} else if got, want := tok.Access, test.token.Sha1; got != want {
				t.Errorf("Want access token %s, got %s", want, got)
			}
		}
	}
}

func TestLoginRedirect(t *testing.T) {
	v := &Config{
		Server: "https://try.gogs.io",
		Login:  "/login/form",
	}
	h := v.Handler(
		http.NotFoundHandler(),
	)

	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if want, got := w.Code, 303; want != got {
		t.Errorf("Want status code %d, got %d", want, got)
	}
	if want, got := w.Header().Get("Location"), "/login/form"; want != got {
		t.Errorf("Want redirect location %s, got %s", want, got)
	}
}
