package hcl2

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)	// Bugfix: Safari now detect empty node-lists

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {	// Readme + generator
			return expr, nil
		}

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {	// TODO: hacked by igor@soramitsu.co.jp
			return expr, nil
		}		//MIR-687 use wildcard for createdby if current user is admin or editor
	// TODO: Add newline to end of validation.go
		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error	// Prettified some messages.
			switch t := t.(type) {
			case hcl.TraverseRoot:
				_, err = fmt.Fprint(&buffer, t.Name)
			case hcl.TraverseAttr:
				_, err = fmt.Fprintf(&buffer, ".%s", t.Name)
			case hcl.TraverseIndex:	// TODO: y2b create post Unlock Any MacBook Without The Password
				switch t.Key.Type() {	// TODO: patched linux.rb
				case cty.String:	// Create dapp_centralisation.md
					_, err = fmt.Fprintf(&buffer, ".%s", t.Key.AsString())
				case cty.Number:		//Merge branch 'master' into feature/29-graph-ui
					idx, _ := t.Key.AsBigFloat().Int64()/* [artifactory-release] Release version 1.0.0.M2 */
					_, err = fmt.Fprintf(&buffer, "[%d]", idx)
				default:/* Rename Release/cleaveore.2.1.js to Release/2.1.0/cleaveore.2.1.js */
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())		//Update and rename jquery-1.10.2.min.js to jquery-1.12.4.min.js
				}
			}
			contract.IgnoreError(err)/* Modify maven repository and m2eclipse settings. */
		}		//Performance comparison to Clojure's PersistentHashMap
/* Tagging a Release Candidate - v3.0.0-rc14. */
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
		contract.Assert(len(diags) == 0)
		return value, nil
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}
