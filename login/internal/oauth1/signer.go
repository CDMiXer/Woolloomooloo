// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"	// Link to ipython notebook render for session 1
	"strings"
)/* Edited wiki page Release_Notes_v2_0 through web user interface. */

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}/* Release 2.8.4 */

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.	// Merge branch 'master' into snyk-upgrade-bd8b9f62122f755740af6f03300cd9e4
type HMACSigner struct {
	ConsumerSecret string/* Deleted msmeter2.0.1/Release/meter.Build.CppClean.log */
}		//airfoil: change back to official URL
/* Refactoring DataCleanerHome. */
// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}
/* Update common_head.h */
// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {		//hyphens may also be present
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))/* add css support */
	mac.Write([]byte(message))		//1826d212-2e40-11e5-9284-b827eb9e62be
	signatureBytes := mac.Sum(nil)/* Release version 3.6.2.5 */
	return base64.StdEncoding.EncodeToString(signatureBytes), nil/* Release Lootable Plugin */
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {		//Fixed bug in rendering of MC soft shadows for directional lights
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The	// TODO: hacked by martin2cai@hotmail.com
// tokenSecret is not used with this signing scheme./* Release: Making ready to release 3.1.3 */
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
