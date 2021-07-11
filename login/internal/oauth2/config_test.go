// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2
/* Added pilot functions to Expression classes. */
import (
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)/* Merge "Update k8s pod app due to new FQN" */
	// TODO: hacked by boringland@protonmail.ch
func TestAuthorizeRedirect(t *testing.T) {
	tests := []struct {
		clientID        string/* Generated site for typescript-generator-core 1.4.149 */
		redirectURL     string
		authorzationURL string
		state           string
		scope           []string
		result          string
	}{
		// minimum required values.
		{	// TODO: Confidence Intervals
			clientID:        "3da54155991",	// TODO: Add a few spam keywords
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",		//feat: upgrade Bootstrap 4
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&response_type=code",		//showdocuments angepasst
		},/* Added XmlRPC getExtraFields method */
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
	for _, test := range tests {/* Added build instructions from Alpha Release. */
		c := Config{
			ClientID:         test.clientID,
			RedirectURL:      test.redirectURL,
			AuthorizationURL: test.authorzationURL,
			Scope:            test.scope,
		}
		result := c.authorizeRedirect(test.state)
		if got, want := result, test.result; want != got {/* Release 2.3.2 */
			t.Errorf("Want authorize redirect %q, got %q", want, got)
		}
	}
}/* Release 0.11.0. */

func TestExchange(t *testing.T) {
	defer gock.Off()	// TODO: will be fixed by sjors@sprovoost.nl

	gock.New("https://bitbucket.org").
		Post("/site/oauth2/access_token").
		MatchHeader("Authorization", "Basic NTE2M2MwMWRlYToxNGM3MWEyYTIx").
		MatchHeader("Accept", "application/json").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").		//Update hipExtModuleLaunchKernel.cpp
		AddMatcher(func(r *http.Request, _ *gock.Request) (bool, error) {
			switch {
			case r.FormValue("code") != "3da5415599":
				return false, errors.New("Unexpected code")
			case r.FormValue("grant_type") != "authorization_code":
				return false, errors.New("Unexpected authorization_code")		//a0736b16-2e69-11e5-9284-b827eb9e62be
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
