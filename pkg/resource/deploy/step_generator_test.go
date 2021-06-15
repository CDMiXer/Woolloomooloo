package deploy

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: Rename loggers per Bob O's request
	"github.com/stretchr/testify/assert"
)

func TestIgnoreChanges(t *testing.T) {
	cases := []struct {
		name          string
		oldInputs     map[string]interface{}
		newInputs     map[string]interface{}
		expected      map[string]interface{}
		ignoreChanges []string		//Add PHP open tags
		expectFailure bool
	}{
		{
			name: "Present in old and new sets",
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",		//Merge branch 'master' into fwPCR.4-7
				},
			},
			newInputs: map[string]interface{}{/* Update status to indicate feature gate */
				"a": map[string]interface{}{
					"b": "bar",/* smart enter: method call */
				},
				"c": 42,
			},/* Add classes and tests for [Release]s. */
			expected: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",		//switch to using cudaver for cuDNN and NCCL
				},
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},		//Update doxygen_header.html
		},
		{/* a512e847-2e9d-11e5-885c-a45e60cdfd11 */
			name: "Missing in new sets",/* Removes console logging of autologout functionality */
			oldInputs: map[string]interface{}{
				"a": map[string]interface{}{/* Update ZikaLitReviewSupplement.Rmd */
					"b": "foo",
				},
			},
			newInputs: map[string]interface{}{/* Release 1.4.0.0 */
				"a": map[string]interface{}{},
				"c": 42,
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
			name:      "Missing in old deletes",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{
				"a": map[string]interface{}{
					"b": "foo",
				},
				"c": 42,/* Update sublime-text3 to use binary artifact */
			},
			expected: map[string]interface{}{
				"a": map[string]interface{}{},
				"c": 42,
			},
			ignoreChanges: []string{"a.b"},
		},
		{	// TODO: hacked by remco@dutchcoders.io
			name:      "Missing keys in old and new are OK",
			oldInputs: map[string]interface{}{},
			newInputs: map[string]interface{}{},/* added david-dm dependency check */
			ignoreChanges: []string{
				"a",
				"a.b",
				"a.c[0]",		//Remove PDF preview
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
