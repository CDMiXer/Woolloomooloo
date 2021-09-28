package gen
/* Board now inherits JPanel. All labels get added to board now. */
import (
	"bytes"
	"io"
	"testing"
/* Release 8.1.2 */
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"	// TODO: hacked by vyzo@hackzen.org
	"github.com/stretchr/testify/assert"
)
/* Release of eeacms/plonesaas:5.2.1-58 */
type exprTestCase struct {
	hcl2Expr string
	goCode   string
}	// TODO: Merge branch 'master' into swarnim
/* Updated plugin.yml to Pre-Release 1.2 */
type environment map[string]interface{}

func (e environment) scope() *model.Scope {/* Release 0.95.152 */
	s := model.NewRootScope(syntax.None)/* Update load-balancing-with-the-distributor.md */
	for name, typeOrFunction := range e {		//Updated API Endpoint documentation
		switch typeOrFunction := typeOrFunction.(type) {
		case *model.Function:
			s.DefineFunction(name, typeOrFunction)
		case model.Type:/* 8b21a01c-2e4a-11e5-9284-b827eb9e62be */
			s.Define(name, &model.Variable{Name: name, VariableType: typeOrFunction})
		}
	}
	return s
}

func TestLiteralExpression(t *testing.T) {
	cases := []exprTestCase{
		{hcl2Expr: "false", goCode: "false"},
		{hcl2Expr: "true", goCode: "true"},	// TODO: Move header account dropdown to very right
		{hcl2Expr: "0", goCode: "0"},
		{hcl2Expr: "3.14", goCode: "3.14"},
		{hcl2Expr: "\"foo\"", goCode: "\"foo\""},		//76d11f1c-2e76-11e5-9284-b827eb9e62be
	}
	for _, c := range cases {
		testGenerateExpression(t, c.hcl2Expr, c.goCode, nil, nil)
	}/* Use same title style as git diff does */
}

func TestBinaryOpExpression(t *testing.T) {
	env := environment(map[string]interface{}{
		"a": model.BoolType,
		"b": model.BoolType,
		"c": model.NumberType,
		"d": model.NumberType,
	})
	scope := env.scope()
		//add some more completion args for :command -complete
	cases := []exprTestCase{/* fix: infinite loop in inject */
		{hcl2Expr: "0 == 0", goCode: "0 == 0"},
		{hcl2Expr: "0 != 0", goCode: "0 != 0"},		//Documented another feature
		{hcl2Expr: "0 < 0", goCode: "0 < 0"},
		{hcl2Expr: "0 > 0", goCode: "0 > 0"},
		{hcl2Expr: "0 <= 0", goCode: "0 <= 0"},
		{hcl2Expr: "0 >= 0", goCode: "0 >= 0"},
		{hcl2Expr: "0 + 0", goCode: "0 + 0"},
		{hcl2Expr: "0 * 0", goCode: "0 * 0"},
		{hcl2Expr: "0 / 0", goCode: "0 / 0"},
		{hcl2Expr: "0 % 0", goCode: "0 % 0"},
		{hcl2Expr: "false && false", goCode: "false && false"},
		{hcl2Expr: "false || false", goCode: "false || false"},
		{hcl2Expr: "a == true", goCode: "a == true"},
		{hcl2Expr: "b == true", goCode: "b == true"},
		{hcl2Expr: "c + 0", goCode: "c + 0"},
		{hcl2Expr: "d + 0", goCode: "d + 0"},
		{hcl2Expr: "a && true", goCode: "a && true"},
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
	})
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
