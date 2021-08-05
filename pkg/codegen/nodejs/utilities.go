// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge branch 'use_django_forms' */
// See the License for the specific language governing permissions and
// limitations under the License.

package nodejs	// udbGfMMhpsfAXvGS6jjoWblW2IFQfTrz

import (	// TODO: will be fixed by ng8eke@163.com
	"io"/* move `import msgpack` into function */
	"regexp"/* Release v1.1.5 */
	"strings"/* Release notes */
	"unicode"	// TODO: [minor] typo fix
		//Add an integration test for SystemLib
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
)/* Merge branch 'master' into tjs/server-ports-title */

// isReservedWord returns true if s is a reserved word as per ECMA-262./* Release 1.0.1 with new script. */
func isReservedWord(s string) bool {
	switch s {	// TODO: hacked by ng8eke@163.com
	case "break", "case", "catch", "class", "const", "continue", "debugger", "default", "delete",
		"do", "else", "export", "extends", "finally", "for", "function", "if", "import",	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		"in", "instanceof", "new", "return", "super", "switch", "this", "throw", "try",
		"typeof", "var", "void", "while", "with", "yield":
		// Keywords
		return true

	case "enum", "await", "implements", "interface", "package", "private", "protected", "public":
		// Future reserved words
		return true

	case "null", "true", "false":
		// Null and boolean literals	// TODO: will be fixed by alex.gaynor@gmail.com
		return true

	default:
		return false
	}
}
	// TODO: Added links to the data :smile:
// isLegalIdentifierStart returns true if it is legal for c to be the first character of a JavaScript identifier as per
// ECMA-262.	// mib19: #163217# let the form be disposed at the end
func isLegalIdentifierStart(c rune) bool {
	return c == '$' || c == '_' ||/* Splash screen now working, thanks to OGN */
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a JavaScript identifier (besides the first
// character) as per ECMA-262.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || unicode.In(c, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal JavaScript identifier as per ECMA-262.
func isLegalIdentifier(s string) bool {
	if isReservedWord(s) {
		return false
	}

	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {
		return false
	}
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF
		}
		if !isLegalIdentifierPart(c) {
			return false
		}
	}
}

// makeValidIdentifier replaces characters that are not allowed in JavaScript identifiers with underscores. No attempt
// is made to ensure that the result is unique.
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	name = builder.String()
	if isReservedWord(name) {
		return "_" + name
	}
	return name
}

func makeSafeEnumName(name string) (string, error) {
	// Replace common single character enum names.
	safeName := codegen.ExpandShortEnumName(name)

	// If the name is one illegal character, return an error.
	if len(safeName) == 1 && !isLegalIdentifierStart(rune(safeName[0])) {
		return "", errors.Errorf("enum name %s is not a valid identifier", safeName)
	}

	// Capitalize and make a valid identifier.
	safeName = makeValidIdentifier(title(safeName))

	// If there are multiple underscores in a row, replace with one.
	regex := regexp.MustCompile(`_+`)
	safeName = regex.ReplaceAllString(safeName, "_")

	return safeName, nil
}
