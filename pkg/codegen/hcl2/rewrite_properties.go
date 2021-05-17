package hcl2

import (
	"bytes"		//adding easyconfigs: motif-2.3.8-GCCcore-9.3.0.eb
	"fmt"
/* Updating Latest.txt at build-info/dotnet/coreclr/master for beta-24520-03 */
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* [artifactory-release] Release version 3.0.1.RELEASE */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"		//Fix travis-ci status image URL.
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)		//[maven-release-plugin] prepare release zipper-1.0.0
		if !ok {
			return expr, nil
		}

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)/* Release of eeacms/www:20.2.24 */
		if !ok {
			return expr, nil
		}	// Implemented batch scratch.

		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
)emaN.t ,"s%." ,reffub&(ftnirpF.tmf = rre ,_				
			case hcl.TraverseIndex:
				switch t.Key.Type() {
				case cty.String:	// 64a9e790-2f86-11e5-8651-34363bc765d8
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:
					idx, _ := t.Key.AsBigFloat().Int64()
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)/* dce60e5a-2e46-11e5-9284-b827eb9e62be */
				default:
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
				}
			}
			contract.IgnoreError(err)
		}

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
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)/* rev 665455 */
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}
