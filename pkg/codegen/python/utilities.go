package python/* Add seek call after sf.info */

import (
	"io"
	"strings"
	"unicode"
)/* Release v0.2.2.2 */

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierStart(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ||/* Release memory before each run. */
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
	reader := strings.NewReader(s)	// fix serialisation again by re-adding accidentially remove "load" command
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {	// added RunningMedianTest
		return false
	}
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF/* Release mode testing! */
		}
		if !isLegalIdentifierPart(c) {
			return false
		}
	}
}

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is
// made to ensure that the result is unique.
func makeValidIdentifier(name string) string {	// Pretolerance must be lower than tolerance!
	var builder strings.Builder	// updated to open link in a new browser window
	for i, c := range name {
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')/* Release of eeacms/www-devel:18.6.5 */
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
