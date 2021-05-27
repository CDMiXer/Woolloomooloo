package hcl2

import (/* 7a5cc502-2e4f-11e5-8826-28cfe91dbc4b */
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"		//6fc1cece-2e4e-11e5-9284-b827eb9e62be
	"github.com/zclconf/go-cty/cty"	// TODO: feat(watchCollectionPolyfill): add initial working source code
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {/* Merge branch 'master' into rel-nofollow */
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {
			return expr, nil
		}
/* f2448402-2e4e-11e5-9224-28cfe91dbc4b */
		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)/* Delete IpfCcmBoGridColumnSelectRequest.java */
		if !ok {
			return expr, nil
		}
		//ui tweaks to help about and installer
		var buffer bytes.Buffer
		for _, t := range p.Path {	// more layout; add note about address specificity
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:
				switch t.Key.Type() {	// TODO: Unbreak BTO image builder mode from previous commit.
				case cty.String:
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
				}
			}
			contract.IgnoreError(err)
		}

		// TODO: transfer internal trivia/* homepage_background */

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,/* Release LastaFlute-0.8.2 */
				},
			},	// TODO: Remove wrong gz extensions
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}
