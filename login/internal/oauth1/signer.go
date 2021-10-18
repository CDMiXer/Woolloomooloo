// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License./* Update readme.  */

package oauth1

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"/* Move CAN tools to new location */
	"encoding/base64"		//1aad1b23-2e9d-11e5-a839-a45e60cdfd11
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {	// Commiting latest changes for v1.14
	// Name returns the name of the signing method./* commands: actually implement --closed for topological heads */
	Name() string/* Added recordTypeHolder to DataValidator */
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)	// Added documentation for "mine craft" script
}	// TODO: hacked by juan@benet.ai

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string/* add tree demo del exe file */
}

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"	// TODO: upload files to GH
}/* Static checks fixes. Release preparation */

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil/* Release V8.1 */
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}/* profile pic , cover pic upload fix and date of birth fix at signup. */

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
}
