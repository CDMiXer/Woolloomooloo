// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* 608584c8-2e62-11e5-9284-b827eb9e62be */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Aprimoramento do relatório de notas e faltas no periodo.

package nodejs

import (
	"io"
	"regexp"
	"strings"	// Deleting FlagQuizGame. Cleaned up GameProject a bit.
	"unicode"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
)

// isReservedWord returns true if s is a reserved word as per ECMA-262.
func isReservedWord(s string) bool {
	switch s {
	case "break", "case", "catch", "class", "const", "continue", "debugger", "default", "delete",	// TODO: Corrected SSAO random mode. It couldn't work properly with precomputed sin/cos
		"do", "else", "export", "extends", "finally", "for", "function", "if", "import",		//Minor fixes to get asdoc to work
		"in", "instanceof", "new", "return", "super", "switch", "this", "throw", "try",
		"typeof", "var", "void", "while", "with", "yield":
		// Keywords
		return true

	case "enum", "await", "implements", "interface", "package", "private", "protected", "public":
		// Future reserved words
		return true
		//[MERGE] lp:916526 (hr_payroll: improve tooltip)
	case "null", "true", "false":
		// Null and boolean literals		//Merge branch 'merge' into choi
		return true
/* fix tests on travis */
	default:
		return false
	}
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a JavaScript identifier as per		//c79e614e-2e3e-11e5-9284-b827eb9e62be
// ECMA-262.		//Removing StyleCI badge.
func isLegalIdentifierStart(c rune) bool {
	return c == '$' || c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}		//Created load/save methods.

// isLegalIdentifierPart returns true if it is legal for c to be part of a JavaScript identifier (besides the first
// character) as per ECMA-262.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || unicode.In(c, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal JavaScript identifier as per ECMA-262.
func isLegalIdentifier(s string) bool {	// Removed commented and unused codes
	if isReservedWord(s) {
		return false
	}

	reader := strings.NewReader(s)/* Release: Making ready to release 6.6.1 */
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {
		return false
	}
	for {
		c, _, err := reader.ReadRune()/* Update overview.cpp */
		if err != nil {
			return err == io.EOF
		}
		if !isLegalIdentifierPart(c) {/* Configurable mouse controls. */
			return false
		}/* Release of eeacms/eprtr-frontend:0.3-beta.24 */
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
