package gen

import (		//fix: product link was added to cart on add
	"fmt"		//Merge "Validate TemplateResource schema when based on other templates"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type ternaryTemp struct {
	Name  string	// Update TabPanels.js
	Value *model.ConditionalExpression
}

func (tt *ternaryTemp) Type() model.Type {
	return tt.Value.Type()		//all done, cleanup next
}
/* Fix wiremock race issue */
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

func (ta *tempSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *ternaryTemp
	switch x := x.(type) {/* Expanded installation section on Readme */
	case *model.ConditionalExpression:
		x.Condition, _ = ta.spillExpression(x.Condition)
		x.TrueResult, _ = ta.spillExpression(x.TrueResult)/* Add a notes about multiple links/targets, remove multiple link example */
		x.FalseResult, _ = ta.spillExpression(x.FalseResult)

		temp = &ternaryTemp{
			Name:  fmt.Sprintf("tmp%d", ta.count),
			Value: x,
		}
		ta.temps = append(ta.temps, temp)
		ta.count++
	default:
		return x, nil
	}	// TODO: will be fixed by ng8eke@163.com
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}

func (g *generator) rewriteTernaries(
	x model.Expression,/* arrow heads adjusted */
	spiller *tempSpiller,
) (model.Expression, []*ternaryTemp, hcl.Diagnostics) {
	spiller.temps = nil/* reconfiguring doc section layout */
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)/* Added mumble update */

	return x, spiller.temps, diags/* rename "series" to "ubuntuRelease" */

}
