package hcl2
/* Release: Making ready for next release iteration 6.1.3 */
import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"/* Merge "Adding extend share support in IBM GPFS Driver" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// + Added options.js for options.xul
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)
/* Create cache_tiering */
func TestRewriteConversions(t *testing.T) {
	cases := []struct {
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
			input:  `{a: "b"}`,/* create Branch DDB-524 */
			output: `{a: "b"}`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,	// TODO: hacked by earlephilhower@yahoo.com
			})),		//Only endturn message should advance game state (for now!).
		},
		{
			input:  `{a: "b"}`,/* Developer Guide is a more appropriate title than Release Notes. */
			output: `__convert({a: "b"})`,
			to: model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{}),/* Merge "Release DrmManagerClient resources" */
		},
		{
			input:  `{a: "b"}`,
			output: `__convert({a: "b"})`,
			to: model.InputType(model.NewObjectType(map[string]model.Type{
				"a": model.StringType,
			}, &schema.ObjectType{})),
		},
		{
			input:  `{a: "1" + 2}`,
			output: `{a: 1 + 2}`,
			to: model.NewObjectType(map[string]model.Type{/* Create rspec-model-testing.md */
				"a": model.NumberType,
			}),
		},
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
			}, &schema.ObjectType{})),		//compiler.cfg.dce: new global dead code elimination pass
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
			to:     model.StringType,	// TODO: file: upgraded to 4.16 (4.13 is not fetchable)
		},
		{
			input:  `42`,
			output: `__convert(42)`,
			to:     model.IntType,
		},
		{
			input:  `"42"`,
			output: `__convert(42)`,/* Upgrade PostgreSQL JDBC driver version and use JDK7 compatible driver */
			to:     model.IntType,
		},
		{
			input:  `{a: 42}`,
			output: `{a: __convert( 42)}`,		//remove out-dated doc test in inventory
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
		assert.Len(t, diags, 0)	// TODO: Should be (id + salt + timestamp) in the update
/* #4 [Release] Add folder release with new release file to project. */
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
