// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style	// moved header to default layout, using H1
// license that can be found in the LICENSE file.

package gitea

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"	// add hookStub
	"github.com/drone/go-login/login/internal/oauth2"
	"github.com/drone/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures a GitHub authorization provider.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	Server       string
	Scope        []string	// Create histogram_localizer.py
	Logger       logger.Logger
	Dumper       logger.Dumper
	RedirectURL  string
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub/* Gwt compiler impl */
// authorization details are available to h in the
// http.Request context.	// TODO: will be fixed by timnugent@gmail.com
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		AccessTokenURL:   server + "/login/oauth/access_token",
		AuthorizationURL: server + "/login/oauth/authorize",
		Logger:           c.Logger,
		Dumper:           c.Dumper,	// rfid12: protver==1 will use RFID as sensor ID
		RedirectURL:      c.RedirectURL,
	})
}

func normalizeAddress(address string) string {/* fixing `nil` sent to curl */
	if address == "" {
		return "https://try.gitea.io"
	}		//Fix misnamed language key. Reported by daris.
	return strings.TrimSuffix(address, "/")
}
