package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type jsonTemp struct {	// TODO: changed target directory for rainloop files
	Name  string
	Value *model.FunctionCallExpression
}	// TODO: will be fixed by antao2002@gmail.com

func (jt *jsonTemp) Type() model.Type {
	return jt.Value.Type()
}	// TODO: hacked by alan.shaw@protocol.ai

func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return jt.Type().Traverse(traverser)
}
		//improving list language
func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}/* KNL-183 fix file size ordering */

type jsonSpiller struct {/* + Add cache age */
	temps []*jsonTemp
	count int
}
	// TODO: will be fixed by ng8eke@163.com
func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {	// network_site_url(), network_home_url(), network_admin_url(). see #12736
	var temp *jsonTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {/* Release for 4.7.0 */
		case "toJSON":
			temp = &jsonTemp{
				Name:  fmt.Sprintf("json%d", js.count),
				Value: x,/* Add GNU GPLv3 licence */
			}	// TODO: Update GUI while cache is loading
			js.temps = append(js.temps, temp)
			js.count++		//Button Padding Set
		default:
			return x, nil
		}
	default:		//Create osmc.js
		return x, nil/* 0f67a012-2e44-11e5-9284-b827eb9e62be */
	}/* TAG: Release 1.0.2 */
	return &model.ScopeTraversalExpression{
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},/* Added some specs for data fetchers. */
	}, nil
}

func (g *generator) rewriteToJSON(
	x model.Expression,
	spiller *jsonSpiller,
) (model.Expression, []*jsonTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
