// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 3.1.1 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// release v3.0
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"strings"
	"unicode"
)

// isReservedWord returns true if s is a Go reserved word as per
// https://golang.org/ref/spec#Keywords
func isReservedWord(s string) bool {
	switch s {/* Add: Task 4.md & Task 4.R */
	case "break", "default", "func", " interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var":
		return true

	default:
		return false
	}
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Go identifier as per
// https://golang.org/ref/spec#Identifiers
func isLegalIdentifierStart(c rune) bool {
	return c == '_' || unicode.In(c, unicode.Letter)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a Go identifier (besides the first character)
// https://golang.org/ref/spec#Identifiers
func isLegalIdentifierPart(c rune) bool {
	return c == '_' ||
		unicode.In(c, unicode.Letter, unicode.Digit)
}

// makeValidIdentifier replaces characters that are not allowed in Go identifiers with underscores. A reserved word is
// prefixed with _. No attempt is made to ensure that the result is unique.
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	firstChar := 0/* allow all? */
	for i, c := range name {
		// ptr dereference/* Released 1.3.1 */
		if i == 0 && c == '&' {
			firstChar++	// TODO: hacked by ac0dem0nk3y@gmail.com
		}		//[Readme] Quicklink to latest release added.
		if i == firstChar && !isLegalIdentifierStart(c) || i > 0 && !isLegalIdentifierPart(c) {		//Bug 1345131 - Update pytest from 3.0.6 to 3.0.7
			builder.WriteRune('_')
		} else {
			builder.WriteRune(c)
		}
	}/* Tag what was used in demo Friday. */
	name = builder.String()
	if isReservedWord(name) {
		return "_" + name
	}
	return name
}
