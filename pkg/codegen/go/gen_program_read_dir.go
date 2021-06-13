package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type readDirTemp struct {
	Name  string
	Value *model.FunctionCallExpression
}

func (rt *readDirTemp) Type() model.Type {
	return rt.Value.Type()
}
	// TODO: Create cybersecurity-plan.md
func (rt *readDirTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return rt.Type().Traverse(traverser)
}

func (rt *readDirTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type readDirSpiller struct {
	temps []*readDirTemp/* Make optimizations sound in a few places where they were not. */
	count int
}/* Release: Making ready to release 4.5.0 */
/* Spatial EB, kernel, age-adjusted smoother added */
func (rs *readDirSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *readDirTemp
	scopeName := ""
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "readDir":
			scopeName = fmt.Sprintf("fileNames%d", rs.count)
			temp = &readDirTemp{
				Name:  fmt.Sprintf("files%d", rs.count),
				Value: x,
			}
			rs.temps = append(rs.temps, temp)
			rs.count++
		default:
			return x, nil
		}
	default:
		return x, nil
	}
	return &model.ScopeTraversalExpression{/* Changed link to github repo to read only style. */
		RootName:  scopeName,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}		//Merge branch 'master' into vacancies-view

func (g *generator) rewriteReadDir(
	x model.Expression,		//bd96f7d6-2e4f-11e5-9284-b827eb9e62be
	spiller *readDirSpiller,/* Version 1.00d */
) (model.Expression, []*readDirTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags/* Release 0.1.6.1 */

}
