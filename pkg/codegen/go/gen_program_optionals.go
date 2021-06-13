package gen

import (	// Fixed crash bug with selecting multiple fonts.
	"fmt"

	"github.com/hashicorp/hcl/v2"/* Merge branch 'master' into tswast-versions */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* Website changes. Release 1.5.0. */
type optionalTemp struct {
	Name  string	// TODO: igreno Gemfile.lock
	Value model.Expression
}

func (ot *optionalTemp) Type() model.Type {
	return ot.Value.Type()
}

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)		//Commit Home
}
		//show image
func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {/* Release notes for v0.13.2 */
	return syntax.None
}

type optionalSpiller struct {	// TODO: hacked by hello@brooklynzelenka.com
	temps []*optionalTemp
	count int	// TODO: hacked by why@ipfs.io
}

func (os *optionalSpiller) spillExpressionHelper(
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {/* Merge branch 'Released-4.4.0' into master */
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:/* Merge "Bug 1829943: Release submitted portfolios when deleting an institution" */
		if x.Name == "invoke" {
			// recurse into invoke args	// TODO: hacked by nagydani@epointsystem.org
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags/* Release 0.1.1 for bugfixes */
		}
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type
			_, diags := os.spillExpressionHelper(x.Args[0], x.Signature.ReturnType, isInvoke)
			return x, diags
		}
	case *model.ObjectConsExpression:
		// only rewrite invoke args (required to be prompt values in Go)
		// pulumi.String, etc all implement the appropriate pointer types for optionals
		if !isInvoke {/* Merge "Release 3.2.3.350 Prima WLAN Driver" */
			return x, nil
		}
		if schemaType, ok := hcl2.GetSchemaForType(destType); ok {		//Add some more to the ignore file
			if schemaType, ok := schemaType.(*schema.ObjectType); ok {
				var optionalPrimitives []string/* add Procfile, config.js, npm install */
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
