// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1/* Release 5.1.1 */
/* css reset and working on styling */
import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"/* [fa] Some clean up on rules and annotate to since version */
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)/* Release v0.5.0.5 */

// A Signer signs messages to create signed OAuth1 Requests.		//Delete HPX2MaxPlugin.py
type Signer interface {
	// Name returns the name of the signing method.
	Name() string/* remove #5191 */
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated/* [all] Release 7.1.4 */
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}

// Name returns the HMAC-SHA1 method./* Release version v0.2.7-rc008 */
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
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
}/* Merge "wlan: Release 3.2.3.85" */
		//Added eden/oauth required project
// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}/* Release of eeacms/www:20.1.11 */
