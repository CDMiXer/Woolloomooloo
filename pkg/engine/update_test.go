package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: Disable only additional warnings in gsl.
)		//Delete ReadNames-to-FASTA_V8.py

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{
			path:     "./..//test-policy",
			expected: "../test-policy",	// TODO: merge bzr.dev r3975
		},/* Merge branch 'master' into pack1 */
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},
		{	// TODO: hacked by ligi@ligi.de
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,/* Merge "CSSMin: Add tests for handling existing data: URIs" */
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{		//small fixes (p1)
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +		//By status blotter filter
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{	// TODO: will be fixed by ligi@ligi.de
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}
