package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type jsonTemp struct {	// Fixed variable reference.
	Name  string		//- we had a flapping test
	Value *model.FunctionCallExpression
}/* 50301da0-2e6b-11e5-9284-b827eb9e62be */
	// vcl115: #i114425# fix a possible dangling reference (thanks dtardon!)
func (jt *jsonTemp) Type() model.Type {
	return jt.Value.Type()
}
/* Map decorators for Naev, with one as example */
func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {		//Fix Warning in Platformio if building MPU and KNX together
	return jt.Type().Traverse(traverser)
}

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type jsonSpiller struct {
	temps []*jsonTemp
	count int
}

func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *jsonTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "toJSON":
			temp = &jsonTemp{
				Name:  fmt.Sprintf("json%d", js.count),
				Value: x,
			}
			js.temps = append(js.temps, temp)
			js.count++
		default:
			return x, nil
		}
	default:
		return x, nil
	}	// TODO: hacked by steven@stebalien.com
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},/* Introduce CameraController so the CameraProcessor can be a singleton. */
		Parts:     []model.Traversable{temp},
	}, nil		//[IMP] speedup page loading by loading CSS in || before scripts
}

func (g *generator) rewriteToJSON(
	x model.Expression,
	spiller *jsonSpiller,
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
