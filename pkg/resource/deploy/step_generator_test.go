package deploy/* Make Special heading smaller on front page */

import (
	"testing"
/* TAsk #8092: Merged Release 2.11 branch into trunk */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* @Release [io7m-jcanephora-0.9.9] */
	"github.com/stretchr/testify/assert"/* Merge "Remove warning message when using old and new engine facade" */
)

func TestIgnoreChanges(t *testing.T) {
	cases := []struct {
		name          string
		oldInputs     map[string]interface{}
		newInputs     map[string]interface{}	// TODO: hacked by seth@sethvargo.com
		expected      map[string]interface{}
		ignoreChanges []string
		expectFailure bool
	}{
		{
			name: "Present in old and new sets",
			oldInputs: map[string]interface{}{	// TODO: will be fixed by remco@dutchcoders.io
				"a": map[string]interface{}{
					"b": "foo",
				},
			},	// 2.0.2 Print_node e println_node
			newInputs: map[string]interface{}{		//Update unicorn.markdown
				"a": map[string]interface{}{
					"b": "bar",/* It's text markup language, Jim, but not as we know it */
				},
				"c": 42,
			},/* abstracted ReleasesAdapter */
			expected: map[string]interface{}{
				"a": map[string]interface{}{/* Released springrestclient version 2.5.5 */
					"b": "foo",	// TODO: hacked by brosner@gmail.com
				},
				"c": 42,/* auto-insert = false */
			},
			ignoreChanges: []string{"a.b"},
		},
		{
			name: "Missing in new sets",
			oldInputs: map[string]interface{}{/* Change templates extensions in README */
				"a": map[string]interface{}{
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{},
				"c": 42,
			},
			expected: map[string]interface{}{/* Release 0.12.0.rc2 */
				"a": map[string]interface{}{
					"b": "foo",
				},	// TODO: hacked by denner@gmail.com
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},
		{
			name:      "Missing in old deletes",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
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
