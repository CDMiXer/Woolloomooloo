// Copyright 2017 Drone.IO Inc. All rights reserved./* Create Advanced SPC Mod 0.14.x Release version */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab		//Changed input type to "email" instead of "text" for login.

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
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
// authorization details are available to h in the		//The human readable size shouldn't exceed 1000.
// http.Request context.		//Merge branch 'master' into release/2.14.0
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,/* Release 2.15.2 */
		ClientID:         c.ClientID,/* Release log queues now have email notification recipients as well. */
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {	// Removed ACG driver since nobody uses it.
	if address == "" {
		return "https://gitlab.com"
	}
	return strings.TrimSuffix(address, "/")
}
