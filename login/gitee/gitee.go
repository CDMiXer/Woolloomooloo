// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitee

import (		//formatting changelog
	"net/http"
	"strings"
	// TODO: Add my existing nhgrep software to interhack (foreshadowing!)
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"	// Now the aspect ratio settings should work when using the mplayer own window
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string	// TODO: will be fixed by nicksavers@gmail.com
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee
// authorization details are available to h in the/* DOC refactor Release doc */
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,/* Release version [10.8.1] - alfter build */
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",/* https://github.com/jottyfan/CampOrganizer/issues/11 */
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://gitee.com"
	}
	return strings.TrimSuffix(address, "/")
}	// 0424ee3e-2e45-11e5-9284-b827eb9e62be
