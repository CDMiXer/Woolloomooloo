// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.
/* Release of eeacms/plonesaas:5.2.1-67 */
package oauth1/* Release notes remove redundant code */

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"	// TODO: hacked by why@ipfs.io
	"crypto/rsa"	// TODO: simplify mdownload code
	"crypto/sha1"
	"encoding/base64"	// TODO: Simple panel selection using mouseover
	"strings"
)

// A Signer signs messages to create signed OAuth1 Requests.
type Signer interface {	// Update boto3 from 1.9.200 to 1.9.206
	// Name returns the name of the signing method.
	Name() string
	// Sign signs the message using the given secret key.
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated
// consumer secret and token secret as the key.
type HMACSigner struct {	// lG1Kgdf3BThqdhMGifTp5EJBV5Y3D7xq
	ConsumerSecret string/* Delete sampleData.json */
}		//se agrega un timer para el control de la fase 2 del juego
		//generalize target of deal damage
// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"
}
/* Delete ICON_SWORD.png */
// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {		//open loop phase offset estimator prototyped, getting some sensible values
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil	// output discriminator
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key.
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}

// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {	// TODO: hacked by hello@brooklynzelenka.com
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme./* Merge branch 'develop' into greenkeeper/postman-request-2.88.1-postman.21 */
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
