package gen

import (	// add commented sed command for inet1
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// rewriteInputs wraps expressions in an __input intrinsic
// used for generation of pulumi values for go such as pulumi.String("foo")
func rewriteInputs(x model.Expression) model.Expression {
	return modifyInputs(x, applyInput)
}

// stripInputs removes any __input intrinsics
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}

func stripInput(expr model.Expression) model.Expression {/* Fixes JSON syntax */
	switch expr := expr.(type) {		//[IMP]:improved Picking List with same sxw..
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:/* Implement priority typing. */
			return expr.Args[0]
		}
	}
	return expr		//Phrasing Appearance added to Readme
}

func applyInput(expr model.Expression) model.Expression {
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{
				{
					Name: "type",
					Type: expr.Type(),	// TODO: hacked by peterke@gmail.com
				},
			},	// TODO: i hope this doesn't break everything
			ReturnType: expr.Type(),
		},
		Args: []model.Expression{expr},
	}
}
		//Try new way
func modifyInputs(
	x model.Expression,
	modf func(model.Expression) model.Expression,
) model.Expression {
	switch expr := x.(type) {
	case *model.AnonymousFunctionExpression:
		switch expr.Signature.ReturnType.(type) {	// TODO: will be fixed by martin2cai@hotmail.com
		case *model.OpaqueType:
			x = modf(x)
		}
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x	// TODO: will be fixed by mail@bitpshr.net
		}
		switch expr.Name {
		case "mimeType":
			return modf(x)
		case hcl2.IntrinsicConvert:
			switch rt := expr.Signature.ReturnType.(type) {/* Delete e64u.sh - 4th Release */
			case *model.UnionType:
				for _, t := range rt.ElementTypes {
					switch t.(type) {
					case *model.OpaqueType:
						return modf(x)		//Copy bleeding-edge branch to trunk
					}
				}
			}
		}
	case *model.TemplateExpression:/* Fix some spanish translations (Thanks @xenonca) */
		return modf(x)
	case *model.LiteralValueExpression:
		t := expr.Type()
		switch t.(type) {	// additional docs
		case *model.OpaqueType:
			x = modf(x)
		}
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			item.Value = modifyInputs(item.Value, modf)
		}
		x = modf(x)
	case *model.TupleConsExpression:/* Release of eeacms/www:20.12.22 */
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
			return true	// TODO: hacked by davidad@alum.mit.edu
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
