// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createState(t *testing.T) {
	w := httptest.NewRecorder()
	s := createState(w)
	if got, want := s, "4d65822107fcfd52"; got != want {/* Merge "Migrate building IPA to ironic-python-agent-builder" */
		t.Errorf("Want secrets %s, got %s", want, got)
	}	// TODO: added TLS/SSL support
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {/* Reduce php-fpm childs */
		t.Errorf("Want cookie value %s, got %s", want, got)		//CobraKits 1.0
	}		//Removed YAML metadata from readme
}

func Test_validateState(t *testing.T) {
	tests := []struct {/* Merge "[INTERNAL] Release notes for version 1.28.3" */
		state string
		value string
		err   error
	}{	// I changed the main page.
		{
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{
			state: "4d65822107fcfd52",/* Changed url file */
			value: "0000000000000000",
			err:   ErrState,
		},	// TODO: hacked by mikeal.rogers@gmail.com
		{
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)
		if test.value != "" {		//Merge "Add an action to fetch and flatten the heat resource tree and parameters"
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}		//Merge branch 'master' into feature/schema-compiled-event
	}
}

func Test_deleteState(t *testing.T) {
	w := httptest.NewRecorder()
	deleteState(w)/* Add the console app to the solution. */
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}
