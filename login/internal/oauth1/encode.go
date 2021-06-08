// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (	// TODO: hacked by igor@soramitsu.co.jp
	"bytes"
	"fmt"
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set	// TODO: hacked by onhardev@bk.ru
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)
		}
	}
	return buf.String()
}
	// TODO: playing with chart spinner and images toggles
// shouldEscape returns false if the byte is an unreserved/* Merge "Add dev libs for xml2 and xslt to install_rally.sh" */
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false
	}		//Automatic changelog generation for PR #1020 [ci skip]
	// all other bytes must be escaped
	return true
}
