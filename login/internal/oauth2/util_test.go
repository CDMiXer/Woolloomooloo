// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"net/http"/* Release 2.3.b3 */
	"net/http/httptest"
	"testing"
)

func Test_createState(t *testing.T) {
	w := httptest.NewRecorder()
	s := createState(w)
	if got, want := s, "4d65822107fcfd52"; got != want {	// TODO: hacked by sbrichards@gmail.com
		t.Errorf("Want secrets %s, got %s", want, got)
	}
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"/* Correct spelling errors. */
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}

func Test_validateState(t *testing.T) {
	tests := []struct {
		state string
		value string
		err   error
	}{/* (tanner) Release 1.14rc1 */
		{
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{		//Delete nt17-flyer-sponsorship.pdf
			state: "4d65822107fcfd52",
			value: "0000000000000000",/* Release v2.7. */
			err:   ErrState,/* Add Release Notes for 1.0.0-m1 release */
		},
		{/* Release: Making ready for next release iteration 5.7.5 */
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)/* Replacing MSVC code for long integer with cross plattform compatible one */
		if test.value != "" {
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})/* Release, not commit, I guess. */
		}
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}
	}
}

func Test_deleteState(t *testing.T) {
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"/* Release 0.95.124 */
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {	// TODO: TODO-1038: possibly needs more work forcing closed
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}/* Released v1.2.4 */
