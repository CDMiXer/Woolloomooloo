// Copyright 2017 Drone.IO Inc. All rights reserved./* Release 0.95.209 */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2/* Released v1.0.0-alpha.1 */

import (
	"encoding/json"
	"net/http"
	"net/url"/* now producing Fix at specified offset, not just location */
	"strings"

	"github.com/drone/go-login/login/logger"
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`		//set defaults for better user experience from ABMOF paper
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`	// TODO: will be fixed by why@ipfs.io
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.
	Client *http.Client

	// ClientID is the identifier issued to the application
	// during the registration process.
	ClientID string		//menu links working

	// ClientSecret is the secret issued to the application	// TODO: bump to 0.19
	// during the registration process.
	ClientSecret string

	// Scope is the scope of the access request./* Prepare Release v3.8.0 (#1152) */
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client./* Released DirectiveRecord v0.1.27 */
	RedirectURL string

	// AccessTokenURL is used by the client to exchange an/* Released on PyPI as 0.9.9. */
	// authorization grant for an access token.
	AccessTokenURL string

	// AuthorizationURL is used by the client to obtain
	// authorization from the resource owner.
	AuthorizationURL string

	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id
	// and client_secret in the formdata.
	BasicAuthOff bool

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger
/* Added FloatMath.mix; */
	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.	// TODO: rolled off April, May
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{		//Updated outcome from experiment
		"response_type": {"code"},
		"client_id":     {c.ClientID},
	}
	if len(c.Scope) != 0 {
		v.Set("scope", strings.Join(c.Scope, " "))
	}
	if len(state) != 0 {
		v.Set("state", state)/* RESTEASY-699: Correct typo in MediaTypeHeaderDelegate. */
	}
	if len(c.RedirectURL) != 0 {
		v.Set("redirect_uri", c.RedirectURL)
	}
	u, _ := url.Parse(c.AuthorizationURL)/* Add Releases and Cutting version documentation back in. */
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
