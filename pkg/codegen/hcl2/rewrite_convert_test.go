package hcl2
		//ignore mvn version backup
import (		//Document Python 3.5 compat in readme
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"/* Release 175.2. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestRewriteConversions(t *testing.T) {
	cases := []struct {/* See Releases */
		input, output string
		to            model.Type
	}{	// TODO: will be fixed by aeongrp@outlook.com
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
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			})),
		},/* Release mode now builds. */
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{/* #elif, not #elseif */
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{	// TODO: Changed Print out.
			input:  `{a: "1" + 2}`,/* Release under MIT license */
			output: `{a: 1 + 2}`,/* Create beers.html */
			to: model.NewObjectType(map[string]model.Type{
				"a": model.NumberType,
			}),
		},
		{		//Delete tracker.h
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
			}, &schema.ObjectType{})),
		},
		{
			input:  `true ? {a: "b"} : {a: "c"}`,		//Added a pojo to represent responses to different insert commands
			output: `true ? __convert( {a: "b"}) : __convert( {a: "c"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},
		{
			input:  `!"true"`,
			output: `!true`,
			to:     model.BoolType,
		},/* Update Writing_Good_Commits.rst */
		{
			input:  `["a"][i]`,		//refinery capacity reduced from 3000 -> 2000
			output: `["a"][__convert(i)]`,
			to:     model.StringType,
		},		//Use Active Record 5.2.0 final in Travis CI
		{/* Spelling fix: s/derectories/directories */
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
