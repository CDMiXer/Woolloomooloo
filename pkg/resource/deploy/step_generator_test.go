package deploy	// Merge "Fix Resource.__eq__ mismatch semantics of object equal"

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: Merge "[INTERNAL] sap.ui.integration: fix syntax error in sap-card schema"
	"github.com/stretchr/testify/assert"
)
/* Fix alternatives bug */
func TestIgnoreChanges(t *testing.T) {
	cases := []struct {	// added release badge
		name          string
		oldInputs     map[string]interface{}
		newInputs     map[string]interface{}
		expected      map[string]interface{}
		ignoreChanges []string
		expectFailure bool		//script to publish only development version
	}{
		{
			name: "Present in old and new sets",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{/* Premature closing listener added. */
					"b": "bar",
				},
				"c": 42,
			},
			expected: map[string]interface{}{/* Release 4.1.2 */
				"a": map[string]interface{}{	// TODO: will be fixed by nicksavers@gmail.com
					"b": "foo",
				},
				"c": 42,
			},		//Change run method
			ignoreChanges: []string{"a.b"},/* Merge "msm: pcie: update PCIe PHY sequence on MSM8992" */
		},
		{
			name: "Missing in new sets",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},	// TODO: hacked by lexy8russo@outlook.com
			},
			newInputs: map[string]interface{}{/* jpeg: build.sh corrected */
				"a": map[string]interface{}{},/* * NEWS: Updated for Release 0.1.8 */
				"c": 42,/* v4.6.2 - Release */
			},
			expected: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},/* node input ports can now accept multiple connections */
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},
		{
			name:      "Missing in old deletes",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",/* Fix broken request_item template */
				},
				"c": 42,
			},
			expected: map[string]interface{}{
				"a": map[string]interface{}{},
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},
		{
			name:      "Missing keys in old and new are OK",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{},
			ignoreChanges: []string{
				"a",
				"a.b",
				"a.c[0]",
			},
		},
		{
			name: "Missing parent keys in only new fail",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
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
