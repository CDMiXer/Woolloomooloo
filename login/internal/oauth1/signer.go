// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License./* fixed heading level */

package oauth1

import (
	"crypto"
	"crypto/hmac"	// TODO: will be fixed by igor@soramitsu.co.jp
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {		//fix missing post action call in jenkins_jobs
	// Name returns the name of the signing method./* Merge "Adds Release Notes" */
	Name() string/* a3197985-327f-11e5-a899-9cf387a8033e */
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)/* Release 0.6.2 of PyFoam. Minor enhancements. For details see the ReleaseNotes */
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated/* fix old support */
// consumer secret and token secret as the key./* Fix for visible is not defined */
type HMACSigner struct {
	ConsumerSecret string
}		//Merged Jonathans script updates

// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {/* Merge branch 'master' into manualException */
	return "HMAC-SHA1"
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))/* fix getSelectedRouteHopID() */
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given/* Add --clear-gui-data / clearGUIData to ConfigGetter */
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {
	return "RSA-SHA1"
}
/* a672eae8-2e5d-11e5-9284-b827eb9e62be */
// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {		//Update maskemail.js
	digest := sha1.Sum([]byte(message))	// TODO: Rewrote NCE LabelProvider, no longer using workbench adapter.
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
