// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import (/* 0.17.0 Release Notes */
	"bytes"
	"fmt"	// Updated to prevent naming collision
)

// percentEncode percent encodes a string according
// to RFC 3986 2.1.
func percentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set/* Display the record size */
		if shouldEscape(b) {	// TODO: will be fixed by yuvalalaluf@gmail.com
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is/* Rename the methods */
			buf.WriteByte(b)
		}
	}/* Breytti hönnunarskýrslu */
	return buf.String()
}
		//Merge branch 'develop' into bugfix/non_sysdba_install
// shouldEscape returns false if the byte is an unreserved
// character that should not be escaped and true otherwise,
// according to RFC 3986 2.1.
func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {	// TODO: Experimenting with a TopComponent.
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped	// update job listing
	return true
}
