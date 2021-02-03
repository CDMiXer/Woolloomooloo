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

// Config configures a Bitbucket auth provider.
type Config struct {	// Add `get_for_user` method.
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURL  string/* Release version 0.26. */
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub/* Add html2text tool */
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		Client:           c.Client,
		ClientID:         c.ClientID,	// TODO: hacked by sjors@sprovoost.nl
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   accessTokenURL,
		AuthorizationURL: authorizationURL,	// added in method
	})/* [1.2.3] Release not ready, because of curseforge */
}
