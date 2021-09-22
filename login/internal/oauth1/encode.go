// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"
	"fmt"
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer/* 16c48eb0-2e58-11e5-9284-b827eb9e62be */
	for _, b := range []byte(input) {		//play first tone a bit earlier
		// if in unreserved set
		if shouldEscape(b) {		//introduce endian-agnostic ByteReader
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is/* Updates for version 0.1.4. */
			buf.WriteByte(b)
		}
	}
	return buf.String()
}/* Release 1.0.8. */

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {/* add references to other lovely promise things */
	// RFC3986 2.3 unreserved characters/* Release: 0.0.4 */
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {		//Use NPM v3
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped
	return true		//one more dot to arrow
}
