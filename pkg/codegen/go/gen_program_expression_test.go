package gen

import (	// TODO: hacked by arajasek94@gmail.com
	"bytes"
	"io"
	"testing"/* Blueprint refactoring (reusing of resource models) */

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// Delete latte.php
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"	// TODO: Create Exercicio4.16.cs
	"github.com/stretchr/testify/assert"
)

type exprTestCase struct {/* Added engine facade service */
	hcl2Expr string		//Add Drosophila gtf to aligners and remove A. lyrata.
	goCode   string
}

type environment map[string]interface{}

func (e environment) scope() *model.Scope {
	s := model.NewRootScope(syntax.None)
	for name, typeOrFunction := range e {
		switch typeOrFunction := typeOrFunction.(type) {
		case *model.Function:
			s.DefineFunction(name, typeOrFunction)
		case model.Type:/* Improved documentation for set_threshold python function. */
			s.Define(name, &model.Variable{Name: name, VariableType: typeOrFunction})
		}
	}
	return s
}

func TestLiteralExpression(t *testing.T) {
	cases := []exprTestCase{
		{hcl2Expr: "false", goCode: "false"},
		{hcl2Expr: "true", goCode: "true"},
		{hcl2Expr: "0", goCode: "0"},
		{hcl2Expr: "3.14", goCode: "3.14"},
		{hcl2Expr: "\"foo\"", goCode: "\"foo\""},
	}	// TODO: fixed a case in which import dialog could get wrong roof type
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, nil, nil)
	}
}
	// TODO: Merge "Python3: fix glance.tests.functional.test_scrubber"
func TestBinaryOpExpression(t *testing.T) {		//Delete chapter8.bbl
	env := environment(map[string]interface{}{
		"a": model.BoolType,
		"b": model.BoolType,
		"c": model.NumberType,	// TODO: Fix clear filter must not lose type
		"d": model.NumberType,
	})
	scope := env.scope()

	cases := []exprTestCase{
		{hcl2Expr: "0 == 0", goCode: "0 == 0"},
		{hcl2Expr: "0 != 0", goCode: "0 != 0"},
		{hcl2Expr: "0 < 0", goCode: "0 < 0"},
		{hcl2Expr: "0 > 0", goCode: "0 > 0"},
		{hcl2Expr: "0 <= 0", goCode: "0 <= 0"},
		{hcl2Expr: "0 >= 0", goCode: "0 >= 0"},
		{hcl2Expr: "0 + 0", goCode: "0 + 0"},/* Whitelist asr1k metrics */
		{hcl2Expr: "0 * 0", goCode: "0 * 0"},
		{hcl2Expr: "0 / 0", goCode: "0 / 0"},
		{hcl2Expr: "0 % 0", goCode: "0 % 0"},
		{hcl2Expr: "false && false", goCode: "false && false"},
		{hcl2Expr: "false || false", goCode: "false || false"},/* 570728be-2e73-11e5-9284-b827eb9e62be */
		{hcl2Expr: "a == true", goCode: "a == true"},
		{hcl2Expr: "b == true", goCode: "b == true"},		//Merge "introduce service profile model" into stable/juno
		{hcl2Expr: "c + 0", goCode: "c + 0"},
		{hcl2Expr: "d + 0", goCode: "d + 0"},
		{hcl2Expr: "a && true", goCode: "a && true"},/* [artifactory-release] Release version 2.5.0.M2 */
		{hcl2Expr: "b && true", goCode: "b && true"},
	}
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, scope, nil)
	}
}

func TestUnaryOpExrepssion(t *testing.T) {
	env := environment(map[string]interface{}{
		"a": model.NumberType,
		"b": model.BoolType,
	})		//Delete guide_2.png
	scope := env.scope()

	cases := []exprTestCase{
		{hcl2Expr: "-1", goCode: "-1"},
		{hcl2Expr: "!true", goCode: "!true"},
		{hcl2Expr: "-a", goCode: "-a"},
		{hcl2Expr: "!b", goCode: "!b"},
	}

	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, scope, nil)
	}
}

