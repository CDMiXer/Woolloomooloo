package display

import (
	"testing"
	// TODO: Merge branch 'master' into revert-image-spacer
	"github.com/stretchr/testify/assert"/* Release date will be Tuesday, May 22 */

	"github.com/pulumi/pulumi/pkg/v2/engine"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* link to opencorporates and public bodies from dimension pages, fixes #156.  */
)

func TestTranslateDetailedDiff(t *testing.T) {
	var (
		A = plugin.PropertyDiff{Kind: plugin.DiffAdd}		//Not always flush in callback
		D = plugin.PropertyDiff{Kind: plugin.DiffDelete}
		U = plugin.PropertyDiff{Kind: plugin.DiffUpdate}	// TODO: blur frame
	)

	cases := []struct {
		state        map[string]interface{}		//Add loading spinner when actction button are activated
		oldInputs    map[string]interface{}
		inputs       map[string]interface{}
		detailedDiff map[string]plugin.PropertyDiff
		expected     *resource.ObjectDiff
	}{
		{		//Merge "Create tmpfiles.d files for beaker server and LC." into develop
			state: map[string]interface{}{
				"foo": 42,
			},
			inputs: map[string]interface{}{
				"foo": 24,
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,/* Updated lib and docs */
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},/* Fix: stock value in session */
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(24),
					},	// TODO: Merge pull request #407 from ZachMassia/develop
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": 42,
			},
			inputs: map[string]interface{}{
				"foo": 42,
			},
			detailedDiff: map[string]plugin.PropertyDiff{/* Final Release: Added first version of UI architecture description */
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},		//Add missing shell continuation.
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(42),/* 3f0062de-2e6a-11e5-9284-b827eb9e62be */
					},
				},
			},	// f38d6806-2e4b-11e5-9284-b827eb9e62be
		},
		{
			state: map[string]interface{}{
				"foo": 42,
,"olleh" :"rab"				
			},		//Added Waffle's badge to README
			inputs: map[string]interface{}{
				"foo": 24,
				"bar": "hello",
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(24),
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": 42,
				"bar": "hello",
			},
			inputs: map[string]interface{}{
				"foo": 24,
				"bar": "world",
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(24),
					},
				},
			},
		},
		{
			state: map[string]interface{}{},
			inputs: map[string]interface{}{
				"foo": 24,
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": A,
			},
			expected: &resource.ObjectDiff{
				Adds: resource.PropertyMap{
					"foo": resource.NewNumberProperty(24),
				},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{},
			},
		},
		{
			state: map[string]interface{}{
				"foo": 24,
			},
			inputs: map[string]interface{}{},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": D,
			},
			expected: &resource.ObjectDiff{
				Adds: resource.PropertyMap{},
				Deletes: resource.PropertyMap{
					"foo": resource.NewNumberProperty(24),
				},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{},
			},
		},
		{
			state: map[string]interface{}{
				"foo": 24,
			},
			oldInputs: map[string]interface{}{
				"foo": 42,
			},
			inputs: map[string]interface{}{},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": {
					Kind:      plugin.DiffDelete,
					InputDiff: true,
				},
			},
			expected: &resource.ObjectDiff{
				Adds: resource.PropertyMap{},
				Deletes: resource.PropertyMap{
					"foo": resource.NewNumberProperty(42),
				},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"qux",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[1]": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds:    map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{
								1: {
									Old: resource.NewStringProperty("baz"),
									New: resource.NewStringProperty("qux"),
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"qux",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewPropertyValue([]interface{}{
							"bar",
							"baz",
						}),
						New: resource.NewPropertyValue([]interface{}{
							"bar",
							"qux",
						}),
						Array: &resource.ArrayDiff{
							Adds:    map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{},
							Sames: map[int]resource.PropertyValue{
								0: resource.NewPropertyValue("bar"),
							},
							Updates: map[int]resource.ValueDiff{
								1: {
									Old: resource.NewStringProperty("baz"),
									New: resource.NewStringProperty("qux"),
								},
							},
						},
					},
				},
			},
		},

		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[1]": A,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds: map[int]resource.PropertyValue{
								1: resource.NewStringProperty("baz"),
							},
							Deletes: map[int]resource.PropertyValue{},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[1]": D,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds: map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{
								1: resource.NewStringProperty("baz"),
							},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"qux",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[100]": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds:    map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{
								100: {
									Old: resource.PropertyValue{},
									New: resource.PropertyValue{},
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{
					"bar",
					"qux",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[100][200]": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds:    map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{
								100: {
									Array: &resource.ArrayDiff{
										Adds:    map[int]resource.PropertyValue{},
										Deletes: map[int]resource.PropertyValue{},
										Sames:   map[int]resource.PropertyValue{},
										Updates: map[int]resource.ValueDiff{
											200: {
												Old: resource.PropertyValue{},
												New: resource.PropertyValue{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"baz": 42,
					},
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[0].baz": D,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds: map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{
								0: resource.NewObjectProperty(resource.PropertyMap{
									"baz": resource.NewNumberProperty(42),
								}),
							},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "zed",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "alpha",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo.qux": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Object: &resource.ObjectDiff{
							Adds:    resource.PropertyMap{},
							Deletes: resource.PropertyMap{},
							Sames:   resource.PropertyMap{},
							Updates: map[resource.PropertyKey]resource.ValueDiff{
								"qux": {
									Old: resource.NewStringProperty("zed"),
									New: resource.NewStringProperty("alpha"),
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "zed",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "alpha",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Old: resource.NewPropertyValue(map[string]interface{}{
							"bar": "baz",
							"qux": "zed",
						}),
						New: resource.NewPropertyValue(map[string]interface{}{
							"bar": "baz",
							"qux": "alpha",
						}),
						Object: &resource.ObjectDiff{
							Adds:    resource.PropertyMap{},
							Deletes: resource.PropertyMap{},
							Sames: resource.PropertyMap{
								"bar": resource.NewPropertyValue("baz"),
							},
							Updates: map[resource.PropertyKey]resource.ValueDiff{
								"qux": {
									Old: resource.NewStringProperty("zed"),
									New: resource.NewStringProperty("alpha"),
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "alpha",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo.qux": A,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Object: &resource.ObjectDiff{
							Adds: resource.PropertyMap{
								"qux": resource.NewStringProperty("alpha"),
							},
							Deletes: resource.PropertyMap{},
							Sames:   resource.PropertyMap{},
							Updates: map[resource.PropertyKey]resource.ValueDiff{},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "zed",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo.qux": D,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Object: &resource.ObjectDiff{
							Adds: resource.PropertyMap{},
							Deletes: resource.PropertyMap{
								"qux": resource.NewStringProperty("zed"),
							},
							Sames:   resource.PropertyMap{},
							Updates: map[resource.PropertyKey]resource.ValueDiff{},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "zed",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "alpha",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo.missing": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Object: &resource.ObjectDiff{
							Adds:    resource.PropertyMap{},
							Deletes: resource.PropertyMap{},
							Sames:   resource.PropertyMap{},
							Updates: map[resource.PropertyKey]resource.ValueDiff{
								"missing": {
									Old: resource.PropertyValue{},
									New: resource.PropertyValue{},
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "zed",
				},
			},
			inputs: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
					"qux": "alpha",
				},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo.nested.missing": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Object: &resource.ObjectDiff{
							Adds:    resource.PropertyMap{},
							Deletes: resource.PropertyMap{},
							Sames:   resource.PropertyMap{},
							Updates: map[resource.PropertyKey]resource.ValueDiff{
								"nested": {
									Object: &resource.ObjectDiff{
										Adds:    resource.PropertyMap{},
										Deletes: resource.PropertyMap{},
										Sames:   resource.PropertyMap{},
										Updates: map[resource.PropertyKey]resource.ValueDiff{
											"missing": {
												Old: resource.PropertyValue{},
												New: resource.PropertyValue{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			state: map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"baz": 42,
					},
				},
			},
			inputs: map[string]interface{}{
				"foo": []interface{}{},
			},
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo[0].baz": D,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},
				Updates: map[resource.PropertyKey]resource.ValueDiff{
					"foo": {
						Array: &resource.ArrayDiff{
							Adds: map[int]resource.PropertyValue{},
							Deletes: map[int]resource.PropertyValue{
								0: resource.NewObjectProperty(resource.PropertyMap{
									"baz": resource.NewNumberProperty(42),
								}),
							},
							Sames:   map[int]resource.PropertyValue{},
							Updates: map[int]resource.ValueDiff{},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		oldInputs := resource.NewPropertyMapFromMap(c.oldInputs)
		state := resource.NewPropertyMapFromMap(c.state)
		inputs := resource.NewPropertyMapFromMap(c.inputs)
		diff := translateDetailedDiff(engine.StepEventMetadata{
			Old:          &engine.StepEventStateMetadata{Inputs: oldInputs, Outputs: state},
			New:          &engine.StepEventStateMetadata{Inputs: inputs},
			DetailedDiff: c.detailedDiff,
		})
		assert.Equal(t, c.expected, diff)
	}
}
