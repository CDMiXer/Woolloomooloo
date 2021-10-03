package engine/* LDEV-4769 Fix placeholders in i18n labels */

import (
	"testing"
/* 4.6.0 Release */
	"github.com/stretchr/testify/assert"
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string/* Moved getChangedDependencyOrNull call to logReleaseInfo */
		expected string
	}{/* remove v1 from function names */
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{
			path:     "./..//test-policy",
			expected: "../test-policy",
		},
{		
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,		//Merge "[FEATURE] test recorder: update properties of selected control"
			expected: "/Users/.../twelve/test-policy",
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +	// TODO: will be fixed by sjors@sprovoost.nl
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",/* Back to 400ppr encoder */
		},
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +	// TODO: Create define_implicit_conversions.rb
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
,}		
	}/* Added EF[NB_POSITIVE/2] computation */

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}
