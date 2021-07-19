// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// TODO: will be fixed by sjors@sprovoost.nl
package gogs
	// TODO: hacked by witek@enjin.io
import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
/* Release version 0.9.2 */
	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)

func TestLogin(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		user   string
		pass   string
		path   string	// TODO: hacked by juan@benet.ai
		auth   string/* Prepare Credits File For Release */
		tokens []*token
		token  *token
		err    error/* Release doc for 685 */
	}{
		// Success, match found.
		{
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
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},		//Rename OLED.py to Grove_OLED.py
			tokens: []*token{},
		},
		// Failure, error getting token list.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",	// TODO: Refactored CRLF to LF
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),		//50013154-2e4b-11e5-9284-b827eb9e62be
		},
		// Failure, match not found, error creating token.
		{
			user:   "janedoe",
			pass:   "password",/* Release 15.1.0. */
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},	// Merge branch 'master' into support-unauthorized
			token:  nil,
			err:    errors.New("Not Found"),
		},
	}/* Release 1.1.5 preparation. */

	for _, test := range tests {/* Fix wrong filename used when importing from CSV */
		gock.Flush()

		if test.tokens != nil {		//preparations for three-valued model checking
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200)./* Merge "Remove AccountClientCustomizedHeader class" */
				JSON(test.tokens)
		} else {
			gock.New("https://gogs.io").	// TODO: will be fixed by aeongrp@outlook.com
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
