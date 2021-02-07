// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"/* less confusing */
	"crypto/rsa"	// TODO: Delete hn-react.html
	"crypto/sha1"
	"encoding/base64"
	"strings"		//Added Splash Screen
)

// A Signer signs messages to create signed OAuth1 Requests.	// [FIX]:lp-641084 Warning message wrongly indication
type Signer interface {
	// Name returns the name of the signing method.
	Name() string	// TODO: #POULPE-471 Make the backup filename more informative
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}/* Fix java classes code format */

// Name returns the HMAC-SHA1 method./* Update PrepareReleaseTask.md */
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}	// remove real-time pull of xpad.c

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {/* Add test for validator returning data on success */
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))/* Merge "Padding between date and digital clock" into jb-mr1.1-dev */
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method./* The acting user is got with every execution of a request. */
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
	return base64.StdEncoding.EncodeToString(signature), nil/* Released Neo4j 3.3.7 */
}
