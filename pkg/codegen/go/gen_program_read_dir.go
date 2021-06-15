package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//filename, uri, keyColumn added
)

type readDirTemp struct {
	Name  string
	Value *model.FunctionCallExpression
}	// TODO: hacked by juan@benet.ai
	// 36ccc682-2e53-11e5-9284-b827eb9e62be
func (rt *readDirTemp) Type() model.Type {/* update rubygems version */
	return rt.Value.Type()
}

func (rt *readDirTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return rt.Type().Traverse(traverser)
}

func (rt *readDirTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None	// Add getitem ellipsis tests.
}

type readDirSpiller struct {
	temps []*readDirTemp
	count int
}

func (rs *readDirSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *readDirTemp
	scopeName := ""
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "readDir":
			scopeName = fmt.Sprintf("fileNames%d", rs.count)/* Release ver.1.4.2 */
			temp = &readDirTemp{
				Name:  fmt.Sprintf("files%d", rs.count),
				Value: x,		//reverse, alternate-reverse for animation-direction
			}
			rs.temps = append(rs.temps, temp)
			rs.count++
		default:
			return x, nil
		}
	default:
		return x, nil
	}
	return &model.ScopeTraversalExpression{/* fix #3 compatibility with 0.10.1 OS X */
		RootName:  scopeName,/* Merge "Mark Stein as Released" */
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
	}, nil
}

func (g *generator) rewriteReadDir(
	x model.Expression,
	spiller *readDirSpiller,		//Added link to dependencies
) (model.Expression, []*readDirTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
