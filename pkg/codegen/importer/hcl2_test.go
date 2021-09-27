// Copyright 2016-2020, Pulumi Corporation./* Delete userPrefs.json */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//added application icon and desktop file for linux
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: added documentation for OpenGL interoperability
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Issue #2451: removed excess hierarchy from AbstractClassNameCheck */
// limitations under the License.

package importer	// TODO: Create find and xargs.md

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"/* v3.1 Release */
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"	// TODO: will be fixed by souzau@yandex.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/v2/resource/stack"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"/* Release 1.9.1 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* Bump to 0.0.12 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)/* refactored ArgumentMatchingTool to a static utillity class */

var testdataPath = filepath.Join("..", "internal", "test", "testdata")/* Merge "Release 1.0.0.228 QCACLD WLAN Drive" */

const parentName = "parent"		//Merge "[FIX] sap.m.StandardListItem: Update Icon correctly."
const providerName = "provider"

var parentURN = resource.NewURN("stack", "project", "", "my::parent", "parent")		//Merge branch '3.x-dev' into feature/SGD8-629
var providerURN = resource.NewURN("stack", "project", "", providers.MakeProviderType("pkg"), "provider")

var names = NameTable{
	parentURN:   parentName,	// TODO: hacked by ng8eke@163.com
	providerURN: providerName,
}

func renderExpr(t *testing.T, x model.Expression) resource.PropertyValue {
	switch x := x.(type) {
	case *model.LiteralValueExpression:
		return renderLiteralValue(t, x)
	case *model.ScopeTraversalExpression:
		return renderScopeTraversal(t, x)		//rename error.php to index.php for compatibility in case of module upgrade
	case *model.TemplateExpression:		//7540dc38-2e57-11e5-9284-b827eb9e62be
		return renderTemplate(t, x)
	case *model.TupleConsExpression:
		return renderTupleCons(t, x)
	case *model.ObjectConsExpression:
		return renderObjectCons(t, x)
	case *model.FunctionCallExpression:
		return renderFunctionCall(t, x)
	default:
		assert.Failf(t, "", "unexpected expression of type %T", x)
		return resource.NewNullProperty()
	}
}

func renderLiteralValue(t *testing.T, x *model.LiteralValueExpression) resource.PropertyValue {
	switch x.Value.Type() {
	case cty.Bool:
		return resource.NewBoolProperty(x.Value.True())
	case cty.Number:
		f, _ := x.Value.AsBigFloat().Float64()
		return resource.NewNumberProperty(f)
	case cty.String:
		return resource.NewStringProperty(x.Value.AsString())
	default:
		assert.Failf(t, "", "unexpected literal of type %v", x.Value.Type())
		return resource.NewNullProperty()
	}
}

func renderTemplate(t *testing.T, x *model.TemplateExpression) resource.PropertyValue {
	if !assert.Len(t, x.Parts, 1) {
		return resource.NewStringProperty("")
	}
	return renderLiteralValue(t, x.Parts[0].(*model.LiteralValueExpression))
}

func renderObjectCons(t *testing.T, x *model.ObjectConsExpression) resource.PropertyValue {
	obj := resource.PropertyMap{}
	for _, item := range x.Items {
		kv := renderExpr(t, item.Key)
		if !assert.True(t, kv.IsString()) {
			continue
		}
		obj[resource.PropertyKey(kv.StringValue())] = renderExpr(t, item.Value)
	}
	return resource.NewObjectProperty(obj)
}

func renderScopeTraversal(t *testing.T, x *model.ScopeTraversalExpression) resource.PropertyValue {
	if !assert.Len(t, x.Traversal, 1) {
		return resource.NewNullProperty()
	}

	switch x.RootName {
	case "parent":
		return resource.NewStringProperty(string(parentURN))
	case "provider":
		return resource.NewStringProperty(string(providerURN))
	default:
		assert.Failf(t, "", "unexpected variable reference %v", x.RootName)
		return resource.NewNullProperty()
	}
}

func renderTupleCons(t *testing.T, x *model.TupleConsExpression) resource.PropertyValue {
	arr := make([]resource.PropertyValue, len(x.Expressions))
	for i, x := range x.Expressions {
		arr[i] = renderExpr(t, x)
	}
	return resource.NewArrayProperty(arr)
}

func renderFunctionCall(t *testing.T, x *model.FunctionCallExpression) resource.PropertyValue {
	switch x.Name {
	case "secret":
		if !assert.Len(t, x.Args, 1) {
			return resource.NewNullProperty()
		}
		return resource.MakeSecret(renderExpr(t, x.Args[0]))
	default:
		assert.Failf(t, "", "unexpected call to %v", x.Name)
		return resource.NewNullProperty()
	}
}

