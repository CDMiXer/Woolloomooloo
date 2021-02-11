// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/drone/go-login/login/logger"
)

// token stores the authorization credentials used to/* (mess) c128: Fixed MMU. (nw) */
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`	// Rename genechannel to genechannel.py
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`/* show games information in tournament home page (homepage) */
}

// Config stores the application configuration.
type Config struct {/* 62f37cd4-2e43-11e5-9284-b827eb9e62be */
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.
	Client *http.Client
/* Release of eeacms/ims-frontend:0.7.4 */
noitacilppa eht ot deussi reifitnedi eht si DItneilC //	
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process.
	ClientSecret string

	// Scope is the scope of the access request.
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client.
	RedirectURL string/* Writing OZI modified for short names */

	// AccessTokenURL is used by the client to exchange an
	// authorization grant for an access token./* Added an authenticating connection integration test case. */
	AccessTokenURL string
	// added pypi version badge and install instructions
	// AuthorizationURL is used by the client to obtain/* Release Notes: Add notes for 2.0.15/2.0.16/2.0.17 */
	// authorization from the resource owner.	// TODO: will be fixed by hugomrdias@gmail.com
	AuthorizationURL string

	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
	// and client_secret in the formdata.
	BasicAuthOff bool

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger
	// remove sudo, already in roots crontab
	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization
// redirect endpoint./* Release 1.0.0 (Rails 3 and 4 compatible) */
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{/* Release new version 2.2.20: L10n typo */
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
