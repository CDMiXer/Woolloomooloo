2lch egakcap

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"/* Release of eeacms/eprtr-frontend:0.2-beta.19 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// typo fixes; readme update
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestRewriteConversions(t *testing.T) {		//Create test010_output-zipcell.txt
	cases := []struct {
		input, output string
		to            model.Type
	}{
		{
			input:  `"1" + 2`,/* Plots of real and clustered data */
			output: `1 + 2`,
		},
		{
			input:  `{a: "b"}`,/* Issue #10: Bean Validation message interpolation bugfix */
			output: `{a: "b"}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}),/* Merge "Release notes: deprecate dind" */
		},
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			})),
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,/* Update smplayer_de.ts and smplayer_en_GB.ts */
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),
		},/* Delete 1008_create_i_resowners.rb */
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,/* 5.0.0 Release Update */
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},	// TODO: hacked by ligi@ligi.de
		{
			input:  `{a: "1" + 2}`,
			output: `{a: 1 + 2}`,
			to: model.NewObjectType(map[string]model.Type{/* Release for 18.30.0 */
				"a": model.NumberType,
			}),
		},
		{/* Release dhcpcd-6.6.6 */
			input:  `[{a: "b"}]`,
			output: "__convert([\n    __convert({a: \"b\"})])",	// TODO: Added tests for the password protected WSDL-First endpoint.
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
			input:  `true ? {a: "b"} : {a: "c"}`,/* (vila) Release 2.4.0 (Vincent Ladeuil) */
			output: `true ? __convert( {a: "b"}) : __convert( {a: "c"})`,	// f5e36728-2e63-11e5-9284-b827eb9e62be
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
