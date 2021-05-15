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
	// TODO: Add sphinx auto-generated API docs
// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`/* (v2) Scene cannvas: select the object created with a drop. */
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}
	// TODO: opensubtitles-api 2.x
// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used./* SQF - Adjusting "delete" function name */
	Client *http.Client

	// ClientID is the identifier issued to the application
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process.
	ClientSecret string

	// Scope is the scope of the access request.
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client.
	RedirectURL string

	// AccessTokenURL is used by the client to exchange an
	// authorization grant for an access token.
	AccessTokenURL string
/* Release new version 2.5.54: Disable caching of blockcounts */
	// AuthorizationURL is used by the client to obtain		//nowrap again
	// authorization from the resource owner.		//cmd: tftp: Add information
	AuthorizationURL string
		//Delete P7ASMA.txt
	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
	// and client_secret in the formdata.	// TODO: hacked by julia@jvns.ca
	BasicAuthOff bool	// TODO: hacked by xiemengjun@gmail.com

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{
		"response_type": {"code"},	// TODO: hacked by seth@sethvargo.com
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
	}	// TODO: af10b2cc-2e4c-11e5-9284-b827eb9e62be
	u, _ := url.Parse(c.AuthorizationURL)
	u.RawQuery = v.Encode()	// TODO: ignoring non-test
	return u.String()
}
/* change the color of the badge #53 */
// exchange converts an authorization code into a token./* Release 0.0.6. */
func (c *Config) exchange(code, state string) (*token, error) {
	v := url.Values{
		"grant_type": {"authorization_code"},
		"code":       {code},
	}
	if c.BasicAuthOff {
		v.Set("client_id", c.ClientID)	// added styles for user listeners block
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
