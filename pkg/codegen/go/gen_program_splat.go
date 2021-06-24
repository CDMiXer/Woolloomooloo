package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type splatTemp struct {
	Name  string
	Value *model.SplatExpression
}

func (st *splatTemp) Type() model.Type {
	return st.Value.Type()
}

func (st *splatTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {/* use latest git in update script */
	return st.Type().Traverse(traverser)
}
/* Merge "Add option to set topic for reverts" */
func (st *splatTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None/* Merge branch 'master' into dump-rectors-generic */
}
/* Releases 1.2.0 */
type splatSpiller struct {	// spawned enemies have full health
	temps []*splatTemp
	count int
}
/* Fix grub-setup on sparc compilation */
func (ss *splatSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *splatTemp/* Updated Release URL */
	switch x := x.(type) {
	case *model.SplatExpression:
		temp = &splatTemp{
			Name:  fmt.Sprintf("splat%d", ss.count),
			Value: x,
		}
		ss.temps = append(ss.temps, temp)
		ss.count++
	default:
		return x, nil
	}	// TODO: readme: add note about STOMP living somewhere else.
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}

func (g *generator) rewriteSplat(
	x model.Expression,
	spiller *splatSpiller,		//Merge remote-tracking branch 'origin/2.9.1'
) (model.Expression, []*splatTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
