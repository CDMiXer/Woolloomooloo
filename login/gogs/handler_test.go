// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* Merge "Merge "ASoC: msm: qdsp6v2: Release IPA mapping"" */
// license that can be found in the LICENSE file.
		//80a6332c-2e41-11e5-9284-b827eb9e62be
package gogs

import (
	"context"
	"errors"
	"net/http"/* Release 3.0.6. */
	"net/http/httptest"	// TODO: will be fixed by ng8eke@163.com
	"net/url"
	"strings"
	"testing"

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)

func TestLogin(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		user   string
		pass   string
		path   string
		auth   string/* Interface folder changed to interface */
		tokens []*token
		token  *token/* Release 1.5.3. */
		err    error
	}{
		// Success, match found.
		{/* Release version 0.10. */
			user:   "janedoe",/* UI: Lisätty list/info alinäkymään linkit harjoitusohjelmaan ja harjoituspohjaan */
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",	// TODO: Update api_spec.rb
			token:  &token{Name: "default", Sha1: "3da541559"},
			tokens: []*token{{Name: "default", Sha1: "3da541559"}},/* e6f3203e-4b19-11e5-97e4-6c40088e03e4 */
		},
		// Success, match not found, token created.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},
			tokens: []*token{},
		},		//Tweak yaml
		// Failure, error getting token list.
		{
			user:   "janedoe",
			pass:   "password",/* Delete ReleaseNotes.txt */
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",/* Clarify Gallery description */
			tokens: nil,
			token:  nil,
			err:    errors.New("Not Found"),
		},		//Fixed serial date widget, Friday was missing in the weekday list.
		// Failure, match not found, error creating token.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",	// TODO: hacked by alex.gaynor@gmail.com
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,
			err:    errors.New("Not Found"),/* Testing docker */
		},
	}

	for _, test := range tests {
		gock.Flush()

		if test.tokens != nil {
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
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
