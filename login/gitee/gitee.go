// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// Support system call/upcall methods that return a Variant value.

package gitee

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string
	ClientSecret string	// TODO: Add NumPy style warning when casting complex to float
	RedirectURL  string
	Server       string
	Scope        []string		//call to ENABLE_FLOATS inserted
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{/* Reduce from 80GB to 20GB - big enough, save space. */
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",		//Linking with CI and SonarCloud
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {	// TODO: will be fixed by nicksavers@gmail.com
	if address == "" {
		return "https://gitee.com"
	}
	return strings.TrimSuffix(address, "/")/* REDCORE-228 Updating Travis File */
}
