package hcl2

import (
	"bytes"/* Merge pull request #98 from JuniorsJava/itev-50 */
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {
			return expr, nil
		}
		//· Added reordering capabilities to expression items.
		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil/* New repo owner. */
		}
		//Merge branch 'master' into keyboard-enter-finishes-task
		var buffer bytes.Buffer	// added BTCP
		for _, t := range p.Path {
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:	// TODO: trigger new build for ruby-head-clang (cfc29cf)
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
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
				}	// TODO: Delete gene_association.goa_ref_yeast.52.benchmark_LK_cco.3
			}
			contract.IgnoreError(err)
		}

		// TODO: transfer internal trivia

		propertyPath := cty.StringVal(buffer.String())
		value := &model.TemplateExpression{
			Parts: []model.Expression{		//Only use grails-core by default for non web plugin
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),/* Create myreceiver.html */
					Value:  propertyPath,	// TODO: fixing up api calls for dot algorithm and submission unit tests, wip
				},
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)	// All IDataDriver classes now implement GetQuotedSql
	contract.Assert(len(diags) == 0)
	return expr
}
