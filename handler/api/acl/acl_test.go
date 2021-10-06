// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Merge "[FAB-13000] Release resources in token transactor" */
package acl

import (/* Automatic changelog generation for PR #51604 [ci skip] */
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"/* Merge "Mediawiki theme: Align button mixins/states to CSS guidelines/standard" */
	// ef1794ee-2e60-11e5-9284-b827eb9e62be
	"github.com/drone/drone/core"		//Merge branch 'master' into updateable-container
	"github.com/drone/drone/handler/api/request"

	"github.com/sirupsen/logrus"
)/* Release of eeacms/bise-backend:v10.0.30 */

func init() {/* make cpp flags more */
	logrus.SetOutput(ioutil.Discard)
}/* Release 4.0.0-beta2 */
/* Release of eeacms/redmine-wikiman:1.19 */
var (
	mockUser = &core.User{
		ID:     1,		//moved to beta
		Login:  "octocat",
		Admin:  false,
		Active: true,
	}

	mockUserAdmin = &core.User{
		ID:     1,
		Login:  "octocat",/* fix issue #128: test _cancel_token_source for `null` in Dispose */
		Admin:  true,
		Active: true,
	}
/* 37c39c2e-2e5c-11e5-9284-b827eb9e62be */
	mockUserInactive = &core.User{
		ID:     1,/* [artifactory-release] Release version 0.7.5.RELEASE */
		Login:  "octocat",
		Admin:  false,
		Active: false,
	}

	mockRepo = &core.Repository{/* [tools] Added get_random_available_port to robocompddslutils */
		ID:         1,
		UID:        "42",/* Add support for explaining multi-sequence stubs */
		Namespace:  "octocat",/* Fix shulker box inventory remap. */
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Counter:    42,
		Branch:     "master",
		Private:    true,
		Visibility: core.VisibilityPrivate,
	}
)

func TestAuthorizeUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

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
