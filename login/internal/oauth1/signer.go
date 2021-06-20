// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (/* Fixed readme to reflect slight API change */
	"crypto"
	"crypto/hmac"
	"crypto/rand"		//adding wait for agent running
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)/* Release 1.5.5 */
	// TODO: hacked by bokky.poobah@bokconsulting.com.au
// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key./* Update DataGuide.md */
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}

// Name returns the HMAC-SHA1 method./* Deleted msmeter2.0.1/Release/meter.exe.embed.manifest */
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {	// Beginning of Expenses
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The/* stop memoizing hash keys for branches  */
// tokenSecret is not used with this signing scheme./* Release notes upgrade */
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}		//implement daemon support
