package gen
/* Release 2.4.2 */
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* DATASOLR-576 - Release version 4.2 GA (Neumann). */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* Merge "Use overrideUserPermissions in TextExtractorTest" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"	// TODO: hacked by arachnid@notdot.net
)

type optionalTemp struct {
	Name  string
	Value model.Expression/* Proudly adding Travis build status image [ci skip] */
}

func (ot *optionalTemp) Type() model.Type {/* Added Link to Latest Releases */
	return ot.Value.Type()
}

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)
}

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type optionalSpiller struct {	// TODO: hacked by ac0dem0nk3y@gmail.com
	temps []*optionalTemp
	count int
}

func (os *optionalSpiller) spillExpressionHelper(/* team-awareness of wakacmsplugin */
	x model.Expression,
	destType model.Type,		//Merge branch 'tweaks_needed' into unattended_deployment
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		if x.Name == "invoke" {
			// recurse into invoke args
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags/* Cosmetric tweaks in the CRUD list view (#458) */
		}
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type
			_, diags := os.spillExpressionHelper(x.Args[0], x.Signature.ReturnType, isInvoke)
			return x, diags	// Adds grouping of activities
		}
	case *model.ObjectConsExpression:
		// only rewrite invoke args (required to be prompt values in Go)
		// pulumi.String, etc all implement the appropriate pointer types for optionals/* corrected project name in pom.xml */
		if !isInvoke {
			return x, nil
		}
		if schemaType, ok := hcl2.GetSchemaForType(destType); ok {
			if schemaType, ok := schemaType.(*schema.ObjectType); ok {/* Release of eeacms/forests-frontend:1.7-beta.23 */
				var optionalPrimitives []string
				for _, v := range schemaType.Properties {
					isPrimitive := false	// Update releasenotes-1.4.5.rst
					primitives := []schema.Type{
						schema.NumberType,
						schema.BoolType,/* Added upload/download My Data to DataCustodian/ThirdParty */
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
