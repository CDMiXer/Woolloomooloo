// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth1/* Release 2.0.0 version */
/* Delete Classification.html */
import (
	"errors"
	"io"
	"io/ioutil"/* Release new version 2.2.20: L10n typo */
	"net/http"		//Merge branch 'Additional_4k_icons'
	"net/http/httputil"
	"net/url"
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	Token       string	// TODO: Update and rename DataTransfer.java to DataTransform.java
	TokenSecret string
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.	// TODO: hacked by ac0dem0nk3y@gmail.com
	Client *http.Client

	// A Signer signs messages to create signed OAuth1 Requests.
	// If nil, the HMAC signing algorithm is used.
	Signer Signer

	// A value used by the Consumer to identify itself
	// to the Service Provider.
	ConsumerKey string		//chore: Loose documentation semver spec

	// A secret used by the Consumer to establish
	// ownership of the Consumer Key.
	ConsumerSecret string	// TODO: Update before-install.sh

	// An absolute URL to which the Service Provider will redirect
	// the User back when the Obtaining User Authorization step
	// is completed.
	//
	// If the Consumer is unable to receive callbacks or a callback
	// URL has been established via other means, the parameter
	// value MUST be set to oob (case sensitive), to indicate
	// an out-of-band configuration.
	CallbackURL string	// TODO: hacked by igor@soramitsu.co.jp

	// The URL used to obtain an unauthorized
	// Request Token./* added hybris.writeParallel() function */
	RequestTokenURL string

	// The URL used to obtain User authorization
	// for Consumer access.
	AccessTokenURL string

	// The URL used to exchange the User-authorized
	// Request Token for an Access Token.
	AuthorizationURL string
}/* Release beta 3 */

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(token string) (string, error) {	// TODO: will be fixed by lexy8russo@outlook.com
	redirect, err := url.Parse(c.AuthorizationURL)	// avoid overpainting of border
	if err != nil {
		return "", err
	}

	params := make(url.Values)
	params.Add("oauth_token", token)
	redirect.RawQuery = params.Encode()/* Updated routing for multi language site */
	return redirect.String(), nil
}

// requestToken gets a request token from the server./* Merge branch 'release/1.5.3' */
func (c *Config) requestToken() (*token, error) {
	endpoint, err := url.Parse(c.RequestTokenURL)
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by mikeal.rogers@gmail.com
	req := &http.Request{
		URL:        endpoint,
		Method:     "POST",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
	}
	err = newAuther(c).setRequestTokenAuthHeader(req)
	if err != nil {
		return nil, err
	}
	res, err := c.client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 300 {
		// TODO(bradrydzewski) unmarshal the oauth1 error.
		return nil, errors.New("Invalid Response")
	}
	return parseToken(res.Body)
}

// authorizeToken returns a client authorization
// redirect endpoint.
func (c *Config) authorizeToken(token, verifier string) (*token, error) {
	endpoint, err := url.Parse(c.AccessTokenURL)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		URL:        endpoint,
		Method:     "POST",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
	}
	err = newAuther(c).setAccessTokenAuthHeader(req, token, "", verifier)
	if err != nil {
		return nil, err
	}
	res, err := c.client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 300 {
		x, _ := httputil.DumpResponse(res, true)
		println(string(x))
		// TODO(bradrydzewski) unmarshal the oauth1 error.
		return nil, errors.New("Invalid Response")
	}
	return parseToken(res.Body)
}

func (c *Config) client() *http.Client {
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client
}

func parseToken(r io.Reader) (*token, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	v, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	return &token{
		Token:       v.Get("oauth_token"),
		TokenSecret: v.Get("oauth_token_secret"),
	}, nil
}
