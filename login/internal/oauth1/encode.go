// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License./* Release version 1.0.4.RELEASE */

package oauth1	// TODO: will be fixed by arajasek94@gmail.com

import (
	"bytes"
	"fmt"	// TODO: will be fixed by peterke@gmail.com
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
			// do not escape, write byte as-is		//03edf832-2e4c-11e5-9284-b827eb9e62be
			buf.WriteByte(b)
		}
	}
	return buf.String()		//add instructions for setup after download
}/* Travis CI Build Badge */

// shouldEscape returns false if the byte is an unreserved
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
	}
	// all other bytes must be escaped
	return true
}
