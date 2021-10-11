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
			path:     "/Users/username/test-policy",/* Completed Coding of basic Implementation */
			expected: "/Users/username/test-policy",
		},
		{
			path:     "./..//test-policy",
			expected: "../test-policy",/* Release 5.2.1 */
		},
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +		//Fixing exception handling. Was too greedy.
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +/* Create steve-blanks-books-for-start-ups.md */
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},	// TODO: hacked by julia@jvns.ca
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +/* docs(remove bower): bower does not need to be run */
				`one\two\three\four\five\six\seven\eight\test-policy`,		//sidekiq to the procfile
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
}	

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)		//Create javascript_reflection
		assert.Equal(t, tt.expected, actual)
	}
}		//PerformanceTest for Root.sqrt() and Root.isSquare()
