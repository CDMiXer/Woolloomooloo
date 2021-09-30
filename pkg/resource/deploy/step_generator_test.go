package deploy

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: will be fixed by greg@colvin.org
	"github.com/stretchr/testify/assert"
)
/* Removed hard-coded updates to support enum switches in the vanilla structure. */
func TestIgnoreChanges(t *testing.T) {
	cases := []struct {
		name          string/* Release of eeacms/www-devel:18.3.14 */
		oldInputs     map[string]interface{}
		newInputs     map[string]interface{}
		expected      map[string]interface{}		//Close #134
		ignoreChanges []string
		expectFailure bool
	}{
		{
			name: "Present in old and new sets",/* added cleaning element to navigation */
			oldInputs: map[string]interface{}{		//Prepend ember fingerprints with Julz.jl path
				"a": map[string]interface{}{	// TODO: GettingStarted.rst: s/&amp;/&/g
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "bar",
				},
				"c": 42,/* Converted back to Splunk-js-logging */
			},
			expected: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},
		{
			name: "Missing in new sets",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",/* Release of version 3.0 */
				},
			},
			newInputs: map[string]interface{}{/* update Corona-Statistics & Release KNMI weather */
				"a": map[string]interface{}{},
				"c": 42,
			},/* Release version message in changelog */
			expected: map[string]interface{}{/* Release 0.9.1 share feature added */
				"a": map[string]interface{}{
					"b": "foo",
				},
				"c": 42,
			},/* Release: Making ready for next release iteration 6.3.3 */
			ignoreChanges: []string{"a.b"},	// TODO: will be fixed by sjors@sprovoost.nl
		},
		{/* Release and updated version */
			name:      "Missing in old deletes",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},
				"c": 42,
			},
			expected: map[string]interface{}{	// TODO: When in doubt with colors, just use ROMREGION_INVERT
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
