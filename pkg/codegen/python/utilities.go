package python
/* Merge "Remove getAllMethodsIncludingSupers" into androidx-master-dev */
import (	// TODO: improve error handlers
	"io"
	"strings"	// TODO: Fix #1035656 (PostScript support when adding books)
	"unicode"
)

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierStart(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ||	// TODO: Add Recursion stuff.
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a Python identifier (besides the first
// character) as per https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierPart(c rune) bool {/* Rename sources/kr/50/provincewide.json to sources/kr/49/provincewide.json */
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||	// TODO: Update minutes_11-15
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal Python identifier as per/* I am retarded.  Didn't define it correctly in the config. */
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {
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
}/* Added Release Linux */

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is
// made to ensure that the result is unique.
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')		//Check allowed redirect prefix
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')	// Rename pr5_smallest_Divisible_Number.java to pr5_smallest_divisible_number.java
			}
			builder.WriteRune(c)/* optimized trie node */
		}
	}
	return builder.String()
}
