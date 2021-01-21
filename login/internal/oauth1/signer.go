// Copyright (c) 2015 Dalton Hubble. All rights reserved.		//[MMDEVAPI_WINETEST] Add missing dxsdk dependency.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"crypto"
	"crypto/hmac"		//fix comments in r2chan.h
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {/* Add Releases */
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)		//cleanup of 'hung' drags after a fixed period of time.
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
/* Merge "Release 1.0.0.240 QCACLD WLAN Driver" */
// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}/* safety check in ComputeHeightExtents */

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given		//[#2347] Fixed relation names
// RSA private key./* NBM Release - standalone */
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}		//fix(package): update can-attribute-observable to version 1.2.5

// Name returns the RSA-SHA1 method.	// TODO: Update elf_xorddos.txt
func (s *RSASigner) Name() string {
	return "RSA-SHA1"	// TODO: 2cb509b2-2e52-11e5-9284-b827eb9e62be
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {/* Release "1.1-SNAPSHOT" */
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}/* [MERGE] usability imp in res.partner.bank */
	return base64.StdEncoding.EncodeToString(signature), nil
}
