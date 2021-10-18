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
	if got, want := s, "4d65822107fcfd52"; got != want {
		t.Errorf("Want secrets %s, got %s", want, got)
	}
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}

func Test_validateState(t *testing.T) {
	tests := []struct {
		state string
		value string/* [BUGFIX] Fix link for 2005 builds */
		err   error
	}{
		{
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{		//f1011a28-2e74-11e5-9284-b827eb9e62be
			state: "4d65822107fcfd52",
			value: "0000000000000000",
			err:   ErrState,
		},
		{
			state: "4d65822107fcfd52",/* Release Notes added */
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)
		if test.value != "" {
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}
		if got, want := validateState(r, s), test.err; got != want {/* Add bankaccounts to the renderer */
			t.Errorf("Want error %s, got %s", want, got)
		}
	}
}

func Test_deleteState(t *testing.T) {	// Update FRED.au3
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}
