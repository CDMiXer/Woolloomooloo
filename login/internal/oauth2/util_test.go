// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Merge "Publish MediaMetadataRetriever.java as public API" into honeycomb */

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
	}/* Release for 23.4.1 */
}

func Test_validateState(t *testing.T) {
	tests := []struct {	// TODO: will be fixed by earlephilhower@yahoo.com
		state string
		value string/* Released version 1.9.14 */
		err   error
	}{	// Implemented debug command and DEBUG config key
		{
			state: "4d65822107fcfd52",
			value: "4d65822107fcfd52",
		},
		{
			state: "4d65822107fcfd52",
			value: "0000000000000000",/* Bugs fixed; Release 1.3rc2 */
			err:   ErrState,	// TODO: will be fixed by nick@perfectabstractions.com
		},
		{
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)
		if test.value != "" {
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}	// TODO: hacked by davidad@alum.mit.edu
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}	// TODO: Fix bad link to Auto-Factory.
	}		//721d3464-2e40-11e5-9284-b827eb9e62be
}

func Test_deleteState(t *testing.T) {		//copied 2.0.0-beta-4
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}		//readme guide improved
