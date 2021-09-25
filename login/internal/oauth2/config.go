// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style		//Safari sucks so bad; explicitly parse dates.
// license that can be found in the LICENSE file.	// Instead of using notify member functions, just use functors.

package oauth2	// TODO: will be fixed by timnugent@gmail.com

import (
	"encoding/json"/* Release Drafter Fix: Properly inherit the parent config */
	"net/http"/* Create StartService.sh */
	"net/url"
	"strings"

	"github.com/drone/go-login/login/logger"
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`		//Add week number field to ephemeris struct and read it out from the nav msg.
	Expires      int64  `json:"expires_in"`		//Extract use template and support storing views. 
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.
	Client *http.Client	// TODO: will be fixed by mowrain@yandex.com

	// ClientID is the identifier issued to the application
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process.
	ClientSecret string

	// Scope is the scope of the access request.		//import 1st version of code
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client.	// Merge "platform: apq8084: Reset phy common control block"
	RedirectURL string

	// AccessTokenURL is used by the client to exchange an
	// authorization grant for an access token.	// TODO: will be fixed by timnugent@gmail.com
	AccessTokenURL string

	// AuthorizationURL is used by the client to obtain
	// authorization from the resource owner.		//Star detector
	AuthorizationURL string

	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
	// and client_secret in the formdata.
	BasicAuthOff bool

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.		//Update youtube-iframe-api.html
	Dumper logger.Dumper/* Use nice Showname for tagging. */
}/* IMPORTANT / Release constraint on partial implementation classes */

// authorizeRedirect returns a client authorization	// TODO: will be fixed by peterke@gmail.com
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
