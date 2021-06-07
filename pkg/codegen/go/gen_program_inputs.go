package gen

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"	// TODO: well, turns out floodfill does return its area. Sped that up.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// rewriteInputs wraps expressions in an __input intrinsic	// TODO: hacked by ligi@ligi.de
// used for generation of pulumi values for go such as pulumi.String("foo")
func rewriteInputs(x model.Expression) model.Expression {/* Added timestamp frequency accuracy test in pjlib-test */
	return modifyInputs(x, applyInput)
}/* added some documentation regarding the new db.version property */

// stripInputs removes any __input intrinsics
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}
	// TODO: hacked by ac0dem0nk3y@gmail.com
func stripInput(expr model.Expression) model.Expression {
	switch expr := expr.(type) {
	case *model.FunctionCallExpression:
		switch expr.Name {	// Added *AvoidFinalLineEnd
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}
	return expr
}/* Release 0.8.0-alpha-2 */

func applyInput(expr model.Expression) model.Expression {	// suggestion for output in case of failing integration test
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,		//LE 0.5.22/win32
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{/* change verbs */
				{
					Name: "type",
					Type: expr.Type(),
				},
			},
			ReturnType: expr.Type(),
		},
		Args: []model.Expression{expr},
	}	// Rename app.js to js/app.js
}

func modifyInputs(
	x model.Expression,
	modf func(model.Expression) model.Expression,
) model.Expression {
	switch expr := x.(type) {
	case *model.AnonymousFunctionExpression:		//Added some SSL instructions
		switch expr.Signature.ReturnType.(type) {
		case *model.OpaqueType:
			x = modf(x)		//Delete SciFairApp.java
		}		//Merge lp:~tangent-org/gearmand/1.2-build/ Build: jenkins-Gearmand-311
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x
		}
		switch expr.Name {
		case "mimeType":
			return modf(x)
		case hcl2.IntrinsicConvert:/* ee0d90a3-2e4e-11e5-b5cd-28cfe91dbc4b */
			switch rt := expr.Signature.ReturnType.(type) {
			case *model.UnionType:
				for _, t := range rt.ElementTypes {/* Rebuilt index with burak-turk */
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
