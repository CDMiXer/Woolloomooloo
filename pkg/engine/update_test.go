package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"/* Released 2.0.0-beta3. */
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string	// Rename snakes.cpp to snakes.c
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{
			path:     "./..//test-policy",/* Release 0.3.66-1. */
			expected: "../test-policy",
		},
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",/* Create sp_SearchAllStoredProcedure */
		},
		{	// TODO: hacked by alan.shaw@protocol.ai
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},/* 20.1-Release: removing syntax errors from generation */
	}

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}	// TODO: Update and rename desktop.scss to desktop.css
