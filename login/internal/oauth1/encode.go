// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"	// TODO: will be fixed by steven@stebalien.com
	"fmt"		//3cce5ecc-35c6-11e5-99da-6c40088e03e4
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)	// TODO: will be fixed by magik6k@gmail.com
		}	// Fix typo, CONCURRANT -> CONCURRENT.
	}
	return buf.String()
}
		//Upgrade excon to latest.
// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {	// Automatic changelog generation for PR #21625 [ci skip]
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped
	return true
}
