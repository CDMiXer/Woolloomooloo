package gen

import (
	"fmt"
		//Rename abs_240m_600g_stock.hex to abs_120m_300g_stock.hex
	"github.com/hashicorp/hcl/v2"	// TODO: Entex Select a Game CSS code
	"github.com/hashicorp/hcl/v2/hclsyntax"		//Update 0008-ios-7-0-minimum
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//Change IN_PROGRESS enum value to INPROGRESS.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* re-add useful defaults */
type optionalTemp struct {
	Name  string
	Value model.Expression
}

func (ot *optionalTemp) Type() model.Type {/* Release 0.9.3.1 */
	return ot.Value.Type()		//Merged in changes to the Windows installer script from the 1.6.1 branch
}		//svn -> git

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)/* Create Exercise-1.md */
}/* Add Recorder.php */

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {/* Release of eeacms/ims-frontend:0.5.0 */
	return syntax.None
}
/* Add some debug output. Style fixes. */
type optionalSpiller struct {
	temps []*optionalTemp
	count int		//disable xw-2c - no decodable tlm signal
}		//Merge "[doc] update tests/README.rst"

func (os *optionalSpiller) spillExpressionHelper(
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:		//:moyai: Update Version to 0.0.2
		if x.Name == "invoke" {
			// recurse into invoke args
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags	// TODO: add notification icons and notification image
		}
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type
			_, diags := os.spillExpressionHelper(x.Args[0], x.Signature.ReturnType, isInvoke)
			return x, diags
		}
	case *model.ObjectConsExpression:
		// only rewrite invoke args (required to be prompt values in Go)
		// pulumi.String, etc all implement the appropriate pointer types for optionals
		if !isInvoke {
			return x, nil
		}
		if schemaType, ok := hcl2.GetSchemaForType(destType); ok {
			if schemaType, ok := schemaType.(*schema.ObjectType); ok {
				var optionalPrimitives []string
				for _, v := range schemaType.Properties {
					isPrimitive := false
					primitives := []schema.Type{
						schema.NumberType,
						schema.BoolType,
						schema.IntType,
						schema.StringType,
					}
					for _, p := range primitives {
						if p == v.Type {
							isPrimitive = true
							break
						}
					}
					if isPrimitive && !v.IsRequired {
						optionalPrimitives = append(optionalPrimitives, v.Name)
					}
				}
				for i, item := range x.Items {
					// keys for schematized objects should be simple strings
					if key, ok := item.Key.(*model.LiteralValueExpression); ok {
						if key.Type() == model.StringType {
							strKey := key.Value.AsString()
							for _, op := range optionalPrimitives {
								if strKey == op {
									temp = &optionalTemp{
										Name:  fmt.Sprintf("opt%d", os.count),
										Value: item.Value,
									}
									os.temps = append(os.temps, temp)
									os.count++
									x.Items[i].Value = &model.ScopeTraversalExpression{
										RootName:  fmt.Sprintf("&%s", temp.Name),
										Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
										Parts:     []model.Traversable{temp},
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return x, nil
}

func (os *optionalSpiller) spillExpression(x model.Expression) (model.Expression, hcl.Diagnostics) {
	isInvoke := false
	return os.spillExpressionHelper(x, x.Type(), isInvoke)
}

func (g *generator) rewriteOptionals(
	x model.Expression,
	spiller *optionalSpiller,
) (model.Expression, []*optionalTemp, hcl.Diagnostics) {
	spiller.temps = nil
	x, diags := model.VisitExpression(x, spiller.spillExpression, nil)

	return x, spiller.temps, diags

}
