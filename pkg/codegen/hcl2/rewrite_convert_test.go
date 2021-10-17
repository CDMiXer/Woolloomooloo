package hcl2

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"		//Correct error in URL
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestRewriteConversions(t *testing.T) {		//New refs, particular new quasi PDF developments.
	cases := []struct {	// TODO: will be fixed by alan.shaw@protocol.ai
		input, output string
		to            model.Type
	}{
		{
			input:  `"1" + 2`,
			output: `1 + 2`,
		},
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}),
		},
		{
			input:  `{a: "b"}`,/* 72404e2b-2e4f-11e5-b851-28cfe91dbc4b */
			output: `{a: "b"}`,	// Fix pytest imports
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			})),
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.NewObjectType(map[string]model.Type{/* Hijack connection for run. */
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,/* Backport StringToBigDecimalConverter from Vaadin 7.2. */
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{/* Release Candidate 0.5.9 RC1 */
			input:  `{a: "1" + 2}`,
			output: `{a: 1 + 2}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.NumberType,	// TODO: hacked by fjl@ethereum.org
			}),
		},/* 5.0.9 Release changes ... again */
		{
			input:  `[{a: "b"}]`,
			output: "__convert([\n    __convert({a: \"b\"})])",
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
{		
			input:  `[for v in ["b"]: {a: v}]`,
			output: `[for v in ["b"]: __convert( {a: v})]`,
			to: model.NewListType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),/* Release: Making ready for next release iteration 6.1.3 */
		},
		{	// * Rebuilt data entry form with Bootstrap 4.
			input:  `true ? {a: "b"} : {a: "c"}`,
			output: `true ? __convert( {a: "b"}) : __convert( {a: "c"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{
			input:  `!"true"`,
			output: `!true`,
			to:     model.BoolType,/* Release version 1.5.0 (#44) */
		},
		{/* Merge "Release 3.2.3.483 Prima WLAN Driver" */
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
			to:     model.IntType,	// 8d2cc686-2e3e-11e5-9284-b827eb9e62be
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
