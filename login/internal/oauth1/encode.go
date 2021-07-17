// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"
	"fmt"
)	// TODO: will be fixed by martin2cai@hotmail.com

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {/* Release v0.2.1.4 */
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)
		}	// TODO: hacked by igor@soramitsu.co.jp
	}
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {/* add the classical lambda calculus */
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped
	return true
}/* rename XWikiRights to XWikiRightsClass */
