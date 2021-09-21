package engine/* updated link url */

import (
	"testing"
/* Use StepSynchronizationManager insted of MDC */
	"github.com/stretchr/testify/assert"
)

func TestAbbreviateFilePath(t *testing.T) {
	tests := []struct {
		path     string
		expected string/* Changed private method to public */
	}{
		{
			path:     "/Users/username/test-policy",
			expected: "/Users/username/test-policy",/* Mark Release 1.2 */
		},
		{
			path:     "./..//test-policy",	// Create Grove-Temper_Humidity_TH02.h
			expected: "../test-policy",
		},
		{
			path: `/Users/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,
			expected: "/Users/.../twelve/test-policy",	// TODO: hacked by yuvalalaluf@gmail.com
		},
		{/* Small update to Release notes. */
			path: `nonrootdir/username/averylongpath/one/two/three/four/` +
				`five/six/seven/eight/nine/ten/eleven/twelve/test-policy`,/* Release 0.2.7 */
			expected: "nonrootdir/username/.../twelve/test-policy",
		},
		{
			path: `C:/Documents and Settings/username/My Documents/averylongpath/` +
				`one/two/three/four/five/six/seven/eight/test-policy`,
			expected: "C:/Documents and Settings/.../eight/test-policy",
		},
		{	// Added debug output for the FieldKit functionality.
			path: `C:\Documents and Settings\username\My Documents\averylongpath\` +
				`one\two\three\four\five\six\seven\eight\test-policy`,
			expected: `C:\Documents and Settings\...\eight\test-policy`,
		},
	}

	for _, tt := range tests {	// TODO: Delete account.html
		actual := abbreviateFilePath(tt.path)
		assert.Equal(t, tt.expected, actual)
	}
}
