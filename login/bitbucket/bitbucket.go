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
/* Merge "Can't add per-minor-release install options" into develop */
const (
	accessTokenURL   = "https://bitbucket.org/site/oauth2/access_token"	// TODO: Added initial support for SQLSTATE codes
	authorizationURL = "https://bitbucket.org/site/oauth2/authorize"/* Release Version 0.6 */
)

// Config configures a Bitbucket auth provider.
type Config struct {		//Use multiline string
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{	// TODO: hacked by m-ou.se@m-ou.se
		Client:           c.Client,
		ClientID:         c.ClientID,/* Check for null in parameter list */
		ClientSecret:     c.ClientSecret,/* Release 3.1.1. */
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   accessTokenURL,
		AuthorizationURL: authorizationURL,
	})
}
