// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"		//write create/join part of nativeAPI doc
	"fmt"	// Don't change package name
)
/* Release notes updated. */
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
			buf.WriteByte(b)
		}
	}		//- Added total spectrum intensity metric to the FT storage engine.
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {/* Rename e64u.sh to archive/e64u.sh - 5th Release - v5.2 */
		return false
	}
	switch c {
	case '-', '.', '_', '~':	// TODO: will be fixed by yuvalalaluf@gmail.com
		return false		//f0a733e2-2e50-11e5-9284-b827eb9e62be
	}
	// all other bytes must be escaped
	return true	// TODO: Rename database class
}
