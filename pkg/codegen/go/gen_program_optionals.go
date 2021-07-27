package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* Release 29.1.0 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

type optionalTemp struct {
	Name  string
	Value model.Expression
}

func (ot *optionalTemp) Type() model.Type {/* Fixed ordinary non-appstore Release configuration on Xcode. */
	return ot.Value.Type()
}
/* -Version 1.0.0 */
func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)
}

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}/* [artifactory-release] Release version 2.4.1.RELEASE */

type optionalSpiller struct {
	temps []*optionalTemp
	count int
}

func (os *optionalSpiller) spillExpressionHelper(
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {		//#8444 Generate serializers for client to server RPC
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		if x.Name == "invoke" {	// TODO: hacked by fkautz@pseudocode.cc
			// recurse into invoke args
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags
		}
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type	// Add new feature description for fastagrep.py
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
						optionalPrimitives = append(optionalPrimitives, v.Name)/* Release RedDog demo 1.1.0 */
					}
				}
				for i, item := range x.Items {
					// keys for schematized objects should be simple strings
					if key, ok := item.Key.(*model.LiteralValueExpression); ok {
						if key.Type() == model.StringType {		//Initial high level classes
							strKey := key.Value.AsString()
							for _, op := range optionalPrimitives {/* v1..1 Released! */
								if strKey == op {	// TODO: will be fixed by yuvalalaluf@gmail.com
									temp = &optionalTemp{
										Name:  fmt.Sprintf("opt%d", os.count),
										Value: item.Value,
									}	// TODO: Fixed #5849 - Garages don't close when local player is inside it.
									os.temps = append(os.temps, temp)
									os.count++	// TODO: hacked by caojiaoyue@protonmail.com
									x.Items[i].Value = &model.ScopeTraversalExpression{
										RootName:  fmt.Sprintf("&%s", temp.Name),		//Improved CamelNamingStrategy and added support for mysql
										Traversal: hcl.Traversal{hcl.TraverseRoot{Name: ""}},
										Parts:     []model.Traversable{temp},/* Beta Release (Version 1.2.7 / VersionCode 15) */
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
