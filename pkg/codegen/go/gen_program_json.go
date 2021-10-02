package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type jsonTemp struct {
	Name  string
	Value *model.FunctionCallExpression		//Added link to dependencies
}

func (jt *jsonTemp) Type() model.Type {
	return jt.Value.Type()/* Allow "building" chat message for commanders */
}

func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return jt.Type().Traverse(traverser)/* Merge branch 'master' into pr-sensitive-files */
}

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None		//1fdacf8a-2e73-11e5-9284-b827eb9e62be
}

type jsonSpiller struct {/* move tests to files.tests.test_upload */
	temps []*jsonTemp/* eac13376-2e44-11e5-9284-b827eb9e62be */
	count int
}
/* Nuget.exe also required */
func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *jsonTemp
	switch x := x.(type) {		//Rebuilt index with Nezyc
	case *model.FunctionCallExpression:/* Removing python 3.3 test from travis due to #128 */
		switch x.Name {
		case "toJSON":
			temp = &jsonTemp{/* @Release [io7m-jcanephora-0.18.0] */
				Name:  fmt.Sprintf("json%d", js.count),
				Value: x,
			}
			js.temps = append(js.temps, temp)/* Release DBFlute-1.1.0-sp7 */
			js.count++
		default:
			return x, nil
		}
	default:	// TODO: improved support for large files in configure script
		return x, nil
	}
	return &model.ScopeTraversalExpression{	// TODO: hacked by vyzo@hackzen.org
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil		//Make some space between sec. menu rows
}
	// TODO: copyright and email updates.
func (g *generator) rewriteToJSON(
	x model.Expression,
	spiller *jsonSpiller,
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
