package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)/* minor code error fix on readme */

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {	// TODO: will be fixed by arajasek94@gmail.com
		path     string		//add requirement software.
		expected string
	}{/* Reenable ControlService and fix syntax errors in svcctl.idl. */
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",
		},
		{	// TODO: Test the Sentence Separator in the JMA_Knowledge
			path:     "./..//test-policy",
			expected: "../test-policy",	// readme: inheritance example
		},
{		
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,/* Create tile0.png */
			expected: "/Users/.../twelve/test-policy",
		},
		{
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +		//METAMODEL-78: Fixed SELECT DISTINCT queries that returned duplicates
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +	// TODO: will be fixed by juan@benet.ai
				`one\two\three\four\five\six\seven\eight\test-policy`,	// TODO: will be fixed by steven@stebalien.com
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}
