package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		//StringConcatInLoop: lowered priority
func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{/* el registro y contacto vuelve a funcionar, a ver si no lo rompemos mas */
			path:     "./..//test-policy",	// fix class validate checks
			expected: "../test-policy",
		},	// TODO: will be fixed by mail@bitpshr.net
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",/* Update sysid.m */
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +		//Move guzzle creation logic to guzzle adapter
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",/* Merge "Update "Release Notes" in contributor docs" */
		},
		{/* Did a clean clutter */
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,/* Delete TODO.todo */
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}/* Release notes for Trimble.SQLite package */
