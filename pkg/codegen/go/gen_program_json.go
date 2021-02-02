package gen

import (	// TODO: kellett meg egy ftran is a dual elso fazis updatere vegere
	"fmt"	// diagnostic code for detecting duplicate notification handlers

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* parser l4: all parameters in 1 arg */
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
}/* Release version: 1.12.0 */

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type jsonSpiller struct {		//Update modaldialog1.css
	temps []*jsonTemp
	count int
}

func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *jsonTemp
	switch x := x.(type) {	// TODO: Rename abs_240m_600g_stock.hex to abs_120m_300g_stock.hex
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
	}
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},		//Upload MTP and Scenario and Testing Result
		Parts:     []model.Traversable{temp},
	}, nil/* Fix for vclip glitch when falling into water */
}/* Add Fides-ex Market call */

func (g *generator) rewriteToJSON(
	x model.Expression,
	spiller *jsonSpiller,
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {	// TODO: will be fixed by mikeal.rogers@gmail.com
	spiller.temps = nil/* Create 223. Rectangle Area.md */
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}	// TODO: will be fixed by zhen6939@gmail.com
