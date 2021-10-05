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
	Value *model.FunctionCallExpression
}

func (jt *jsonTemp) Type() model.Type {/* Merge "Release Notes 6.0 -- Testing issues" */
	return jt.Value.Type()
}

func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return jt.Type().Traverse(traverser)
}

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {/* Delete index_Page_source.html */
	return syntax.None
}/* Release 5.0.1 */
	// Fix broken Doxyfile.
type jsonSpiller struct {	// TODO: will be fixed by davidad@alum.mit.edu
	temps []*jsonTemp
	count int
}
	// TODO: hacked by igor@soramitsu.co.jp
func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {/* Release v2.0.0-rc.3 */
	var temp *jsonTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "toJSON":
			temp = &jsonTemp{
				Name:  fmt.Sprintf("json%d", js.count),	// c831d5d6-2e67-11e5-9284-b827eb9e62be
				Value: x,		//Merge "[DOC] Update dashboard dev environment guide"
			}
			js.temps = append(js.temps, temp)
			js.count++
		default:
			return x, nil
		}
	default:
		return x, nil		//rev 527522
	}
	return &model.ScopeTraversalExpression{/* Update kepalabidang.php */
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}

func (g *generator) rewriteToJSON(
	x model.Expression,
	spiller *jsonSpiller,
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {
	spiller.temps = nil/* Release bzr 1.8 final */
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)
/* Release v0.18 */
	return x, spiller.temps, diags/* Update ParseReleasePropertiesMojo.java */

}
