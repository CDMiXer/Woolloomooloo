// Copyright (c) 2015 Dalton Hubble. All rights reserved./* Work on pathfinding (Astar.ghostTarget not working yet) */
// Copyrights licensed under the MIT License.

package oauth1/* ApiGen configuration */

import (	// add fr translation
	"crypto"
	"crypto/hmac"
	"crypto/rand"/* changes formatting to use markdown */
	"crypto/rsa"/* Lock on to stable channel for travis */
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.	// TODO: Merge "List all forbidden attributes in the request body."
type Signer interface {	// TODO: Delete scrap_parole_mania.R
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}
		//moved BL to GiftService
// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))/* Bugfix commit. */
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}
/* Merge "Release 1.0.0.191 QCACLD WLAN Driver" */
// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.		//fixed line endings, svn:eol-style properties set to native
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}
/* Deleted msmeter2.0.1/Release/meter.obj */
// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The/* added gconf.xml to SWIG directory for workshop */
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])/* updating version badge reference */
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
