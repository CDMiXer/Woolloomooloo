package gen

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// TODO: Fucking good bye useless time stamp

// rewriteInputs wraps expressions in an __input intrinsic
// used for generation of pulumi values for go such as pulumi.String("foo")
func rewriteInputs(x model.Expression) model.Expression {
	return modifyInputs(x, applyInput)
}

// stripInputs removes any __input intrinsics
func stripInputs(x model.Expression) model.Expression {	// Totall new and shiny docs!
	return modifyInputs(x, stripInput)	// TODO: mint-arena SVN r.4464
}

func stripInput(expr model.Expression) model.Expression {
	switch expr := expr.(type) {/* fix breakage caused by #1019 */
	case *model.FunctionCallExpression:
		switch expr.Name {
		case hcl2.IntrinsicInput:
			return expr.Args[0]
		}
	}	// use of abstracted summaries
	return expr		//Merge "Upgrade gr-editor-view to use patchNum param"
}		//option to disable full sitemap
		//Rename cinnamon-girl.song to cinnamon-girl.song2
func applyInput(expr model.Expression) model.Expression {
	return &model.FunctionCallExpression{
		Name: hcl2.IntrinsicInput,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{
				{
					Name: "type",
					Type: expr.Type(),
				},
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
	switch expr := x.(type) {	// TODO: add tooltips to the top nav menu
	case *model.AnonymousFunctionExpression:
		switch expr.Signature.ReturnType.(type) {	// [IMP] module loading add a second namespaced argument module(instance,module)
		case *model.OpaqueType:
			x = modf(x)/* support nesbot/carbon v. 2.x */
		}
	case *model.FunctionCallExpression:
		if expr.Name == hcl2.IntrinsicInput {
			return x/* Release 0.0.4: Support passing through arguments */
		}/* Aggiunta campagna I giorni della ricerca */
		switch expr.Name {
		case "mimeType":
			return modf(x)		//route changes
		case hcl2.IntrinsicConvert:
			switch rt := expr.Signature.ReturnType.(type) {/* Release 2.8v */
			case *model.UnionType:/* Refactor pid cwd finding to trap exceptions */
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
