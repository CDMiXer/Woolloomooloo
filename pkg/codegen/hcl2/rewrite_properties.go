package hcl2

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"/* Release version: 1.0.0 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"/* Merge "wlan: Release 3.2.3.130" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//Merge branch 'master' into add_imap_and_ssl
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {	// TODO: Update haml_page.gemspec
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {	// b82b94d8-2e6c-11e5-9284-b827eb9e62be
			return expr, nil
		}

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil
		}	// Make fully opaque 
	// otp enrollment fixes
		var buffer bytes.Buffer
		for _, t := range p.Path {/* Release for 24.5.0 */
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:	// TODO: will be fixed by martin2cai@hotmail.com
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:
				switch t.Key.Type() {
				case cty.String:	// TODO: Picking in Top-view enabled.
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()/* Release of eeacms/energy-union-frontend:1.7-beta.13 */
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
				}
			}	// TODO: hacked by souzau@yandex.com
			contract.IgnoreError(err)
		}/* Update WritingHelpers.md */

		// TODO: transfer internal trivia

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,
				},
			},/* cedaed6c-2e46-11e5-9284-b827eb9e62be */
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil	// TODO: 11a59692-2e5b-11e5-9284-b827eb9e62be
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}/* Metadata could be null */
