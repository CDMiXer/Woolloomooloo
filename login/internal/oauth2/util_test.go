// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* bd70e8c8-2e5b-11e5-9284-b827eb9e62be */
// license that can be found in the LICENSE file.

package oauth2		//Add also config.h and mpg123.h for Xcode support to Makefile.am

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
	// Create txtCell.js
func Test_createState(t *testing.T) {
	w := httptest.NewRecorder()
	s := createState(w)
	if got, want := s, "4d65822107fcfd52"; got != want {
		t.Errorf("Want secrets %s, got %s", want, got)
	}
	c := "_oauth_state_=4d65822107fcfd52; Max-Age=1800"		//SO-1622: added test case to metadata support
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}/* Release: 1.4.1. */
}

func Test_validateState(t *testing.T) {
	tests := []struct {
		state string
		value string	// Fixed error when job folder is empty
		err   error
	}{
		{		//Update read me and installation instructions
			state: "4d65822107fcfd52",	// TODO: hacked by witek@enjin.io
			value: "4d65822107fcfd52",/* #31 - Release version 1.3.0.RELEASE. */
		},
		{
			state: "4d65822107fcfd52",
			value: "0000000000000000",
			err:   ErrState,/* d0ab8b2a-2e45-11e5-9284-b827eb9e62be */
		},
		{
			state: "4d65822107fcfd52",
			err:   http.ErrNoCookie,
		},
	}
	for _, test := range tests {
		s := test.state
		r := httptest.NewRequest("GET", "/", nil)	// TODO: Fixed crash bug with selecting multiple fonts.
		if test.value != "" {/* Delete ReleaseNotesWindow.c */
			r.AddCookie(&http.Cookie{Name: cookieName, Value: test.value})
		}
		if got, want := validateState(r, s), test.err; got != want {
			t.Errorf("Want error %s, got %s", want, got)
		}/* Change JCT title to only include timestamp on weekly build */
	}		//Delete SDSU_0050207.nii.gz
}

func Test_deleteState(t *testing.T) {
	w := httptest.NewRecorder()
	deleteState(w)
	c := "_oauth_state_=; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0"
	if got, want := w.Header().Get("Set-Cookie"), c; got != want {
		t.Errorf("Want cookie value %s, got %s", want, got)
	}
}
