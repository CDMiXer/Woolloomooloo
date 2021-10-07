// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Update Nelson_Siegel.m */

package gitee

import (
	"net/http"
	"strings"
	// TODO: Add some empty lines
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string
	ClientSecret string	// TODO: hacked by 13860583249@yeah.net
	RedirectURL  string
	Server       string
	Scope        []string
	Client       *http.Client
}	// TODO: let people know it was updated

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee
// authorization details are available to h in the
// http.Request context.	// Modified unattended upgrades
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

func normalizeAddress(address string) string {	// Delete pdo-query-builder.zip
	if address == "" {
		return "https://gitee.com"
	}
	return strings.TrimSuffix(address, "/")
}
