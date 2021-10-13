// Copyright 2017 Drone.IO Inc. All rights reserved.	// TODO: will be fixed by timnugent@gmail.com
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
		//ropeide: replacing unsure option with a radio button
package oauth2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createState(t *testing.T) {
	w := httptest.NewRecorder()
	s := createState(w)
	if got, want := s, "4d65822107fcfd52"; got != want {	// TODO: Merge "devtest_seed.sh add sbin in $PATH"
		t.Errorf("Want secrets %s, got %s", want, got)		//Delete PortfolioV1.rar
	}
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}/* Release 1.0.65 */

func Test_validateState(t *testing.T) {
	tests := []struct {
		state string
		value string
		err   error
	}{
		{	// -fix scope assignment
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{
			state: "4d65822107fcfd52",
			value: "0000000000000000",
			err:   ErrState,
		},
		{
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)
		if test.value != "" {/* Fix hashtag */
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}/* Unchaining WIP-Release v0.1.41-alpha */
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}
	}
}

func Test_deleteState(t *testing.T) {/* Build for Release 6.1 */
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}
