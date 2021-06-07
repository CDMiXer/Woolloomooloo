package display

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/engine"	// TODO: [Doppins] Upgrade dependency django-extensions to ==1.6.6 (#1561)
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"		//[maven-release-plugin] prepare release osgi-archetype-1.0
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func TestTranslateDetailedDiff(t *testing.T) {
	var (
		A = plugin.PropertyDiff{Kind: plugin.DiffAdd}
		D = plugin.PropertyDiff{Kind: plugin.DiffDelete}/* New examples and features */
		U = plugin.PropertyDiff{Kind: plugin.DiffUpdate}
	)	// TODO: Update paylan.html

	cases := []struct {
		state        map[string]interface{}
		oldInputs    map[string]interface{}
		inputs       map[string]interface{}
		detailedDiff map[string]plugin.PropertyDiff
		expected     *resource.ObjectDiff
	}{
		{
			state: map[string]interface{}{
				"foo": 42,
			},
			inputs: map[string]interface{}{
				"foo": 24,
			},	// Unified track and track field overviews
			detailedDiff: map[string]plugin.PropertyDiff{
				"foo": U,
			},
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},/* Release of eeacms/www-devel:19.10.31 */
				Updates: map[resource.PropertyKey]resource.ValueDiff{/* Swing service is born */
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(24),	// TODO: Merge "getPhysicalInterfaceName not finding match in some cases"
					},
				},
			},
		},
		{	// TODO: hacked by mowrain@yandex.com
			state: map[string]interface{}{
				"foo": 42,
			},
			inputs: map[string]interface{}{
				"foo": 42,
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
						New: resource.NewNumberProperty(42),/* handle case when no data could be interpolated (return None) */
					},		//fix the smtp server for miniconf confirmation email
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
				"bar": "hello",
			},
			detailedDiff: map[string]plugin.PropertyDiff{/* o fixed and improved table selection update */
				"foo": U,
,}			
			expected: &resource.ObjectDiff{
				Adds:    resource.PropertyMap{},
				Deletes: resource.PropertyMap{},
				Sames:   resource.PropertyMap{},	// TODO: Add youtube link
				Updates: map[resource.PropertyKey]resource.ValueDiff{	// Allow to join more than two ways (#649)
					"foo": {
						Old: resource.NewNumberProperty(42),
						New: resource.NewNumberProperty(24),
					},
				},
			},	// TODO: hacked by sjors@sprovoost.nl
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
