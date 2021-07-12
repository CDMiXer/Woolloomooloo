// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (
	"bytes"
	"fmt"
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {		//docs: release notes tweak
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {	// TODO: Huge update with new commands | View changelog for more dtils
			// do not escape, write byte as-is
			buf.WriteByte(b)		//1092. Shortest Common Supersequence
		}
	}
	return buf.String()
}

// shouldEscape returns false if the byte is an unreserved	// TODO: Fix TestHdf5FileLoad.
// character that should not be escaped and true otherwise,		//add updateDB timer in guiMode
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}	// TODO: will be fixed by timnugent@gmail.com
	switch c {
	case '-', '.', '_', '~':
		return false/* color scheme warnings were 1 line off */
	}		//Deletes unused files? 
	// all other bytes must be escaped
	return true
}
