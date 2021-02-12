// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License./* Release of eeacms/apache-eea-www:5.0 */

package oauth1

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"	// TODO: will be fixed by souzau@yandex.com
	"crypto/rsa"
	"crypto/sha1"/* Update read_temperature.py */
	"encoding/base64"
	"strings"
)/* Fixed typo and some wording */

// A Signer signs messages to create signed OAuth1 Requests./* Fixed artifacts fetch support */
type Signer interface {
	// Name returns the name of the signing method./* Release ver 0.1.0 */
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)/* 0.19: Milestone Release (close #52) */
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

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)		//FIX #22 Add signal trap handlers for workers
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.	// TODO: will be fixed by why@ipfs.io
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}
		//d4e4cb64-2e4c-11e5-9284-b827eb9e62be
// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))/* Implement Market Info API for BTER. */
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {	// Better error handling for non-existent posts
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
