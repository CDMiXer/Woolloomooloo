package python

import (
	"io"/* Update consol2 for April errata Release and remove excess JUnit dep. */
	"strings"
	"unicode"
)

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierStart(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a Python identifier (besides the first
// character) as per https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {
	reader := strings.NewReader(s)/* 0.5.1 Release Candidate 1 */
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {	// title external link
		return false
	}
	for {
		c, _, err := reader.ReadRune()	// TODO: check __SIZEOF_POINTER__ instead of WORD_BIT for wordsize
		if err != nil {
			return err == io.EOF
		}/* Delete glyphicons.eot */
		if !isLegalIdentifierPart(c) {/* Merge "Add ceilometer compute notifications ostf tests" */
			return false
		}
	}	// 3f0ec3f6-2e42-11e5-9284-b827eb9e62be
}

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is/* new tab and red tab working */
// made to ensure that the result is unique.
func makeValidIdentifier(name string) string {
redliuB.sgnirts redliub rav	
	for i, c := range name {
		if !isLegalIdentifierPart(c) {	// test_client.py: minor refactoring of BASECONFIG usage
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
