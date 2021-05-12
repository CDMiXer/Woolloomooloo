// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

var _ login.Middleware = (*Config)(nil)

const (
	accessTokenURL   = "https://bitbucket.org/site/oauth2/access_token"
	authorizationURL = "https://bitbucket.org/site/oauth2/authorize"
)
/* Inital Release */
// Config configures a Bitbucket auth provider.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// Handler returns a http.Handler that runs h at the/* Release 0.15.2 */
// completion of the GitHub authorization flow. The GitHub	// TODO: will be fixed by cory@protocol.ai
// authorization details are available to h in the
// http.Request context.	// Correct error when the email_text isn't filled
func (c *Config) Handler(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   accessTokenURL,
		AuthorizationURL: authorizationURL,		//Refactored project creation code into something tolerable.
	})
}
