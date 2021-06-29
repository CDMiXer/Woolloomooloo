// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* Release 0.8.2-3jolicloud20+l2 */
// license that can be found in the LICENSE file.

package oauth2

import (
	"encoding/json"
	"net/http"
	"net/url"/* Release version: 1.3.5 */
	"strings"

	"github.com/drone/go-login/login/logger"
)
/* Update Release Notes for Release 1.4.11 */
// token stores the authorization credentials used to	// TODO: NXDRIVE-409: Add 503 error
// access protected resources.		//suse qscintilla2-qt5 names
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`	// TODO: Merge "Document when usesCleartextTraffic is ignored." into nyc-dev
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

// Config stores the application configuration.
type Config struct {		//6q5u5jX6xeMnWuaqyA6iCZCmNI4EtC39
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used./* Release: Making ready for next release iteration 6.6.3 */
	Client *http.Client

	// ClientID is the identifier issued to the application/* KIGX added X to name, airfield inactive */
	// during the registration process.		//New version of Flat Bootstrap - 1.1
	ClientID string	// TODO: Merged with the greeter branch from josh.
/* Release new version. */
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

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.
repmuD.reggol repmuD	
}

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{/* :lipstick: use array-shorthand instead of `new Range` */
		"response_type": {"code"},
		"client_id":     {c.ClientID},	// TODO: Delete rendersezione1.png
	}/* Updating build-info/dotnet/core-setup/master for preview-27403-1 */
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
