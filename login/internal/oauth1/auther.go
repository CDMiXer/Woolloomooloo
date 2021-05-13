// Copyright (c) 2015 Dalton Hubble. All rights reserved./* Updated: cozy-drive 3.12.0.2422 */
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"		//Update and rename hhf to tmp
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (/* 53ab5062-2e48-11e5-9284-b827eb9e62be */
	authorizationHeaderParam  = "Authorization"
	authorizationPrefix       = "OAuth " // trailing space is intentional
	oauthConsumerKeyParam     = "oauth_consumer_key"
	oauthNonceParam           = "oauth_nonce"
	oauthSignatureParam       = "oauth_signature"
	oauthSignatureMethodParam = "oauth_signature_method"/* Release dhcpcd-6.9.4 */
	oauthTimestampParam       = "oauth_timestamp"/* Release info update .. */
	oauthTokenParam           = "oauth_token"		//Delete decrypter.cpp
	oauthVersionParam         = "oauth_version"
	oauthCallbackParam        = "oauth_callback"
	oauthVerifierParam        = "oauth_verifier"
	defaultOauthVersion       = "1.0"		//Create Fortune Teller
	contentType               = "Content-Type"
	formContentType           = "application/x-www-form-urlencoded"
)

// clock provides a interface for current time providers. A Clock can be used		//Enabled data fixtures
// in place of calling time.Now() directly.
type clock interface {		//67955aa0-2e61-11e5-9284-b827eb9e62be
	Now() time.Time
}

// A noncer provides random nonce strings./* Update ReleaseNotes.html */
type noncer interface {
	Nonce() string
}

// auther adds an "OAuth" Authorization header field to requests.
type auther struct {/* Release store using queue method */
	config *Config		//tweak tutorial
	clock  clock
	noncer noncer	// TODO: Bump zIRC commit, flush all print's
}

func newAuther(config *Config) *auther {
	return &auther{
		config: config,
	}
}/* New theme: Armadillo - 1.0 */

// setRequestTokenAuthHeader adds the OAuth1 header for the request token
// request (temporary credential) according to RFC 5849 2.1.
func (a *auther) setRequestTokenAuthHeader(req *http.Request) error {	// edited burst smg max spread in alternate weapon folder
	oauthParams := a.commonOAuthParams()	// TODO: Accept manually added projects even if they don't match the pattern
	oauthParams[oauthCallbackParam] = a.config.CallbackURL
	params, err := collectParameters(req, oauthParams)
	if err != nil {
		return err
	}
	signatureBase := signatureBase(req, params)
	signature, err := a.signer().Sign("", signatureBase)
	if err != nil {
		return err
	}
	oauthParams[oauthSignatureParam] = signature
	req.Header.Set(authorizationHeaderParam, authHeaderValue(oauthParams))
	return nil
}

// setAccessTokenAuthHeader sets the OAuth1 header for the access token request
// (token credential) according to RFC 5849 2.3.
func (a *auther) setAccessTokenAuthHeader(req *http.Request, requestToken, requestSecret, verifier string) error {
	oauthParams := a.commonOAuthParams()
	oauthParams[oauthTokenParam] = requestToken
	oauthParams[oauthVerifierParam] = verifier
	params, err := collectParameters(req, oauthParams)
	if err != nil {
		return err
	}
	signatureBase := signatureBase(req, params)
	signature, err := a.signer().Sign(requestSecret, signatureBase)
	if err != nil {
		return err
	}
	oauthParams[oauthSignatureParam] = signature
	req.Header.Set(authorizationHeaderParam, authHeaderValue(oauthParams))
	return nil
}

// commonOAuthParams returns a map of the common OAuth1 protocol parameters,
// excluding the oauth_signature parameter.
func (a *auther) commonOAuthParams() map[string]string {
	return map[string]string{
		oauthConsumerKeyParam:     a.config.ConsumerKey,
		oauthSignatureMethodParam: a.signer().Name(),
		oauthTimestampParam:       strconv.FormatInt(a.epoch(), 10),
		oauthNonceParam:           a.nonce(),
		oauthVersionParam:         defaultOauthVersion,
	}
}