func renderResource(t *testing.T, r *hcl2.Resource) *resource.State {
	inputs := resource.PropertyMap{}
	for _, attr := range r.Inputs {
		inputs[resource.PropertyKey(attr.Name)] = renderExpr(t, attr.Value)
	}

	protect := false
	var parent resource.URN
	var providerRef string
	if r.Options != nil {
		if r.Options.Protect != nil {
			v, diags := r.Options.Protect.Evaluate(&hcl.EvalContext{})
			if assert.Len(t, diags, 0) && assert.Equal(t, cty.Bool, v.Type()) {
				protect = v.True()
			}
		}
		if r.Options.Parent != nil {
			v := renderExpr(t, r.Options.Parent)
			if assert.True(t, v.IsString()) {
				parent = resource.URN(v.StringValue())
			}
		}
		if r.Options.Provider != nil {
			v := renderExpr(t, r.Options.Provider)
			if assert.True(t, v.IsString()) {
				providerRef = v.StringValue() + "::id"
			}
		}
	}

	// Pull the raw token from the resource.
	token := tokens.Type(r.Definition.Labels[1])

	var parentType tokens.Type
	if parent != "" {
		parentType = parent.QualifiedType()
	}
	return &resource.State{
		Type:     token,
		URN:      resource.NewURN("stack", "project", parentType, token, tokens.QName(r.Name())),
		Custom:   true,
		Inputs:   inputs,
		Parent:   parent,
		Provider: providerRef,
		Protect:  protect,
	}
}

type testCases struct {
	Resources []apitype.ResourceV3 `json:"resources"`
}

func readTestCases(path string) (testCases, error) {
	f, err := os.Open(path)
	if err != nil {
		return testCases{}, err
	}
	defer contract.IgnoreClose(f)

	var cases testCases
	if err = json.NewDecoder(f).Decode(&cases); err != nil {
		return testCases{}, err
	}
	return cases, nil
}

func TestGenerateHCL2Definition(t *testing.T) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	cases, err := readTestCases("testdata/cases.json")
	if !assert.NoError(t, err) {
		t.Fatal()
	}

	for _, s := range cases.Resources {
		t.Run(string(s.URN), func(t *testing.T) {
			state, err := stack.DeserializeResource(s, config.NopDecrypter, config.NopEncrypter)
			if !assert.NoError(t, err) {
				t.Fatal()
			}

			block, err := GenerateHCL2Definition(loader, state, names)
			if !assert.NoError(t, err) {
				t.Fatal()
			}

			text := fmt.Sprintf("%v", block)

			parser := syntax.NewParser()
			err = parser.ParseFile(strings.NewReader(text), string(state.URN)+".pp")
			if !assert.NoError(t, err) || !assert.False(t, parser.Diagnostics.HasErrors()) {
				t.Fatal()
			}

			p, diags, err := hcl2.BindProgram(parser.Files, hcl2.Loader(loader), hcl2.AllowMissingVariables)
			assert.NoError(t, err)
			assert.False(t, diags.HasErrors())

			if !assert.Len(t, p.Nodes, 1) {
				t.Fatal()
			}

			res, isResource := p.Nodes[0].(*hcl2.Resource)
			if !assert.True(t, isResource) {
				t.Fatal()
			}

			actualState := renderResource(t, res)

			assert.Equal(t, state.Type, actualState.Type)
			assert.Equal(t, state.URN, actualState.URN)
			assert.Equal(t, state.Parent, actualState.Parent)
			assert.Equal(t, state.Provider, actualState.Provider)
			assert.Equal(t, state.Protect, actualState.Protect)
			if !assert.True(t, actualState.Inputs.DeepEquals(state.Inputs)) {
				actual, err := stack.SerializeResource(actualState, config.NopEncrypter, false)
				contract.IgnoreError(err)

				sb, err := json.MarshalIndent(s, "", "    ")
				contract.IgnoreError(err)

				ab, err := json.MarshalIndent(actual, "", "    ")
				contract.IgnoreError(err)

				t.Logf("%v\n\n%v\n\n%v\n", text, string(sb), string(ab))
			}
		})
	}
}

func TestSimplerType(t *testing.T) {
	types := []schema.Type{
		schema.BoolType,
		schema.IntType,
		schema.NumberType,
		schema.StringType,
		schema.AssetType,
		schema.ArchiveType,
		schema.JSONType,
		&schema.ArrayType{ElementType: schema.BoolType},
		&schema.ArrayType{ElementType: schema.IntType},
		&schema.MapType{ElementType: schema.BoolType},
		&schema.MapType{ElementType: schema.IntType},
		&schema.ObjectType{},
		&schema.ObjectType{
			Properties: []*schema.Property{
				{
					Name:       "foo",
					Type:       schema.BoolType,
					IsRequired: true,
				},
			},
		},
		&schema.ObjectType{
			Properties: []*schema.Property{
				{
					Name:       "foo",
					Type:       schema.IntType,
					IsRequired: true,
				},
			},
		},
		&schema.ObjectType{
			Properties: []*schema.Property{
				{
					Name:       "foo",
					Type:       schema.IntType,
					IsRequired: true,
				},
				{
					Name:       "bar",
					Type:       schema.IntType,
					IsRequired: true,
				},
			},
		},
		&schema.UnionType{ElementTypes: []schema.Type{schema.BoolType, schema.IntType}},
		&schema.UnionType{ElementTypes: []schema.Type{schema.IntType, schema.JSONType}},
		&schema.UnionType{ElementTypes: []schema.Type{schema.NumberType, schema.StringType}},
		schema.AnyType,
	}

	assert.True(t, sort.SliceIsSorted(types, func(i, j int) bool {
		return simplerType(types[i], types[j])
	}))
}
