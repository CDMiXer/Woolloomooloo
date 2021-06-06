package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"/* Delete 09_part_iii_sql_soccer.md */
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},	// TODO: will be fixed by hugomrdias@gmail.com
		{/* 269b6bde-2e52-11e5-9284-b827eb9e62be */
			path:     "./..//test-policy",
			expected: "../test-policy",
		},
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +		//Merge "Sort jobs in JobsView based on created_at value"
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",
		},/* Rename .htaccess to public/.htaccess */
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +/* Merge branch 'master' of https://github.com/yangboz/bearded-shame.git */
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,/* Final San Diego Code */
			expected: "nonrootdir/username/.../twelve/test-policy",
		},		//Array based queue 
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",	// Make a bunch of symbols private.
		},
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {/* Delete Update-Release */
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}
