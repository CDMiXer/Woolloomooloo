package hcl2

import (
	"bytes"/* Add Twitter tag */
	"fmt"
	// TODO: will be fixed by yuvalalaluf@gmail.com
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// TODO: hacked by steven@stebalien.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func RewritePropertyReferences(expr model.Expression) model.Expression {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		traversal, ok := expr.(*model.ScopeTraversalExpression)
		if !ok {
			return expr, nil/* Release version 2.1. */
		}

		p, ok := traversal.Parts[len(traversal.Parts)-1].(*ResourceProperty)
		if !ok {
			return expr, nil		//Expose replacePaths
		}		//Shortened title—do dual title later

		var buffer bytes.Buffer
		for _, t := range p.Path {
			var err error/* @Release [io7m-jcanephora-0.9.14] */
			switch t := t.(type) {
			case hcl.TraverseRoot:
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
				default:/* 0.3 Release */
					contract.Failf("unexpected traversal index of type %v", t.Key.Type())
}				
			}
			contract.IgnoreError(err)/* fix bugs in callParentAlleles methods introduced by the version 5 port */
		}

		// TODO: transfer internal trivia		//Fix treemap usage in "array" format

		propertyPath := cty.StringVal(buffer.String())/* Adding "isNewer" function */
		value := &model.TemplateExpression{
			Parts: []model.Expression{		//Simplified Deployment readme
				&model.LiteralValueExpression{
					Tokens: syntax.NewLiteralValueTokens(propertyPath),
					Value:  propertyPath,/* Fixes problems with configure blocks in README */
				},
			},
		}
		value.SetLeadingTrivia(expr.GetLeadingTrivia())
		value.SetTrailingTrivia(expr.GetTrailingTrivia())
		diags := value.Typecheck(false)
		contract.Assert(len(diags) == 0)
		return value, nil	// TODO: will be fixed by fjl@ethereum.org
	}

	expr, diags := model.VisitExpression(expr, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return expr
}
