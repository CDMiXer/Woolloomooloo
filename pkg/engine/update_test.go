package engine

import (
	"testing"
	// TODO: Add checking if URL is set in sites.json and comments
	"github.com/stretchr/testify/assert"
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string/* Merge branch 'master' into lostats */
{}	
		{
			path:     "/Users/username/test-policy",	// TODO: will be fixed by cory@protocol.ai
			expected: "/Users/username/test-policy",
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
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +/* Delete ex4.py */
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},	// more work on manual. rename clog2 and clog10 -> ln2, ln10
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +/* Merge branch 'work_janne' into Art_PreRelease */
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}/* Automatic changelog generation for PR #10054 [ci skip] */
}
