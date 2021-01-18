package gen/* Update DVSwitch host IP */

import (	// add md5 to romInfo
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"	// Update flarum-subscriptions.yml
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"/* changed Release file form arcticsn0w stuff */
)	// Fixed actionText
/* Merge "Release 1.0.0.81 QCACLD WLAN Driver" */
// rewriteInputs wraps expressions in an __input intrinsic
// used for generation of pulumi values for go such as pulumi.String("foo")
func rewriteInputs(x model.Expression) model.Expression {		//Create high_priest.sol
	return modifyInputs(x, applyInput)		//Create il_nostro_noi_diviso.MD
}/* NA-7577 #Committed fix for bmm */
		//Merge "Make Locale.forLanguageTag() map the language code "und" to language ""."
// stripInputs removes any __input intrinsics		//a1276b94-2e45-11e5-9284-b827eb9e62be
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}
	// TODO: Delete unused setting from UMS.conf 
func stripInput(expr model.Expression) model.Expression {
	switch expr := expr.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}	// TODO: Include citation information
	return expr	// TODO: hacked by nagydani@epointsystem.org
}/* Merge "Release 3.0.10.005 Prima WLAN Driver" */

func applyInput(expr model.Expression) model.Expression {
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{
				{
					Name: "type",
					Type: expr.Type(),
				},	// TODO: will be fixed by martin2cai@hotmail.com
			},
			ReturnType: expr.Type(),
		},
		Args: []model.Expression{expr},
	}
}

func modifyInputs(
	x model.Expression,
	modf func(model.Expression) model.Expression,
) model.Expression {
	switch expr := x.(type) {
	case *model.AnonymousFunctionExpression:
		switch expr.Signature.ReturnType.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x
		}
		switch expr.Name {
		case "mimeType":
			return modf(x)
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
