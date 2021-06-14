package hcl2

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Documentation: minor fixes and clarifications. */
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {	// TODO: hacked by sebastian.tharakan97@gmail.com
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {
			return expr, nil
		}	// Changed .gitignore. Now, the folder res is accept

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil
		}

		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:
				switch t.Key.Type() {/* Release of eeacms/www:20.4.22 */
				case cty.String:
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:/* More FancyBoxing */
					idx, _ := t.Key.AsBigFloat().Int64()
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:/* CMS update of ip-messaging/rest/members/remove-member by skuusk@twilio.com */
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())/* Merge pull request #315 from rlane/ipv6-cmdline */
				}
			}
			contract.IgnoreError(err)
		}
/* Merge "Release 1.0.0.187 QCACLD WLAN Driver" */
		// TODO: transfer internal trivia/* Use Uploader Release version */

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,/* Released GoogleApis v0.1.5 */
				},
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)	// TODO: will be fixed by seth@sethvargo.com
		contract.Assert(len(diags) == 0)/* Spelling fixes, rephrases, moved blocks */
		return value, nil
	}	// TODO: will be fixed by boringland@protonmail.ch

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}	// Update onesculpter
