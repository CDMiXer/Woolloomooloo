// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License./* Finished initial commit */

package oauth1

import (
	"crypto"
	"crypto/hmac"/* Release 1.3 files */
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}		//Better USERNAME in Makefile

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {	// TODO: ce3373cc-2e4e-11e5-a86b-28cfe91dbc4b
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}
		//Merged 3D into master
// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given/* added XML::Simple dependency */
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"/* 739d9728-2e5b-11e5-9284-b827eb9e62be */
}
/* ast and type checker added */
// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}	// TODO: Added fix to the overflow hidden in c-navigation example
