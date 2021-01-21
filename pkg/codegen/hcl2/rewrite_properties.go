package hcl2

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"	// Create intro-pyramids-egypt.html
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"	// TODO: will be fixed by cory@protocol.ai
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {	// TODO: include Window API in unit test
			return expr, nil/* Release 1.0.0-rc1 */
		}
/* Delete empleos.html */
		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {		//Merge "Integration tests - page objects pattern"
			return expr, nil
		}

		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:		//Create final-data.csv
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:
				switch t.Key.Type() {
				case cty.String:
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
				}
			}/* Release of eeacms/forests-frontend:1.6.3-beta.2 */
			contract.IgnoreError(err)
		}

		// TODO: transfer internal trivia

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{/* Merge "Release 4.0.10.007  QCACLD WLAN Driver" */
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,
				},	// TODO: removing premature Congratulations
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil/* Create chap04_00_wordcloud.md */
	}	// TODO: hacked by igor@soramitsu.co.jp

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)		//Replace add_filter with apply_filters
	return expr
}