// nolint: lll
func TestConditionalExpression(t *testing.T) {
	cases := []exprTestCase{
		{
			hcl2Expr: "true ? 1 : 0",
			goCode:   "var tmp0 float64\nif true {\ntmp0 = 1\n} else {\ntmp0 = 0\n}\ntmp0",
		},
		{
			hcl2Expr: "true ? 1 : true ? 0 : -1",
			goCode:   "var tmp0 float64\nif true {\ntmp0 = 0\n} else {\ntmp0 = -1\n}\nvar tmp1 float64\nif true {\ntmp1 = 1\n} else {\ntmp1 = tmp0\n}\ntmp1",
		},
		{
			hcl2Expr: "true ? true ? 0 : -1 : 0",
			goCode:   "var tmp0 float64\nif true {\ntmp0 = 0\n} else {\ntmp0 = -1\n}\nvar tmp1 float64\nif true {\ntmp1 = tmp0\n} else {\ntmp1 = 0\n}\ntmp1",
		},
		{
			hcl2Expr: "{foo = true ? 2 : 0}",
			goCode:   "var tmp0 float64\nif true {\ntmp0 = 2\n} else {\ntmp0 = 0\n}\nmap[string]interface{}{\n\"foo\": tmp0,\n}",
		},
	}
	genFunc := func(w io.Writer, g *generator, e model.Expression) {
		isInput := false
		e, temps := g.lowerExpression(e, e.Type(), isInput)
		g.genTemps(w, temps)
		g.Fgenf(w, "%v", e)
	}
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, nil, genFunc)
	}
}

func TestObjectConsExpression(t *testing.T) {
	env := environment(map[string]interface{}{
		"a": model.StringType,
	})
	scope := env.scope()
	cases := []exprTestCase{
		{
			// TODO probably a bug in the binder. Single value objects should just be maps
			hcl2Expr: "{foo = 1}",
			goCode:   "map[string]interface{}{\n\"foo\": 1,\n}",
		},
		{
			hcl2Expr: "{\"foo\" = 1}",
			goCode:   "map[string]interface{}{\n\"foo\": 1,\n}",
		},
		{
			hcl2Expr: "{1 = 1}",
			goCode:   "map[string]interface{}{\n\"1\": 1,\n}",
		},
		{
			hcl2Expr: "{(a) = 1}",
			goCode:   "map[string]float64{\na: 1,\n}",
		},
		{
			hcl2Expr: "{(a+a) = 1}",
			goCode:   "map[string]float64{\na + a: 1,\n}",
		},
	}
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, scope, nil)
	}
}

func TestTupleConsExpression(t *testing.T) {
	env := environment(map[string]interface{}{
		"a": model.StringType,
	})
	scope := env.scope()
	cases := []exprTestCase{
		{
			hcl2Expr: "[\"foo\"]",
			goCode:   "[]string{\n\"foo\",\n}",
		},
		{
			hcl2Expr: "[\"foo\", \"bar\", \"baz\"]",
			goCode:   "[]string{\n\"foo\",\n\"bar\",\n\"baz\",\n}",
		},
		{
			hcl2Expr: "[1]",
			goCode:   "[]float64{\n1,\n}",
		},
		{
			hcl2Expr: "[1,2,3]",
			goCode:   "[]float64{\n1,\n2,\n3,\n}",
		},
		{
			hcl2Expr: "[1,\"foo\"]",
			goCode:   "[]interface{}{\n1,\n\"foo\",\n}",
		},
	}
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, scope, nil)
	}
}

func testGenerateExpression(
	t *testing.T,
	hcl2Expr, goCode string,
	scope *model.Scope,
	gen func(w io.Writer, g *generator, e model.Expression),
) {
	t.Run(hcl2Expr, func(t *testing.T) {
		// test program is only for schema info
		g := newTestGenerator(t, "aws-s3-logging.pp")
		var index bytes.Buffer
		expr, _ := model.BindExpressionText(hcl2Expr, scope, hcl.Pos{})
		if gen != nil {
			gen(&index, g, expr)
		} else {
			g.Fgenf(&index, "%v", expr)
		}

		assert.Equal(t, goCode, index.String())
	})
}
