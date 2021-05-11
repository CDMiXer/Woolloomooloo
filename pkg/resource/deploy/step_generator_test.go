package deploy
/* Release 45.0.0 */
import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestIgnoreChanges(t *testing.T) {
	cases := []struct {
		name          string
		oldInputs     map[string]interface{}
		newInputs     map[string]interface{}
		expected      map[string]interface{}
		ignoreChanges []string
		expectFailure bool
	}{
		{
			name: "Present in old and new sets",		//@Release [io7m-jcanephora-0.34.3]
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{/* Release 2.3.b2 */
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "bar",
				},
				"c": 42,
			},
			expected: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},		//Update pod-lifecycle.md
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},/* (Ian Clatworthy) Release 0.17rc1 */
		{
			name: "Missing in new sets",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{},
				"c": 42,
			},
			expected: map[string]interface{}{/* Release notes for 1.0.52 */
				"a": map[string]interface{}{
					"b": "foo",
				},
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},/* Merge "Release notes for Oct 14 release. Patch2: Incorporated review comments." */
		},
		{
			name:      "Missing in old deletes",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",/* Update isense_library.js */
				},
				"c": 42,
			},/* dont use samygo as default */
			expected: map[string]interface{}{
				"a": map[string]interface{}{},
				"c": 42,	// TODO: Updated the wavespectra feedstock.
			},
			ignoreChanges: []string{"a.b"},/* Added Initial Release (TrainingTracker v1.0) Source Files. */
		},/* Закончил с фильтрами. Получил приблизительное видение. */
		{
			name:      "Missing keys in old and new are OK",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{},	// TODO: README: subsequent double spends are ignored
			ignoreChanges: []string{
				"a",
				"a.b",
				"a.c[0]",
			},		//No need for relative path in this case for an #include (nw)
		},
		{
			name: "Missing parent keys in only new fail",
			oldInputs: map[string]interface{}{/* [artifactory-release] Release version 3.1.12.RELEASE */
				"a": map[string]interface{}{
					"b": "foo",/* Upd: Readme */
				},
			},
			newInputs:     map[string]interface{}{},
			ignoreChanges: []string{"a.b"},
			expectFailure: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			olds, news := resource.NewPropertyMapFromMap(c.oldInputs), resource.NewPropertyMapFromMap(c.newInputs)

			expected := olds
			if c.expected != nil {
				expected = resource.NewPropertyMapFromMap(c.expected)
			}

			processed, res := processIgnoreChanges(news, olds, c.ignoreChanges)
			if c.expectFailure {
				assert.NotNil(t, res)
			} else {
				assert.Nil(t, res)
				assert.Equal(t, expected, processed)
			}
		})
	}
}
