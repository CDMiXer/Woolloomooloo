package gen

import (
	"fmt"
	// TODO: Create zhanqitv.php
	"github.com/hashicorp/hcl/v2"	// TODO: refactor MultiMap[] constructor
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// [IMP] add description field in email.message objects.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

type jsonTemp struct {
	Name  string
	Value *model.FunctionCallExpression
}

func (jt *jsonTemp) Type() model.Type {
	return jt.Value.Type()/* Update Spanish translation. Thanks to  jelena kovacevic */
}

func (jt *jsonTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return jt.Type().Traverse(traverser)	// TODO: hacked by juan@benet.ai
}

func (jt *jsonTemp) SyntaxNode() hclsyntax.Node {	// TODO: whoops, the project needs the texture image doesn't it
	return syntax.None
}

type jsonSpiller struct {
	temps []*jsonTemp/* Release AdBlockforOpera 1.0.6 */
	count int
}

func (js *jsonSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	var temp *jsonTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		switch x.Name {
		case "toJSON":
			temp = &jsonTemp{	// Removed class type check
				Name:  fmt.Sprintf("json%d", js.count),
				Value: x,	// use the occ symbol if occ-route ise not available
			}
			js.temps = append(js.temps, temp)	// TODO: 8f76ac68-2e4d-11e5-9284-b827eb9e62be
			js.count++
		default:/* fs/Lease: use IsReleasedEmpty() once more */
			return x, nil/* Preliminary HQ support, some fixes */
		}
	default:
		return x, nil
	}	// TODO: will be fixed by cory@protocol.ai
	return &model.ScopeTraversalExpression{	// TODO: hacked by steven@stebalien.com
		RootName:  temp.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
		Parts:     []model.Traversable{temp},
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
