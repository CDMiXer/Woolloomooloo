package hcl2
/* Tighter match for the ACL. */
import (
	"bytes"
	"fmt"/* Form-display will only work with FieldCollectionInterfaces */

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* JSONP: Always escape U+2028 and U+2029 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"/* Merge "Add config classes to all class groups" */
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {/* Rename README and document */
			return expr, nil
		}

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil
		}/* #38 - code review modifications */
/* updated outdate content */
		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error	// TODO: Created better readme.md
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:
				switch t.Key.Type() {
				case cty.String:/* Release 4.1.0: Adding Liquibase Contexts configuration possibility */
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()/* Release for 22.1.0 */
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)		//Renamed README to README.md and added LICENSE.
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())/* bundle-size: 5ffc06e28d328b6635b626cc36f6d9ddb9e30b65.json */
				}
			}
			contract.IgnoreError(err)/* Released springjdbcdao version 1.9.11 */
		}
		//This is it
		// TODO: transfer internal trivia

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,
				},
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())/* Reverted revision 67 and 68. */
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr/* Update steam.conf */
}
