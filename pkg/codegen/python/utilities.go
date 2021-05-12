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
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a Python identifier (besides the first
// character) as per https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc)/* Merge "Refactoring: removed unused filters." into gb-ub-photos-bryce */
}
	// TODO: Adding AdjacentFileOutputStream
// isLegalIdentifier returns true if s is a legal Python identifier as per/* Merge "Release notes for recently added features" */
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {/* Updating build-info/dotnet/roslyn/dev16.0 for beta3-19102-02 */
	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {/* Add YesWeHack to partners */
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
	}/* Commit the properties for the 4.2 build. */
}

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is/* Update astroid from 1.6.5 to 2.0 */
// made to ensure that the result is unique./* 70cfcdde-2e4e-11e5-9284-b827eb9e62be */
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
	}/* Release of eeacms/forests-frontend:1.8-beta.2 */
	return builder.String()
}	// npm version
