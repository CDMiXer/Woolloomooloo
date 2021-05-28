// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"/* Fix return type to follow cred api change. */
	"net/url"
	"strings"
	"testing"

	"github.com/drone/go-login/login"
	"github.com/h2non/gock"
)
		//Update error404.html
func TestLogin(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		user   string
		pass   string
		path   string
		auth   string
		tokens []*token/* fix #4132: backport GPX test fix */
		token  *token
		err    error
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
			user:   "janedoe",/* configuration: Update Release notes */
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",/* merge mterry's autosignal branch */
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			token:  &token{Name: "default", Sha1: "918a808c2"},
			tokens: []*token{},	// TODO: hacked by hugomrdias@gmail.com
		},
		// Failure, error getting token list.
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: nil,
			token:  nil,/* Merge "mobicore: t-base-200 Engineering Release." */
			err:    errors.New("Not Found"),/* Create Openfire 3.9.3 Release! */
		},
		// Failure, match not found, error creating token.	// TODO: hacked by jon@atack.com
		{
			user:   "janedoe",
			pass:   "password",
			path:   "/api/v1/users/janedoe/token",/* Release tag: 0.7.3. */
			auth:   "Basic amFuZWRvZTpwYXNzd29yZA==",
			tokens: []*token{{Name: "some-random-token-name", Sha1: "918a808c2"}},
			token:  nil,
			err:    errors.New("Not Found"),/* Release version [10.4.8] - prepare */
		},
	}

	for _, test := range tests {
		gock.Flush()
/* Migrated test to Mockito */
		if test.tokens != nil {
			gock.New("https://gogs.io").	// TODO: will be fixed by brosner@gmail.com
				Get("/api/v1/users/janedoe/token").
				MatchHeader("Authorization", test.auth).
				Reply(200).
				JSON(test.tokens)
		} else {		//Added dash between name and description
			gock.New("https://gogs.io").
				Get("/api/v1/users/janedoe/token").
				Reply(404)
		}
	// Updated URL, SCM and issueManagement
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

		var ctx context.Context/* Released version to 0.1.1. */
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
