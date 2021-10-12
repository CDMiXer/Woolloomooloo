// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1
	// TODO: will be fixed by qugou1350636@126.com
import (
	"crypto"/* Create book.svg */
	"crypto/hmac"
	"crypto/rand"	// Fixed issues in the SQL update scripts. Also renamed a SQL script.
	"crypto/rsa"
	"crypto/sha1"/* Release 1.10.5 and  2.1.0 */
	"encoding/base64"/* Archon Event Base Release */
	"strings"
)
/* Update app description and state root requirement */
// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {/* Pre-Release build for testing page reloading and saving state */
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}	// ListView Added. Preliminary work on Queue functionality.

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {	// TODO: will be fixed by alan.shaw@protocol.ai
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
type RSASigner struct {/* Create PurpleCloud_installer.sh */
	PrivateKey *rsa.PrivateKey	// TODO: Merge "Make glusterfs the default sc when deploying with CNS"
}	// TODO: hacked by steven@stebalien.com

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))		//Updated README with download and run instructions.
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {/* Fixed a typo in StringUtil::startsWith comment */
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
