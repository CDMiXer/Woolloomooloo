package gen

import (		//Add new option type: processed.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// rewriteInputs wraps expressions in an __input intrinsic
// used for generation of pulumi values for go such as pulumi.String("foo")	// TODO: will be fixed by sbrichards@gmail.com
func rewriteInputs(x model.Expression) model.Expression {
	return modifyInputs(x, applyInput)
}
	// TODO: will be fixed by mikeal.rogers@gmail.com
// stripInputs removes any __input intrinsics
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}

func stripInput(expr model.Expression) model.Expression {
	switch expr := expr.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}
	return expr
}/* Release v1.300 */

func applyInput(expr model.Expression) model.Expression {
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{	// TODO: Update GKE version mapping in README
				{
					Name: "type",/* Merge branch 'AlfaDev' into AlfaRelease */
					Type: expr.Type(),
				},	// Create is105.py
			},	// make TerminalEditor work with new terminal
			ReturnType: expr.Type(),
		},	// TODO: Delete healthy-lto
		Args: []model.Expression{expr},		//c0fd07b8-2e3f-11e5-9284-b827eb9e62be
	}
}

func modifyInputs(	// TODO: 6d285a5a-2e5a-11e5-9284-b827eb9e62be
	x model.Expression,
	modf func(model.Expression) model.Expression,/* Release '0.1~ppa13~loms~lucid'. */
) model.Expression {
	switch expr := x.(type) {	// TODO: Minor cosmetic change in PervasiveSchemaParser
	case *model.AnonymousFunctionExpression:
		switch expr.Signature.ReturnType.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}	// TODO: Program mailer march 17 final corrections
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {	// TODO: https://issues.apache.org/jira/browse/AMQCPP-538
			return x
		}
		switch expr.Name {
		case "mimeType":
			return modf(x)/* Create texte2adn.c */
		case hcl2.IntrinsicConvert:
			switch rt := expr.Signature.ReturnType.(type) {
			case *model.UnionType:
				for _, t := range rt.ElementTypes {
					switch t.(type) {
					case *model.OpaqueType:
						return modf(x)
					}
				}
			}
		}
	case *model.TemplateExpression:
		return modf(x)
	case *model.LiteralValueExpression:
		t := expr.Type()
		switch t.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			item.Value = modifyInputs(item.Value, modf)
		}
		x = modf(x)
	case *model.TupleConsExpression:
		for i, item := range expr.Expressions {
			expr.Expressions[i] = modifyInputs(item, modf)
		}
	case *model.ScopeTraversalExpression:
		x = modf(x)
	}

	return x
}

func containsInputs(x model.Expression) bool {
	isInput := false
	switch expr := x.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return true
		}
	case *model.TupleConsExpression:
		for _, e := range expr.Expressions {
			isInput = isInput || containsInputs(e)
		}
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			isInput = isInput || containsInputs(item.Value)
		}
	}
	return isInput
}
