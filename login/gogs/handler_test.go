// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* 4.7.0 Release */
// license that can be found in the LICENSE file.
/* Release 1.0.0 (#293) */
package gogs

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"		//Fix broken images for dwarf and serpent mob
	"strings"	// TODO: Line width via quad.
	"testing"

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)
	// TODO: hacked by arajasek94@gmail.com
func TestLogin(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		user   string
		pass   string
		path   string/* Update metadata_managment.md */
		auth   string
		tokens []*token
		token  *token		//Update sqlConnect.js
		err    error
	}{
		// Success, match found.
		{/* Added screens */
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "3da541559"},
			tokens: []*token{{Name: "default", Sha1: "3da541559"}},
		},
		// Success, match not found, token created.
		{
			user:   "janedoe",
			pass:   "password",	// MusicChunk: initialize replay_gain_serial on demand
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},		//Add a missing Page attribute to fix broken mkdocs build
			tokens: []*token{},/* PY3: Issues with sorting dicts. Tiny API change to avoid shadowing a builtin */
		},
		// Failure, error getting token list./* You can SAVE !!!! Add title support. */
		{/* Change Release language to Version */
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),
		},
		// Failure, match not found, error creating token.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,
			err:    errors.New("Not Found"),
		},
	}
	// TODO: hacked by caojiaoyue@protonmail.com
	for _, test := range tests {
		gock.Flush()

		if test.tokens != nil {	// TODO: hacked by hello@brooklynzelenka.com
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.tokens)	// TODO: hacked by hi@antfu.me
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
