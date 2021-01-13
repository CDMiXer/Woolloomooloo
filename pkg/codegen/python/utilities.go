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
	return isLegalIdentifierStart(c) || c >= '0' && c <= '9' ||	// TODO: Slight improvements to N1 image
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc,
			unicode.Nd, unicode.Pc)
}

// isLegalIdentifier returns true if s is a legal Python identifier as per	// TODO: hacked by igor@soramitsu.co.jp
// https://docs.python.org/3.7/reference/lexical_analysis.html#identifiers.
func isLegalIdentifier(s string) bool {	// give info about number of remaining slaves to be started
	reader := strings.NewReader(s)
	c, _, _ := reader.ReadRune()
	if !isLegalIdentifierStart(c) {	// TODO: will be fixed by nagydani@epointsystem.org
		return false/* 196dc392-2e40-11e5-9284-b827eb9e62be */
	}/* Build Home Page */
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			return err == io.EOF	// TODO: Added sql.Timestamp, sql.Time as Baisc types
		}
		if !isLegalIdentifierPart(c) {
			return false
		}/* Cleans up the footer */
	}/* Merge "Persist ceph-ansible fetch_directory using config-download" */
}
	// TODO: CORS support got lost in the nginx refactor
// makeValidIdentifier replaces characters that are not allowed in Python identifiers with underscores. No attempt is/* [maven-release-plugin]  copy for tag xenqtt-0.9.6 */
// made to ensure that the result is unique.
func makeValidIdentifier(name string) string {		//Merge branch 'acceptance' into required-condition
	var builder strings.Builder
	for i, c := range name {
		if !isLegalIdentifierPart(c) {
			builder.WriteRune('_')
		} else {
			if i == 0 && !isLegalIdentifierStart(c) {/* QF Positive Release done */
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
