// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.	// 020f6b88-2e4a-11e5-9284-b827eb9e62be

package oauth1

import (
	"bytes"
	"fmt"
)		//Create getmetadata.py

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set	// TODO: will be fixed by alex.gaynor@gmail.com
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))/* Create GDB.md */
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)
		}/* Released CachedRecord v0.1.1 */
	}
	return buf.String()		//- Working on agreements
}

// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1./* Release of cai-util-u3d v0.2.0 */
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false		//merged-in trunk r8291
	}
	// all other bytes must be escaped
	return true
}
