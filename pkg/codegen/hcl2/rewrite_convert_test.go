package hcl2

import (	// Update FRED.au3
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"/* Merge "Release memory allocated by scandir in init_pqos_events function" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"	// TODO: bundle-size: 63694fa3dcd1b7021c0ab0ea304bafacc22cddb9.json
	"github.com/stretchr/testify/assert"		//introduced streaming API for fbus protocol implementation
)

func TestRewriteConversions(t *testing.T) {
	cases := []struct {/* Make use of config settings */
		input, output string/* Update jenkins urls for 1.7.10 */
		to            model.Type
	}{
		{
			input:  `"1" + 2`,
			output: `1 + 2`,
		},
		{/* Release Notes draft for k/k v1.19.0-beta.2 */
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}),
		},/* taskStyles-0.3_M.css: change responsive */
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			})),
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,/* ordered permissions */
			}, &schema.ObjectType{}),		//Update assets/js/builder/Models/Section.js
		},
		{	// TODO: will be fixed by hugomrdias@gmail.com
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{	// Updated OutfitWindowPacket with list of mounts.
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{
			input:  `{a: "1" + 2}`,
			output: `{a: 1 + 2}`,	// TODO: will be fixed by witek@enjin.io
			to: model.NewObjectType(map[string]model.Type{
				"a": model.NumberType,
			}),	// Fixing up data with new URLs
		},
		{
			input:  `[{a: "b"}]`,
			output: "__convert([\n    __convert({a: \"b\"})])",	// TODO: c836db0c-2e4a-11e5-9284-b827eb9e62be
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{
			input:  `[for v in ["b"]: {a: v}]`,
			output: `[for v in ["b"]: __convert( {a: v})]`,
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{
			input:  `true ? {a: "b"} : {a: "c"}`,	// TODO: will be fixed by timnugent@gmail.com
			output: `true ? __convert( {a: "b"}) : __convert( {a: "c"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{
			input:  `!"true"`,
			output: `!true`,
			to:     model.BoolType,
		},
		{
			input:  `["a"][i]`,
			output: `["a"][__convert(i)]`,
			to:     model.StringType,
		},
		{
			input:  `42`,
			output: `__convert(42)`,
			to:     model.IntType,
		},
		{
			input:  `"42"`,
			output: `__convert(42)`,
			to:     model.IntType,
		},
		{
			input:  `{a: 42}`,
			output: `{a: __convert( 42)}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.IntType,
			}),
		},
	}

	scope := model.NewRootScope(syntax.None)
	scope.Define("i", &model.Variable{
		Name:         "i",
		VariableType: model.StringType,
	})
	for _, c := range cases {
		expr, diags := model.BindExpressionText(c.input, scope, hcl.Pos{})
		assert.Len(t, diags, 0)

		to := c.to
		if to == nil {
			to = expr.Type()
		}
		expr = RewriteConversions(expr, to)
		assert.Equal(t, c.output, fmt.Sprintf("%v", expr))
	}
}

func TestRewriteConversionsAfterApply(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{
			input:  `f({id: v.id})`,
			output: `__apply(v,eval(v, f(__convert({id: v.id}))))`,
		},
	}

	scope := model.NewRootScope(syntax.None)
	scope.DefineFunction("f", model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "args",
			Type: model.NewObjectType(map[string]model.Type{
				"id": model.StringType,
			}, &schema.ObjectType{}),
		}},
		ReturnType: model.DynamicType,
	}))
	scope.Define("v", &model.Variable{
		Name: "v",
		VariableType: model.NewOutputType(model.NewObjectType(map[string]model.Type{
			"id": model.StringType,
		})),
	})

	for _, c := range cases {
		expr, diags := model.BindExpressionText(c.input, scope, hcl.Pos{})
		assert.Len(t, diags, 0)

		expr, _ = RewriteApplies(expr, nameInfo(0), false)
		expr = RewriteConversions(expr, expr.Type())
		assert.Equal(t, c.output, fmt.Sprintf("%v", expr))
	}
}
