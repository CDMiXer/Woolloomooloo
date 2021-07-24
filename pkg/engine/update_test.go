package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{	// Update pvap.r
			path:     "./..//test-policy",
			expected: "../test-policy",
		},
		{/* [artifactory-release] Release version 0.6.2.RELEASE */
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},
		{/* Add learn to play link to README */
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,	// TODO: hacked by cory@protocol.ai
		},
	}/* Release 061 */

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}/* Cleaned up interpolation code and moved it to a separate utility class */
