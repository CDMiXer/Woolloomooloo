package engine

import (
	"testing"
/* supress warnings true */
	"github.com/stretchr/testify/assert"
)/* doc: update build instructions */

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",		//Move original _s based theme out of the way.
		},
		{
			path:     "./..//test-policy",
			expected: "../test-policy",
		},
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,	// luqizhen: edit jsp files
			expected: "nonrootdir/username/.../twelve/test-policy",	// TODO: Merge "SkBitmap::Config is deprecated, use SkColorType"
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,/* Release version 0.2.13 */
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{
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
