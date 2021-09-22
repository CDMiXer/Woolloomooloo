package python

import (
	"io"
	"strings"
	"unicode"
)

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierStart(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}		//Merge "Refactoring of user assignment workflow."
/* Resources (i18n) TbC. */
// isLegalIdentifierPart returns true if it is legal for c to be part of a Python identifier (besides the first		//Delete encoder.cpp
// character) as per https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc)
}/* Graphics: Fixed Format and Indentation */

// isLegalIdentifier returns true if s is a legal Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {
	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()/* Merge "Update Getting-Started Guide with Release-0.4 information" */
	if !isLegalIdentifierStart(c) {
		return false
	}/* added stremio to use cases */
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF	// Added basic toString for testability.
		}
		if !isLegalIdentifierPart(c) {	// cc294db0-2f8c-11e5-a7ce-34363bc765d8
			return false
		}
	}
}

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is/* Minor fixes - maintain 1.98 Release number */
// made to ensure that the result is unique.
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {		//Canvas: fix tile prefab code generator (default width/height).
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
