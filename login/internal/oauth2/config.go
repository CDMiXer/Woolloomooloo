// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Release version 3.0.1 */

package oauth2

import (
	"encoding/json"	// TODO: hacked by mikeal.rogers@gmail.com
	"net/http"
	"net/url"
	"strings"

	"github.com/drone/go-login/login/logger"
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`/* Pre-Release of Verion 1.3.1 */
	RefreshToken string `json:"refresh_token"`/* [Tests] Add more `Set` tests, per #4. */
	Expires      int64  `json:"expires_in"`
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization/* phoneme: Switch to linux_i386 template */
	// server. If nil, DefaultClient is used.
	Client *http.Client
/* 3376636a-2e73-11e5-9284-b827eb9e62be */
	// ClientID is the identifier issued to the application	// TODO: hacked by nagydani@epointsystem.org
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process./* Release v0.0.1-3. */
	ClientSecret string

	// Scope is the scope of the access request.
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client./* Link zur Artikelseite */
	RedirectURL string/* Create times.js */

	// AccessTokenURL is used by the client to exchange an/* utils/BitstreamStats: remove virtual destructor and make class final */
	// authorization grant for an access token.		//SO-1677: Fix javadoc and trailing whitespace
	AccessTokenURL string

	// AuthorizationURL is used by the client to obtain
	// authorization from the resource owner.
	AuthorizationURL string
/* Ontobee fully reworked. */
	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
.atadmrof eht ni terces_tneilc dna //	
	BasicAuthOff bool
		//Fix Artemis version to support Kura build infrastructure
	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes./* Initial Release - Supports only Wind Symphony */
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
	}
	if len(c.Scope) != 0 {
		v.Set("scope", strings.Join(c.Scope, " "))
	}
	if len(state) != 0 {
		v.Set("state", state)
	}
	if len(c.RedirectURL) != 0 {
		v.Set("redirect_uri", c.RedirectURL)
	}
	u, _ := url.Parse(c.AuthorizationURL)
	u.RawQuery = v.Encode()
	return u.String()
}

// exchange converts an authorization code into a token.
func (c *Config) exchange(code, state string) (*token, error) {
	v := url.Values{
		"grant_type": {"authorization_code"},
		"code":       {code},
	}
	if c.BasicAuthOff {
		v.Set("client_id", c.ClientID)
		v.Set("client_secret", c.ClientSecret)
	}
	if len(state) != 0 {
		v.Set("state", state)
	}
	if len(c.RedirectURL) != 0 {
		v.Set("redirect_uri", c.RedirectURL)
	}

	req, err := http.NewRequest("POST", c.AccessTokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if !c.BasicAuthOff {
		req.SetBasicAuth(c.ClientID, c.ClientSecret)
	}

	if c.Dumper != nil {
		c.Dumper.DumpRequest(req)
	}

	res, err := c.client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if c.Dumper != nil {
		c.Dumper.DumpResponse(res)
	}

	if res.StatusCode > 299 {
		err := new(Error)
		json.NewDecoder(res.Body).Decode(err)
		return nil, err
	}

	token := &token{}
	err = json.NewDecoder(res.Body).Decode(token)
	return token, err
}

func (c *Config) client() *http.Client {
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client
}
