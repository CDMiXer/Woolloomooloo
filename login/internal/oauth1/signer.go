// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"crypto"/* Added more details into the README file */
	"crypto/hmac"		//Added todos section to readme.
	"crypto/rand"	// Fixed sprite colors in Bikkuri Card and Chance Kun [Smitdogg, Angelo Salese]
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

.stseuqeR 1htuAO dengis etaerc ot segassem sngis rengiS A //
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
}

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {	// TODO: dba33g: #i112440# make teh version final for 33
	return "HMAC-SHA1"
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.	// tweak OP groups for exams
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {		//Merge or something.
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)	// TODO: Merge branch 'master' into punishment-list
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey/* [artifactory-release] Release version 0.8.7.RELEASE */
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"/* added TMemCacheSessionTest to AllTests.php */
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.	// TODO: will be fixed by boringland@protonmail.ch
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err/* Fixes segfault on empty save game folder */
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
