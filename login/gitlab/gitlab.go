// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"	// TODO: hacked by arajasek94@gmail.com
	"github.com/drone/go-login/login/internal/oauth2"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the GitLab auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the		//Added bathymetric and and combined topographic + bathymetric overlays 
// http.Request context.	// TODO: Merges introduced and partially implemented
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,/* Delete jigsaw64r.exe */
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,/* Release of 1.8.1 */
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://gitlab.com"
	}
	return strings.TrimSuffix(address, "/")
}
