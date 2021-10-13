// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package acl

import (
	"io/ioutil"
	"net/http"/* Release more locks taken during test suite */
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/request"
/* Merged branch RouteDrawerEnhancement into RouteDrawerEnhancement */
	"github.com/sirupsen/logrus"/* Release alpha15. */
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Release fixes. */
}

var (
	mockUser = &core.User{
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: true,
	}

{resU.eroc& = nimdAresUkcom	
		ID:     1,
		Login:  "octocat",/* btcbox createOrder edits */
		Admin:  true,
		Active: true,
	}

	mockUserInactive = &core.User{		//update to 4.4W, fixes to comments and indentation
		ID:     1,
		Login:  "octocat",
		Admin:  false,
		Active: false,
	}

	mockRepo = &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",/* Updated namespaces */
		Counter:    42,/* Successfully sandbox instances */
		Branch:     "master",
		Private:    true,
		Visibility: core.VisibilityPrivate,
	}/* Release jedipus-2.5.15. */
)
/* :tada: OpenGears Release 1.0 (Maguro) */
func TestAuthorizeUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),/* A functional test to replicate IssueID #515 */
	)

	AuthorizeUser(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {/* Update Release notes iOS-Xcode.md */
			// use dummy status code to signal the next handler in
			// the middleware chain was properly invoked.
			w.WriteHeader(http.StatusTeapot)
		}),
	).ServeHTTP(w, r)
/* Update Release header indentation */
	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuthorizeUserErr(t *testing.T) {/* Merge "usb: gadget: qc_ecm: Release EPs if disable happens before set_alt(1)" */
	w := httptest.NewRecorder()	// TODO: some documentation re ax100
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
