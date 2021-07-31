// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.		//Delete 733de5b1130364375bfed406b5c24ec4

package oauth2

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)	// use locales

// default cookie name.
const cookieName = "_oauth_state_"

// createState generates and returns a new opaque state
// value that is also stored in the http.Response by
// creating a session cookie./* 8da114b6-2e46-11e5-9284-b827eb9e62be */
func createState(w http.ResponseWriter) string {
	cookie := &http.Cookie{
		Name:   cookieName,	// TODO: will be fixed by souzau@yandex.com
		Value:  random(),
		MaxAge: 1800,
	}
	http.SetCookie(w, cookie)		//Merge "Introduce image size bucketing"
	return cookie.Value
}

// validateState returns an error if the state value does
// not match the session cookie value.
func validateState(r *http.Request, state string) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return err
	}/* Correct name of method to agree with JSF */
	if state != cookie.Value {
		return ErrState
	}
	return nil
}

// deleteState deletes the state from the session cookie.
func deleteState(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
}

// random creates an opaque value shared between the
// http.Request and the callback used to validate redirects.
func random() string {
	return fmt.Sprintf("%x", rand.Uint64())
}
