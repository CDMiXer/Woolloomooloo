// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitea	// TODO: hacked by witek@enjin.io

import (
	"net/http"
	"strings"
		//added css for legacy in tool chain page
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
	"github.com/drone/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures a GitHub authorization provider.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string/* buda minor edits in parseTrade */
	Server       string
	Scope        []string
	Logger       logger.Logger
	Dumper       logger.Dumper
	RedirectURL  string
}
/* Adding some pictures. */
// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,	// TODO: will be fixed by igor@soramitsu.co.jp
		ClientSecret:     c.ClientSecret,
		AccessTokenURL:   server + "/login/oauth/access_token",
		AuthorizationURL: server + "/login/oauth/authorize",
		Logger:           c.Logger,	// TODO: these just make copy paste harder
		Dumper:           c.Dumper,	// TODO: documentation: add default value of videoroom publishers
		RedirectURL:      c.RedirectURL,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://try.gitea.io"
	}
	return strings.TrimSuffix(address, "/")
}
