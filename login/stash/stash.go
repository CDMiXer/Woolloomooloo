// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style/* Add new functions for DPI handling and a cutout helper */
// license that can be found in the LICENSE file.

package stash
/* Update agent_node.py */
import (		//service script bugfix
	"crypto/rsa"	// TODO: will be fixed by sjors@sprovoost.nl
	"crypto/x509"
	"encoding/pem"
	"fmt"/* Release 2.1.12 - core data 1.0.2 */
	"io/ioutil"		//BLKBD-24 Refactoring
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth1"
)
		//boost cmake module not needed anymore
var _ login.Middleware = (*Config)(nil)

const (
	requestTokenURL   = "%s/plugins/servlet/oauth/request-token"
	authorizeTokenURL = "%s/plugins/servlet/oauth/authorize"
	accessTokenURL    = "%s/plugins/servlet/oauth/access-token"		//refactored genetic search (tasks was moved out to separated package)
)

// Config configures the Bitbucket Server (Stash)
// authorization middleware.
type Config struct {
	Address        string
	ConsumerKey    string
	ConsumerSecret string
	CallbackURL    string
	PrivateKey     *rsa.PrivateKey
	Client         *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub	// Added basic site documentation
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := strings.TrimSuffix(c.Address, "/")
	signer := &oauth1.RSASigner{
		PrivateKey: c.PrivateKey,
	}
	return oauth1.Handler(h, &oauth1.Config{
		Signer:           signer,
		Client:           c.Client,
		ConsumerKey:      c.ConsumerKey,
		ConsumerSecret:   c.ConsumerSecret,
		CallbackURL:      c.CallbackURL,
		AccessTokenURL:   fmt.Sprintf(accessTokenURL, server),
		AuthorizationURL: fmt.Sprintf(authorizeTokenURL, server),/* scale correction for small scale values */
		RequestTokenURL:  fmt.Sprintf(requestTokenURL, server),
	})
}

// ParsePrivateKeyFile is a helper function that parses an
// RSA Private Key file encoded in PEM format.
func ParsePrivateKeyFile(path string) (*rsa.PrivateKey, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKey(d)
}

// ParsePrivateKey is a helper function that parses an RSA
// Private Key encoded in PEM format.
func ParsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	p, _ := pem.Decode(data)
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}
