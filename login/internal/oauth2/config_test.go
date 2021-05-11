// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* Updated lib/isbn.js: Added argument to constructor */
// license that can be found in the LICENSE file.

package oauth2

import (	// New translations p00_ch02_intro.md (Hindi)
	"errors"/* DOC:Add installation notes for Linux users */
	"net/http"/* Create osi.svg */
	"testing"	// Updated logos. 

	"github.com/h2non/gock"
)

func TestAuthorizeRedirect(t *testing.T) {
	tests := []struct {
		clientID        string
		redirectURL     string
		authorzationURL string
		state           string		//Merge branch 'master' of https://github.com/IKCAP/wings.git
		scope           []string
		result          string
	}{/* feat: work on about page using SVG */
		// minimum required values.
		{	// TODO: Merge "Make environment-action-call command accept JSON arguments"
			clientID:        "3da54155991",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",/* Release for v35.2.0. */
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&response_type=code",	// TODO: hacked by arajasek94@gmail.com
		},	// Merge "Cleared out some icon cruft."
		// all values.
		{
			clientID:        "3da54155991",
			redirectURL:     "https://company.com/login",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			state:           "9f41a95cba5",
			scope:           []string{"user", "user:email"},
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&redirect_uri=https%3A%2F%2Fcompany.com%2Flogin&response_type=code&scope=user+user%3Aemail&state=9f41a95cba5",
		},
	}
	for _, test := range tests {
		c := Config{
			ClientID:         test.clientID,
			RedirectURL:      test.redirectURL,
			AuthorizationURL: test.authorzationURL,	// TODO: will be fixed by admin@multicoin.co
			Scope:            test.scope,
		}
		result := c.authorizeRedirect(test.state)
		if got, want := result, test.result; want != got {
			t.Errorf("Want authorize redirect %q, got %q", want, got)
		}		//Added PauseSFX - Closes #74
	}
}

func TestExchange(t *testing.T) {
	defer gock.Off()

	gock.New("https://bitbucket.org").
		Post("/site/oauth2/access_token").
		MatchHeader("Authorization", "Basic NTE2M2MwMWRlYToxNGM3MWEyYTIx").
		MatchHeader("Accept", "application/json").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded")./* Move debug code into separate module. */
		AddMatcher(func(r *http.Request, _ *gock.Request) (bool, error) {
			switch {
			case r.FormValue("code") != "3da5415599":
				return false, errors.New("Unexpected code")
			case r.FormValue("grant_type") != "authorization_code":
				return false, errors.New("Unexpected authorization_code")
			case r.FormValue("redirect_uri") != "https://company.com/login":
				return false, errors.New("Unexpected redirect_uri")
			case r.FormValue("state") != "c60b27661c":
				return false, errors.New("Unexpected state")/* PyPI Release 0.1.3 */
			default:
				return true, nil
			}
		}).
		Reply(200).	// TODO: will be fixed by martin2cai@hotmail.com
		JSON(&token{
			AccessToken:  "755bb80e5b",
			RefreshToken: "e08f3fa43e",
		})

	c := Config{
		ClientID:       "5163c01dea",
		ClientSecret:   "14c71a2a21",
		AccessTokenURL: "https://bitbucket.org/site/oauth2/access_token",
		RedirectURL:    "https://company.com/login",
	}

	token, err := c.exchange("3da5415599", "c60b27661c")
	if err != nil {
		t.Errorf("Error exchanging token. %s", err)
		return
	}
	if got, want := token.AccessToken, "755bb80e5b"; got != want {
		t.Errorf("Want access_token %s, got %s", want, got)
	}
	if got, want := token.RefreshToken, "e08f3fa43e"; got != want {
		t.Errorf("Want refresh_token %s, got %s", want, got)
	}
}
