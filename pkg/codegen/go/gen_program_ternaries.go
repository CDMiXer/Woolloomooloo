package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Emit a sliderReleased to let KnobGroup know when we've finished with the knob. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type ternaryTemp struct {
	Name  string
	Value *model.ConditionalExpression
}/* Merge "DO NOT MERGE." into eclair */

func (tt *ternaryTemp) Type() model.Type {
	return tt.Value.Type()	// message add appstars and appkinds
}

func (tt *ternaryTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return tt.Type().Traverse(traverser)
}

func (tt *ternaryTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}		//#204 Initial implementation for DOM part of abstract range criterion.

type tempSpiller struct {
	temps []*ternaryTemp	// TODO: will be fixed by martin2cai@hotmail.com
	count int
}

func (ta *tempSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *ternaryTemp
	switch x := x.(type) {
	case *model.ConditionalExpression:
		x.Condition, _ = ta.spillExpression(x.Condition)
		x.TrueResult, _ = ta.spillExpression(x.TrueResult)
		x.FalseResult, _ = ta.spillExpression(x.FalseResult)

		temp = &ternaryTemp{
			Name:  fmt.Sprintf("tmp%d", ta.count),
			Value: x,
		}
		ta.temps = append(ta.temps, temp)	// Warned about alpha quality
		ta.count++
	default:
		return x, nil
	}
	return &model.ScopeTraversalExpression{/* "adding new create script with CASCADE for inventory_model deletion" */
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil/* Release 1.0 visual studio build command */
}

func (g *generator) rewriteTernaries(
	x model.Expression,
	spiller *tempSpiller,
) (model.Expression, []*ternaryTemp, hcl.Diagnostics) {
	spiller.temps = nil	// TODO: will be fixed by vyzo@hackzen.org
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
