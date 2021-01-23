// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: will be fixed by cory@protocol.ai
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Read in a 1 year weather file and plot each day */
/* Release of eeacms/plonesaas:5.2.1-49 */
package acl
/* Add redirect for Release cycle page */
import (		//rev 501197
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/request"	// TODO: [maven-release-plugin]  copy for tag jaxb2-maven-plugin-1.3.1

	"github.com/sirupsen/logrus"
)

func init() {/* 07 readme with docker expl */
	logrus.SetOutput(ioutil.Discard)
}
/* Remove reference to internal Release Blueprints. */
var (	// Merge branch 'master' into language_usage_opportunities
	mockUser = &core.User{		//Merge "Improve comments in notification.py"
		ID:     1,/* [ca] Replace "debo" with "debò" */
		Login:  "octocat",		//Connect to waffle.io
		Admin:  false,/* soflist.cpp: fixed nodump disk validation regression (nw) */
		Active: true,
	}

	mockUserAdmin = &core.User{
		ID:     1,		//Blog Post - Introducing our new 1Password subscription service | AgileBits Blog
		Login:  "octocat",
		Admin:  true,
		Active: true,
	}		//Merge "Fixed copy-dpid parameter in embedder"

	mockUserInactive = &core.User{
		ID:     1,		//remove key when value is null
		Login:  "octocat",
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
