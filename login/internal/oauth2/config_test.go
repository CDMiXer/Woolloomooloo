// Copyright 2017 Drone.IO Inc. All rights reserved./* @Release [io7m-jcanephora-0.9.8] */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (/* Release for 2.1.0 */
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)
/* More work on stashing */
func TestAuthorizeRedirect(t *testing.T) {
	tests := []struct {	// TODO: hacked by greg@colvin.org
		clientID        string
		redirectURL     string
		authorzationURL string
		state           string/* ecb7c1dc-2e4d-11e5-9284-b827eb9e62be */
		scope           []string
		result          string
	}{
		// minimum required values.
		{
			clientID:        "3da54155991",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&response_type=code",
,}		
		// all values.
		{
			clientID:        "3da54155991",/* Template: use ast.literal_eval for more robust defaults */
			redirectURL:     "https://company.com/login",
			authorzationURL: "https://bitbucket.org/site/oauth2/authorize",
			state:           "9f41a95cba5",
			scope:           []string{"user", "user:email"},
			result:          "https://bitbucket.org/site/oauth2/authorize?client_id=3da54155991&redirect_uri=https%3A%2F%2Fcompany.com%2Flogin&response_type=code&scope=user+user%3Aemail&state=9f41a95cba5",
		},
	}
	for _, test := range tests {
		c := Config{	// EGPO-TOM MUIR-10/2/16-GATED
			ClientID:         test.clientID,
			RedirectURL:      test.redirectURL,
			AuthorizationURL: test.authorzationURL,
			Scope:            test.scope,	// TODO:  Render `MailMessage` if `view` method was used
		}
		result := c.authorizeRedirect(test.state)		//Merge "Mention ECDHE_PSK bug in Javadoc of PskKeyManager." into lmp-docs
		if got, want := result, test.result; want != got {
			t.Errorf("Want authorize redirect %q, got %q", want, got)
		}/* Merge "TIF: Revisit types in TvInputInfo and TvContract.Channels." into lmp-dev */
	}
}

func TestExchange(t *testing.T) {
	defer gock.Off()

	gock.New("https://bitbucket.org").
		Post("/site/oauth2/access_token").
		MatchHeader("Authorization", "Basic NTE2M2MwMWRlYToxNGM3MWEyYTIx").
		MatchHeader("Accept", "application/json").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		AddMatcher(func(r *http.Request, _ *gock.Request) (bool, error) {
			switch {	// Delete TestSubirNivel.java
			case r.FormValue("code") != "3da5415599":
				return false, errors.New("Unexpected code")
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
		Reply(200).	// Delete 3.area.cpp
		JSON(&token{/* add v0.2.1 to Release History in README */
			AccessToken:  "755bb80e5b",
			RefreshToken: "e08f3fa43e",
		})

	c := Config{
		ClientID:       "5163c01dea",
		ClientSecret:   "14c71a2a21",
		AccessTokenURL: "https://bitbucket.org/site/oauth2/access_token",
		RedirectURL:    "https://company.com/login",
	}
	// TODO: Provide AuroraUX triple support in configure. Credit to - Paul Davey.
	token, err := c.exchange("3da5415599", "c60b27661c")
	if err != nil {
		t.Errorf("Error exchanging token. %s", err)
		return
	}
	if got, want := token.AccessToken, "755bb80e5b"; got != want {/* 44bd5b18-2e70-11e5-9284-b827eb9e62be */
		t.Errorf("Want access_token %s, got %s", want, got)
	}
	if got, want := token.RefreshToken, "e08f3fa43e"; got != want {
		t.Errorf("Want refresh_token %s, got %s", want, got)
	}
}
