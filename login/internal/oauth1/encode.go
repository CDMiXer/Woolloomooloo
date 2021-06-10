// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.	// TODO: will be fixed by praveen@minio.io
/* Initial Release - Supports only Wind Symphony */
package oauth1

import (
	"bytes"
	"fmt"
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1./* Add Releases and Cutting version documentation back in. */
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
	}
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved	// TODO: will be fixed by zaq1tomo@gmail.com
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
	return true/* Release build working on Windows; Deleted some old code. */
}/* Delete fhqTreap.cpp */
