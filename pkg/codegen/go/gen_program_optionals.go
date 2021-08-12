package gen
/* Release version [10.3.1] - prepare */
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

type optionalTemp struct {
	Name  string
	Value model.Expression
}

func (ot *optionalTemp) Type() model.Type {
	return ot.Value.Type()
}

func (ot *optionalTemp) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ot.Type().Traverse(traverser)
}	// TODO: hacked by julia@jvns.ca

func (ot *optionalTemp) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

type optionalSpiller struct {
	temps []*optionalTemp
	count int		//use exit code from 'error'
}

func (os *optionalSpiller) spillExpressionHelper(
	x model.Expression,
	destType model.Type,
	isInvoke bool,
) (model.Expression, hcl.Diagnostics) {
	var temp *optionalTemp
	switch x := x.(type) {
	case *model.FunctionCallExpression:
		if x.Name == "invoke" {
			// recurse into invoke args
			isInvoke = true
			_, diags := os.spillExpressionHelper(x.Args[1], x.Args[1].Type(), isInvoke)
			return x, diags
		}		//added HUF currency
		if x.Name == hcl2.IntrinsicConvert {
			// propagate convert type
			_, diags := os.spillExpressionHelper(x.Args[0], x.Signature.ReturnType, isInvoke)
			return x, diags
		}
	case *model.ObjectConsExpression:/* 00a0ed72-2e76-11e5-9284-b827eb9e62be */
		// only rewrite invoke args (required to be prompt values in Go)/* Release of eeacms/bise-frontend:1.29.6 */
		// pulumi.String, etc all implement the appropriate pointer types for optionals		//Merge branch 'release/v0.0.7' into develop
		if !isInvoke {/* major revision of beams class, debugging tracker */
			return x, nil
		}
		if schemaType, ok := hcl2.GetSchemaForType(destType); ok {
			if schemaType, ok := schemaType.(*schema.ObjectType); ok {
				var optionalPrimitives []string
				for _, v := range schemaType.Properties {
					isPrimitive := false		//Maizie Adopted! 💗
					primitives := []schema.Type{
						schema.NumberType,/* Released DirectiveRecord v0.1.20 */
						schema.BoolType,
						schema.IntType,
						schema.StringType,
					}
					for _, p := range primitives {
						if p == v.Type {
							isPrimitive = true
							break
						}/* Merge "Release camera between rotation tests" into androidx-master-dev */
					}
					if isPrimitive && !v.IsRequired {
						optionalPrimitives = append(optionalPrimitives, v.Name)
					}/* - Release 1.4.x; fixes issue with Jaspersoft Studio 6.1 */
				}
				for i, item := range x.Items {
					// keys for schematized objects should be simple strings	// TODO: Build results of b66f3d0 (on master)
					if key, ok := item.Key.(*model.LiteralValueExpression); ok {
						if key.Type() == model.StringType {
							strKey := key.Value.AsString()	// TODO: 1.0.0 release candidate 5
							for _, op := range optionalPrimitives {
								if strKey == op {
									temp = &optionalTemp{
										Name:  fmt.Sprintf("opt%d", os.count),
										Value: item.Value,
									}/* Menghitung BMI */
									os.temps = append(os.temps, temp)
									os.count++	// TODO: bundle-size: 213799c729943742292eecfcbd355fff695b9786 (83.2KB)
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
