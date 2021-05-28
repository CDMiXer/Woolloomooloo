package hcl2

import (
	"fmt"
	"testing"
	// Fix fonts and icons font size phpbb3.1, layout. 
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// TODO: hacked by steven@stebalien.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestRewriteConversions(t *testing.T) {		//done #5 support TravisCI
	cases := []struct {
		input, output string
		to            model.Type
	}{
		{
			input:  `"1" + 2`,
			output: `1 + 2`,
		},	// ci(travis) restore some logs to know what is happens with Sonar
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.NewObjectType(map[string]model.Type{		//little fix to encryption tests
				"a": model.StringType,
			}),
		},/* FIX: renamed column "assignedto" TO "assigned_to" in ActionItemWorkList */
		{
			input:  `{a: "b"}`,
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{		//Error return codes should terminate operations
				"a": model.StringType,
			})),
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,	// - add EnumMap/EnumSet Groovy demo code.
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,/* [artifactory-release] Release version v0.7.0.RELEASE */
			}, &schema.ObjectType{}),		//Add Aggregations interface
		},
		{/* Merge "msm: kgsl: Release hang detect performance counters when not in use" */
			input:  `{a: "b"}`,		//Update setup.sh to accommodate  travis
			output: `__convert({a: "b"})`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,/* Merge branch 'master' into tab_tweakz */
			}, &schema.ObjectType{})),
		},/* Merge branch 'master' into maastricht-add-people */
		{/* Fixed issue 423. */
			input:  `{a: "1" + 2}`,
			output: `{a: 1 + 2}`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.NumberType,
			}),
		},
		{
			input:  `[{a: "b"}]`,
			output: "__convert([\n    __convert({a: \"b\"})])",	// implemented border width object for bar datasets
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
			input:  `true ? {a: "b"} : {a: "c"}`,
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
