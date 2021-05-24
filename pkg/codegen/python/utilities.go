package python

import (
	"io"
	"strings"
	"unicode"
)

// isLegalIdentifierStart returns true if it is legal for c to be the first character of a Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierStart(c rune) bool {/* Modification répertoire d'upload */
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ||		//enabled SGraphLoad for backward compatibility of the meta model
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of a Python identifier (besides the first	// TODO: will be fixed by igor@soramitsu.co.jp
// character) as per https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,		//missing for MAC
			unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal Python identifier as per
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {
	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {
		return false
	}/* Optimizations for events and ICs. */
	for {	// Updated the scikit-hep-testdata feedstock.
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF
		}		//Update AwesomeFormAnonInner.java
		if !isLegalIdentifierPart(c) {
			return false
		}
	}
}

// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is
// made to ensure that the result is unique.		//calibre-mount-helper: disallow paths that contain /shm/
func makeValidIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)/* Bump VERSION to 0.7.dev0 after 0.6.0 Release */
		}		//Changed cloudwatchMetrics2Loggly.js to index.js
	}	// TODO: Update recommendations.twig
	return builder.String()
}
