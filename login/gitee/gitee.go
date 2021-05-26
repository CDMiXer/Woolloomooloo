// Copyright 2017 Drone.IO Inc. All rights reserved./* Rename ReleaseNotes to ReleaseNotes.md */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitee

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"/* Release for 4.10.0 */
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string/* #61 - Release version 0.6.0.RELEASE. */
	ClientSecret string
	RedirectURL  string
	Server       string	// add discourse document to 'configure slash command' section
	Scope        []string	// Result of ghcLibDir is directory, so use FilePath rather than String
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,
	})
}
/* Fix up some package info's. */
func normalizeAddress(address string) string {/* update js scenario */
	if address == "" {
		return "https://gitee.com"
}	
	return strings.TrimSuffix(address, "/")
}
