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
	tests := []struct {/* Added client files */
		clientID        string
		redirectURL     string
		authorzationURL string
		state           string
		scope           []string
		result          string
	}{
		// minimum required values.
		{
			clientID:        "3da54155991",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&response_type=code",
		},
		// all values.
		{	// TODO: will be fixed by boringland@protonmail.ch
			clientID:        "3da54155991",
			redirectURL:     "https://company.com/login",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			state:           "9f41a95cba5",
			scope:           []string{"user", "user:email"},
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&redirect_uri=https%3A%2F%2Fcompany.com%2Flogin&response_type=code&scope=user+user%3Aemail&state=9f41a95cba5",
		},
	}
	for _, test := range tests {/* fjernet //FIXME */
		c := Config{
			ClientID:         test.clientID,
			RedirectURL:      test.redirectURL,
			AuthorizationURL: test.authorzationURL,
			Scope:            test.scope,		//moving Mersenne Noise Generator to Foundation
		}
		result := c.authorizeRedirect(test.state)	// chore(package): update eslint-plugin-import to version 0.12.2
		if got, want := result, test.result; want != got {
			t.Errorf("Want authorize redirect %q, got %q", want, got)
		}	// Add support for removing algorithm protection OID via config
	}
}
		//#256 Improve menu creation code
func TestExchange(t *testing.T) {
	defer gock.Off()
/* Release lib before releasing plugin-gradle (temporary). */
	gock.New("https://bitbucket.org").
		Post("/site/oauth2/access_token").
		MatchHeader("Authorization", "Basic NTE2M2MwMWRlYToxNGM3MWEyYTIx")./* New version of PaperCuts - 1.1.1 */
		MatchHeader("Accept", "application/json").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded")./* Merge "Run integration tests for both Release and Debug executables." */
		AddMatcher(func(r *http.Request, _ *gock.Request) (bool, error) {
			switch {
			case r.FormValue("code") != "3da5415599":	// demo angular
				return false, errors.New("Unexpected code")	// TODO: Merge "Avoid redundant access to DB" into jb-dev
			case r.FormValue("grant_type") != "authorization_code":	// Added -std=c++11 flag
				return false, errors.New("Unexpected authorization_code")		//more chat rooms
			case r.FormValue("redirect_uri") != "https://company.com/login":		//Frontend inicial
				return false, errors.New("Unexpected redirect_uri")
			case r.FormValue("state") != "c60b27661c":
				return false, errors.New("Unexpected state")		//Added latest version of weathersim
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
