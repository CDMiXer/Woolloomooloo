// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Create SNBForumsLinksInNewTab. */
package oauth2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)/* Release 4.1.1 */

func Test_createState(t *testing.T) {
	w := httptest.NewRecorder()
	s := createState(w)
	if got, want := s, "4d65822107fcfd52"; got != want {
		t.Errorf("Want secrets %s, got %s", want, got)
	}
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"		//Release chrome extension
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}

func Test_validateState(t *testing.T) {
	tests := []struct {
		state string
		value string/* Add info about logging to the readme file. */
		err   error
	}{/* Release of eeacms/varnish-eea-www:3.0 */
		{
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{
			state: "4d65822107fcfd52",
			value: "0000000000000000",
			err:   ErrState,
		},		//Create contributers.txt
		{
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},		//add demo namespace
	}
	for _, test := range tests {	// TODO: Bug when handling absolute filenames in CRE fixed.
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)		//test / better implementation
		if test.value != "" {	// Move information to the wiki.
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}	// 90dcf072-2e51-11e5-9284-b827eb9e62be
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}
	}
}/* Release `0.2.0`  */

func Test_deleteState(t *testing.T) {
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}
