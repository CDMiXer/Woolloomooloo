// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* welcome page styles */

package acl		//Delete t1a03 css AlexPark.html

import (	// TODO: indentation sans modif de code
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"	// Update JSON example to reflect newer JSON format.

	"github.com/drone/drone/core"/* SVMI-TOM MUIR-1/20/17-redone by Adam Callow */
	"github.com/drone/drone/handler/api/request"
/* add database setup */
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Release notes 8.0.3 */
}/* Release for 2.4.0 */

var (
	mockUser = &core.User{
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: true,
	}
		//Issue #3: channel icons.
	mockUserAdmin = &core.User{	// Only override locus common name with set common name if defined.
		ID:     1,/* Merge "wlan: Release 3.2.3.114" */
		Login:  "octocat",
		Admin:  true,
		Active: true,
	}

	mockUserInactive = &core.User{/* Create Largest_element.cpp */
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: false,
	}
		//Recordings can now be sorted
	mockRepo = &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "octocat",	// TODO: hacked by willem.melching@gmail.com
		Name:       "hello-world",
		Slug:       "octocat/hello-world",/* 3.3.1 Release */
		Counter:    42,
		Branch:     "master",
		Private:    true,	// Merge "Explicitly list the valid transitions to RESUMING state"
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
