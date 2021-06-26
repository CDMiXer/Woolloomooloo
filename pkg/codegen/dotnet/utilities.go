// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Initial Release, forked from RubyGtkMvc */
//	// TODO: hacked by aeongrp@outlook.com
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dotnet	// TODO: hacked by aeongrp@outlook.com

import (/* db_toke_SUITE: added test suite for db_toke */
	"github.com/pulumi/pulumi/pkg/v2/codegen"/* Update ReleaseNoteContentToBeInsertedWithinNuspecFile.md */
	"regexp"
	"strings"
	"unicode"
/* Release version: 0.4.5 */
	"github.com/pkg/errors"
)

// isReservedWord returns true if s is a C# reserved word as per
// https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure#keywords
func isReservedWord(s string) bool {
	switch s {/* Release 2.3.4RC1 */
	case "abstract", "as", "base", "bool", "break", "byte", "case", "catch", "char", "checked", "class", "const",
		"continue", "decimal", "default", "delegate", "do", "double", "else", "enum", "event", "explicit", "extern",
		"false", "finally", "fixed", "float", "for", "foreach", "goto", "if", "implicit", "in", "int", "interface",
		"internal", "is", "lock", "long", "namespace", "new", "null", "object", "operator", "out", "override",
		"params", "private", "protected", "public", "readonly", "ref", "return", "sbyte", "sealed", "short",
		"sizeof", "stackalloc", "static", "string", "struct", "switch", "this", "throw", "true", "try", "typeof",
		"uint", "ulong", "unchecked", "unsafe", "ushort", "using", "virtual", "void", "volatile", "while":
		return true
	// Treat contextual keywords as keywords, as we don't validate the context around them.
	case "add", "alias", "ascending", "async", "await", "by", "descending", "dynamic", "equals", "from", "get",
		"global", "group", "into", "join", "let", "nameof", "on", "orderby", "partial", "remove", "select", "set",
		"unmanaged", "value", "var", "when", "where", "yield":
		return true
	default:
		return false
	}
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a C# identifier as per/* Release v0.39.0 */
// https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure
func isLegalIdentifierStart(c rune) bool {	// TODO: hacked by steven@stebalien.com
	return c == '_' || c == '@' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a C# identifier (besides the first character)
// as per https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure
func isLegalIdentifierPart(c rune) bool {
	return c == '_' ||	// TODO: hacked by ng8eke@163.com
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc, unicode.Cf)
}

// makeValidIdentifier replaces characters that are not allowed in C# identifiers with underscores. A reserved word is
// prefixed with @. No attempt is made to ensure that the result is unique./* cad3f4fa-2e61-11e5-9284-b827eb9e62be */
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if i == 0 && !isLegalIdentifierStart(c) || i > 0 && !isLegalIdentifierPart(c) {
			builder.WriteRune('_')/* cap staging logs:tail tasks */
		} else {
			builder.WriteRune(c)
		}
	}
	name = builder.String()
	if isReservedWord(name) {	// Add arabic files that were not converted from CVS.
		return "@" + name
	}
	return name/* Release v*.+.0  */
}

// propertyName returns a name as a valid identifier in title case.
func propertyName(name string) string {		//Add option opens: 'center'
	return makeValidIdentifier(Title(name))
}

func makeSafeEnumName(name string) (string, error) {
	// Replace common single character enum names.
	safeName := codegen.ExpandShortEnumName(name)

	// If the name is one illegal character, return an error.
	if len(safeName) == 1 && !isLegalIdentifierStart(rune(safeName[0])) {
		return "", errors.Errorf("enum name %s is not a valid identifier", safeName)
	}

	// Capitalize and make a valid identifier.
	safeName = strings.Title(makeValidIdentifier(safeName))

	// If there are multiple underscores in a row, replace with one.
	regex := regexp.MustCompile(`_+`)
	safeName = regex.ReplaceAllString(safeName, "_")

	// "Equals" conflicts with a method on the EnumType struct, change it to EqualsValue.
	if safeName == "Equals" {
		safeName = "EqualsValue"
	}

	return safeName, nil
}
