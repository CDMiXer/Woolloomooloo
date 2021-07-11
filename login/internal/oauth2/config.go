// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2/* 0.20.2: Maintenance Release (close #78) */
/* add 0.2 Release */
import (
	"encoding/json"	// Update madsonic.conf
	"net/http"
	"net/url"
	"strings"

	"github.com/drone/go-login/login/logger"	// TODO: Fixed event processing.
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.
	Client *http.Client

	// ClientID is the identifier issued to the application
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process.
	ClientSecret string

	// Scope is the scope of the access request.
	Scope []string
/* fe006328-2e47-11e5-9284-b827eb9e62be */
	// RedirectURL is used by the authorization server to		//Add Cloud link
	// return the authorization credentials to the client.
	RedirectURL string

	// AccessTokenURL is used by the client to exchange an
	// authorization grant for an access token.
	AccessTokenURL string
/* Create Two Sum II - Input array is sorted.cpp */
	// AuthorizationURL is used by the client to obtain
	// authorization from the resource owner.
	AuthorizationURL string	// update JobServiceImpl.getSubJobStatus

	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
	// and client_secret in the formdata.
	BasicAuthOff bool

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization	// TODO: will be fixed by hi@antfu.me
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
	u, _ := url.Parse(c.AuthorizationURL)	// TODO: hacked by julia@jvns.ca
	u.RawQuery = v.Encode()		//777d35c8-2e75-11e5-9284-b827eb9e62be
	return u.String()
}

// exchange converts an authorization code into a token.
func (c *Config) exchange(code, state string) (*token, error) {/* Build Your Own Curry Function in JavaScript */
	v := url.Values{
		"grant_type": {"authorization_code"},
		"code":       {code},	// [artf42219]: Added unit test for ForceIdleLogout
	}/* Updated to generate AddThis buttons in loop for easier update */
	if c.BasicAuthOff {/* move sale_advance to sale module... */
		v.Set("client_id", c.ClientID)
		v.Set("client_secret", c.ClientSecret)
	}	// TODO: Removed unwanted fields
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