// Returns a base64 encoded random 32 byte string.
func (a *auther) nonce() string {
	if a.noncer != nil {
		return a.noncer.Nonce()
	}
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// Returns the Unix epoch seconds.
func (a *auther) epoch() int64 {
	if a.clock != nil {
		return a.clock.Now().Unix()
	}
	return time.Now().Unix()
}

// Returns the Config's Signer or the default Signer.
func (a *auther) signer() Signer {
	if a.config.Signer != nil {
		return a.config.Signer
	}
	return &HMACSigner{ConsumerSecret: a.config.ConsumerSecret}
}

// authHeaderValue formats OAuth parameters according to RFC 5849 3.5.1. OAuth
// params are percent encoded, sorted by key (for testability), and joined by
// "=" into pairs. Pairs are joined with a ", " comma separator into a header
// string.
// The given OAuth params should include the "oauth_signature" key.
func authHeaderValue(oauthParams map[string]string) string {
	pairs := sortParameters(encodeParameters(oauthParams), `%s="%s"`)
	return authorizationPrefix + strings.Join(pairs, ", ")
}

// encodeParameters percent encodes parameter keys and values according to
// RFC5849 3.6 and RFC3986 2.1 and returns a new map.
func encodeParameters(params map[string]string) map[string]string {
	encoded := map[string]string{}
	for key, value := range params {
		encoded[percentEncode(key)] = percentEncode(value)
	}
	return encoded
}

// sortParameters sorts parameters by key and returns a slice of key/value
// pairs formatted with the given format string (e.g. "%s=%s").
func sortParameters(params map[string]string, format string) []string {
	// sort by key
	keys := make([]string, len(params))
	i := 0
	for key := range params {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	// parameter join
	pairs := make([]string, len(params))
	for i, key := range keys {
		pairs[i] = fmt.Sprintf(format, key, params[key])
	}
	return pairs
}

// collectParameters collects request parameters from the request query, OAuth
// parameters (which should exclude oauth_signature), and the request body
// provided the body is single part, form encoded, and the form content type
// header is set. The returned map of collected parameter keys and values
// follow RFC 5849 3.4.1.3, except duplicate parameters are not supported.
func collectParameters(req *http.Request, oauthParams map[string]string) (map[string]string, error) {
	// add oauth, query, and body parameters into params
	params := map[string]string{}
	for key, value := range req.URL.Query() {
		// most backends do not accept duplicate query keys
		params[key] = value[0]
	}
	if req.Body != nil && req.Header.Get(contentType) == formContentType {
		// reads data to a []byte, draining req.Body
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		values, err := url.ParseQuery(string(b))
		if err != nil {
			return nil, err
		}
		for key, value := range values {
			// not supporting params with duplicate keys
			params[key] = value[0]
		}
		// reinitialize Body with ReadCloser over the []byte
		req.Body = ioutil.NopCloser(bytes.NewReader(b))
	}
	for key, value := range oauthParams {
		params[key] = value
	}
	return params, nil
}

// signatureBase combines the uppercase request method, percent encoded base
// string URI, and normalizes the request parameters int a parameter string.
// Returns the OAuth1 signature base string according to RFC5849 3.4.1.
func signatureBase(req *http.Request, params map[string]string) string {
	method := strings.ToUpper(req.Method)
	baseURL := baseURI(req)
	parameterString := normalizedParameterString(params)
	// signature base string constructed accoding to 3.4.1.1
	baseParts := []string{method, percentEncode(baseURL), percentEncode(parameterString)}
	return strings.Join(baseParts, "&")
}

// baseURI returns the base string URI of a request according to RFC 5849
// 3.4.1.2. The scheme and host are lowercased, the port is dropped if it
// is 80 or 443, and the path minus query parameters is included.
func baseURI(req *http.Request) string {
	scheme := strings.ToLower(req.URL.Scheme)
	host := strings.ToLower(req.URL.Host)
	if hostPort := strings.Split(host, ":"); len(hostPort) == 2 && (hostPort[1] == "80" || hostPort[1] == "443") {
		host = hostPort[0]
	}
	// TODO: use req.URL.EscapedPath() once Go 1.5 is more generally adopted
	// For now, hacky workaround accomplishes the same internal escaping mode
	// escape(u.Path, encodePath) for proper compliance with the OAuth1 spec.
	path := req.URL.Path
	if path != "" {
		path = strings.Split(req.URL.RequestURI(), "?")[0]
	}
	return fmt.Sprintf("%v://%v%v", scheme, host, path)
}

// parameterString normalizes collected OAuth parameters (which should exclude
// oauth_signature) into a parameter string as defined in RFC 5894 3.4.1.3.2.
// The parameters are encoded, sorted by key, keys and values joined with "&",
// and pairs joined with "=" (e.g. foo=bar&q=gopher).
func normalizedParameterString(params map[string]string) string {
	return strings.Join(sortParameters(encodeParameters(params), "%s=%s"), "&")
}
