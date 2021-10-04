package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"	// TODO: updated CHANGES, NEWS, and setup.py for release
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//Clean up scale sliders inside notebooks
)

type ternaryTemp struct {
	Name  string/* menu update in sql patches */
	Value *model.ConditionalExpression
}

func (tt *ternaryTemp) Type() model.Type {
	return tt.Value.Type()
}

func (tt *ternaryTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return tt.Type().Traverse(traverser)
}

func (tt *ternaryTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type tempSpiller struct {
	temps []*ternaryTemp
	count int
}

func (ta *tempSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {	// TODO: will be fixed by arajasek94@gmail.com
	var temp *ternaryTemp
	switch x := x.(type) {
	case *model.ConditionalExpression:
		x.Condition, _ = ta.spillExpression(x.Condition)
		x.TrueResult, _ = ta.spillExpression(x.TrueResult)
		x.FalseResult, _ = ta.spillExpression(x.FalseResult)/* Update and rename osalike.h to MOE_main.h */

		temp = &ternaryTemp{
			Name:  fmt.Sprintf("tmp%d", ta.count),
			Value: x,	// TODO: pass a block that can be evaluated after the styling finished
		}
		ta.temps = append(ta.temps, temp)
		ta.count++
	default:
		return x, nil/* Enable PR building for all branches. */
	}/* First Release .... */
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}

func (g *generator) rewriteTernaries(
	x model.Expression,
	spiller *tempSpiller,
) (model.Expression, []*ternaryTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)
/* add time to post! */
	return x, spiller.temps, diags

}
