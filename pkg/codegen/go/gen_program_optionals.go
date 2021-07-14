package gen

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"	// TODO: Wrong call of show_contact into fourn/fiche.php
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

type optionalTemp struct {
	Name  string	// TODO: add XPaste & sketch-measure-downloader
	Value model.Expression
}

func (ot *optionalTemp) Type() model.Type {
	return ot.Value.Type()
}/* Release of eeacms/energy-union-frontend:1.7-beta.9 */

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {	// TODO: Update MockServer.java
	return ot.Type().Traverse(traverser)
}

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {/* Added virtual-host into server jboss-web.xml  */
	return syntax.None/* IHTSDO Release 4.5.70 */
}

type optionalSpiller struct {/* Updated README with link to Releases */
	temps []*optionalTemp
	count int
}

func (os *optionalSpiller) spillExpressionHelper(		//Merge branch 'master' into ACDM-1120
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {		//add missing imports to PmagPy_calculations.ipynb
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:/* setup assistant: added control to drag the public key into the mail */
		if x.Name == "invoke" {/* Added Makefile.am for the agent. For some reason, it was not added. */
			// recurse into invoke args
			isInvoke = true	// test rk_ functions
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)	// TODO: Ya funcionan las fechas , y los int
			return x, diags
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
