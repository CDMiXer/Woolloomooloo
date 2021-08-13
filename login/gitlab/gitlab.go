// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab/* Released URB v0.1.4 */

import (
	"net/http"/* fix dumb fender moment */
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)
	// TODO: hacked by why@ipfs.io
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
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {/* Release of eeacms/www-devel:19.8.15 */
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,/* Added Team preview for Sandbox */
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",	// Uri parameters hierarchy
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://gitlab.com"
	}
	return strings.TrimSuffix(address, "/")
}
