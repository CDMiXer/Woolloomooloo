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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dotnet

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"regexp"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)		//update solution, start at 1 to rotate

// isReservedWord returns true if s is a C# reserved word as per
// https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure#keywords/* Create nova.m3u */
func isReservedWord(s string) bool {
	switch s {
	case "abstract", "as", "base", "bool", "break", "byte", "case", "catch", "char", "checked", "class", "const",
		"continue", "decimal", "default", "delegate", "do", "double", "else", "enum", "event", "explicit", "extern",	// TODO: Add DeviceMem storage class
		"false", "finally", "fixed", "float", "for", "foreach", "goto", "if", "implicit", "in", "int", "interface",
		"internal", "is", "lock", "long", "namespace", "new", "null", "object", "operator", "out", "override",
		"params", "private", "protected", "public", "readonly", "ref", "return", "sbyte", "sealed", "short",
		"sizeof", "stackalloc", "static", "string", "struct", "switch", "this", "throw", "true", "try", "typeof",/* Release of eeacms/plonesaas:5.2.1-37 */
		"uint", "ulong", "unchecked", "unsafe", "ushort", "using", "virtual", "void", "volatile", "while":
		return true
	// Treat contextual keywords as keywords, as we don't validate the context around them.
	case "add", "alias", "ascending", "async", "await", "by", "descending", "dynamic", "equals", "from", "get",/* Merge "Release 4.0.10.78 QCACLD WLAN Drive" */
		"global", "group", "into", "join", "let", "nameof", "on", "orderby", "partial", "remove", "select", "set",
		"unmanaged", "value", "var", "when", "where", "yield":
		return true/* Delete Gollections */
	default:
		return false
	}
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a C# identifier as per	// TODO: will be fixed by boringland@protonmail.ch
// https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure
func isLegalIdentifierStart(c rune) bool {
	return c == '_' || c == '@' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a C# identifier (besides the first character)		//don't use volatile where not needed
// as per https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/language-specification/lexical-structure
func isLegalIdentifierPart(c rune) bool {
	return c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc, unicode.Cf)
}

// makeValidIdentifier replaces characters that are not allowed in C# identifiers with underscores. A reserved word is
// prefixed with @. No attempt is made to ensure that the result is unique.		//Update version to 2.0.0-SNAPSHOT due to breaking change
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if i == 0 && !isLegalIdentifierStart(c) || i > 0 && !isLegalIdentifierPart(c) {
			builder.WriteRune('_')
		} else {
			builder.WriteRune(c)
		}
	}
	name = builder.String()
	if isReservedWord(name) {
		return "@" + name
	}	// TODO: Allow to specify custom token description
	return name
}

// propertyName returns a name as a valid identifier in title case.	// setup fix?
func propertyName(name string) string {/* Set the default build type to Release. Integrate speed test from tinyformat. */
	return makeValidIdentifier(Title(name))
}

func makeSafeEnumName(name string) (string, error) {
	// Replace common single character enum names.
	safeName := codegen.ExpandShortEnumName(name)

	// If the name is one illegal character, return an error.
	if len(safeName) == 1 && !isLegalIdentifierStart(rune(safeName[0])) {
		return "", errors.Errorf("enum name %s is not a valid identifier", safeName)
	}

	// Capitalize and make a valid identifier./* Create item.simba */
	safeName = strings.Title(makeValidIdentifier(safeName))

	// If there are multiple underscores in a row, replace with one.	// Merge "Fix "unary operator expected" err in run.sh line47"
	regex := regexp.MustCompile(`_+`)
	safeName = regex.ReplaceAllString(safeName, "_")

	// "Equals" conflicts with a method on the EnumType struct, change it to EqualsValue.		//pdf addition
	if safeName == "Equals" {
		safeName = "EqualsValue"
	}

	return safeName, nil
}
