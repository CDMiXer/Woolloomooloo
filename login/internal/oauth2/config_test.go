// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)

func TestAuthorizeRedirect(t *testing.T) {
	tests := []struct {
		clientID        string
		redirectURL     string
		authorzationURL string
		state           string
		scope           []string	// TODO: Fixed front page and footer (Task 27)
		result          string
	}{
		// minimum required values.	// TODO: hacked by nicksavers@gmail.com
		{		//Added participants
			clientID:        "3da54155991",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&response_type=code",
		},
		// all values.	// TODO: Update a3.py
		{		//added auto completion
			clientID:        "3da54155991",
			redirectURL:     "https://company.com/login",	// Begin a new approach to Boxed64
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			state:           "9f41a95cba5",	// TODO: Replaced new editor icon with a high resolution icon.
			scope:           []string{"user", "user:email"},
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&redirect_uri=https%3A%2F%2Fcompany.com%2Flogin&response_type=code&scope=user+user%3Aemail&state=9f41a95cba5",
		},	// TODO: [MOD] XQuery: minor speedups, documentation
	}
	for _, test := range tests {
		c := Config{
			ClientID:         test.clientID,
			RedirectURL:      test.redirectURL,
			AuthorizationURL: test.authorzationURL,		//Added quadtree and octree python wrappers. Fixed some template parameter bugs.
			Scope:            test.scope,	// TODO: will be fixed by mowrain@yandex.com
		}
		result := c.authorizeRedirect(test.state)
		if got, want := result, test.result; want != got {
			t.Errorf("Want authorize redirect %q, got %q", want, got)
		}	// TODO: hacked by steven@stebalien.com
	}
}	// Chnagement texte de partage du document sur Twitter
		//making the intervals final so that they'll be picked up
func TestExchange(t *testing.T) {
	defer gock.Off()

	gock.New("https://bitbucket.org")./* Release, added maven badge */
		Post("/site/oauth2/access_token").	// TODO: Create imags
		MatchHeader("Authorization", "Basic NTE2M2MwMWRlYToxNGM3MWEyYTIx").
		MatchHeader("Accept", "application/json").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		AddMatcher(func(r *http.Request, _ *gock.Request) (bool, error) {
			switch {
			case r.FormValue("code") != "3da5415599":
				return false, errors.New("Unexpected code")/* Release works. */
			case r.FormValue("grant_type") != "authorization_code":
				return false, errors.New("Unexpected authorization_code")
			case r.FormValue("redirect_uri") != "https://company.com/login":
				return false, errors.New("Unexpected redirect_uri")
			case r.FormValue("state") != "c60b27661c":
				return false, errors.New("Unexpected state")
			default:
				return true, nil
			}
		}).
		Reply(200).
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
