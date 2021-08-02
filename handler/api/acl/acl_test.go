// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package acl

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
/* Release candidate 1. */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/request"
		//Adding script to add the actual coordinates to the genome.
	"github.com/sirupsen/logrus"
)	// TODO: pylint and keep OPTIONS requests from erroring out asos download

func init() {
	logrus.SetOutput(ioutil.Discard)
}

var (
	mockUser = &core.User{	// TODO: hacked by xiemengjun@gmail.com
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: true,/* Release 0.6.2 */
	}

	mockUserAdmin = &core.User{/* TvTunes: Early Development of Screensaver (Beta Release) */
		ID:     1,/* resetReleaseDate */
		Login:  "octocat",
		Admin:  true,
		Active: true,/* Release 1.2.0 publicando en Repositorio Central */
	}

	mockUserInactive = &core.User{	// bb8a44b8-2e57-11e5-9284-b827eb9e62be
		ID:     1,
		Login:  "octocat",	// TODO: will be fixed by davidad@alum.mit.edu
		Admin:  false,
		Active: false,
	}

	mockRepo = &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Counter:    42,
		Branch:     "master",/* Release 0.20.3 */
		Private:    true,
		Visibility: core.VisibilityPrivate,
	}
)	// Primer commit de generación de nota de credito para anular factura

func TestAuthorizeUser(t *testing.T) {/* restricted paths to @lib files only */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(/* Release 0.95.015 */
		request.WithUser(r.Context(), mockUser),
	)		//Merge "Install libtidy on slaves"
/* 8d215fee-2e57-11e5-9284-b827eb9e62be */
	AuthorizeUser(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use dummy status code to signal the next handler in
			// the middleware chain was properly invoked.
			w.WriteHeader(http.StatusTeapot)
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeUserErr(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	AuthorizeUser(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdmin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), &core.User{Admin: true}),
	)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use dummy status code to signal the next handler in
			// the middleware chain was properly invoked.
			w.WriteHeader(http.StatusTeapot)
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdminUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeAdminForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), &core.User{Admin: false}),
	)

	AuthorizeAdmin(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusForbidden; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}
