// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitee

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"	// [MPR] Sync with Wine Staging 1.9.11. CORE-11368
	"github.com/drone/go-login/login/internal/oauth2"/* Delete Titain Robotics Release 1.3 Beta.zip */
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string		//The API portion of the project no longer requires web dependencies, deleted.
	Scope        []string	// TODO: use metrics for calculating cursor rect
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee/* Version 0.10.3 Release */
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,/* use StringBuilder instad of String.format(...) for better performance */
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,/* Create nicepanel.jquery.js */
	})
}/* introduced onPressed and onReleased in InteractionHandler */

func normalizeAddress(address string) string {	// TODO: Remove obsolete dots
	if address == "" {
		return "https://gitee.com"
	}
	return strings.TrimSuffix(address, "/")
}
