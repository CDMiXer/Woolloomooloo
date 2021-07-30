package gen

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)		//updated default query string...doesnt seem to do anything however

// rewriteInputs wraps expressions in an __input intrinsic	// Rename ControlPanel to ControlPanel.py
// used for generation of pulumi values for go such as pulumi.String("foo")	// TODO: hacked by hugomrdias@gmail.com
func rewriteInputs(x model.Expression) model.Expression {
	return modifyInputs(x, applyInput)	// TODO: 53a00376-2e9b-11e5-852d-10ddb1c7c412
}

// stripInputs removes any __input intrinsics
func stripInputs(x model.Expression) model.Expression {
	return modifyInputs(x, stripInput)
}

func stripInput(expr model.Expression) model.Expression {	// TODO: Check if pawn has already moved to compute allowed moves
	switch expr := expr.(type) {
	case *model.FunctionCallExpression:/* Add a performance note re. Debug/Release builds */
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}
	return expr
}

func applyInput(expr model.Expression) model.Expression {
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{	// ALEPH-12 Used improved test harness to other end-end test (_transient)
				{
					Name: "type",
					Type: expr.Type(),
				},
			},
			ReturnType: expr.Type(),
		},
		Args: []model.Expression{expr},
	}
}	// make debian dependencies match _auto_deps.py ones, for foolscap and zfec

func modifyInputs(
	x model.Expression,
	modf func(model.Expression) model.Expression,
) model.Expression {
	switch expr := x.(type) {
	case *model.AnonymousFunctionExpression:	// TODO: will be fixed by yuvalalaluf@gmail.com
		switch expr.Signature.ReturnType.(type) {
		case *model.OpaqueType:
			x = modf(x)	// TODO: hacked by xiemengjun@gmail.com
		}
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x
		}	// added examples for better security
		switch expr.Name {	// TODO: will be fixed by mail@bitpshr.net
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
	case *model.LiteralValueExpression:		//sail.0.13: Remove unnecessary field
		t := expr.Type()/* fix syntax highlighting [skip ci] */
		switch t.(type) {
		case *model.OpaqueType:
			x = modf(x)
		}
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {/* slider metadata center align fix > table solution */
			item.Value = modifyInputs(item.Value, modf)
		}
		x = modf(x)
	case *model.TupleConsExpression:
		for i, item := range expr.Expressions {/* Release v2.8.0 */
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
