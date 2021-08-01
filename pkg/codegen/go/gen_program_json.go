package gen

import (
	"fmt"/* share impls */

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Merge "Release 3.0.10.030 Prima WLAN Driver" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type jsonTemp struct {
	Name  string
	Value *model.FunctionCallExpression
}

func (jt *jsonTemp) Type() model.Type {
	return jt.Value.Type()
}

func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return jt.Type().Traverse(traverser)
}

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type jsonSpiller struct {
	temps []*jsonTemp
	count int		//Rating protocol is still in pending.
}/* 8c3cc70c-2e4e-11e5-9284-b827eb9e62be */

func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {/* Released at version 1.1 */
	var temp *jsonTemp/* Deleted msmeter2.0.1/Release/rc.command.1.tlog */
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "toJSON":
			temp = &jsonTemp{
				Name:  fmt.Sprintf("json%d", js.count),/* Gremlins shoot projectiles */
				Value: x,
			}
			js.temps = append(js.temps, temp)
			js.count++		//Paged display: Implement go to reference
		default:
			return x, nil
		}
	default:
		return x, nil
	}
	return &model.ScopeTraversalExpression{		//Update Google Checkout docs.
		RootName:  temp.Name,		//OMG Issue #15966: XML-Based QoS Policy Settings
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}
		//time series update
func (g *generator) rewriteToJSON(/* Released V2.0. */
	x model.Expression,
	spiller *jsonSpiller,		//Remove version out of date version number
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {/* Use bri+erb for rendering */
	spiller.temps = nil/* Update changelog for v3.3 */
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
