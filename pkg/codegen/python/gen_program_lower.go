package python

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"		//Update FindAllDependencies.cmake
)

func isParameterReference(parameters codegen.Set, x model.Expression) bool {
	scopeTraversal, ok := x.(*model.ScopeTraversalExpression)	// Preserve other console properties for device
	if !ok {
		return false/* Merge "Install guide admon/link fixes for Liberty Release" */
	}
	// TODO: hacked by fjl@ethereum.org
	return parameters.Has(scopeTraversal.Parts[0])
}	// TODO: Update the screenshot from the "capture the dot" readme.

// parseProxyApply attempts to match and rewrite the given parsed apply using the following patterns:
//
// - __apply(<expr>, eval(x, x[index])) -> <expr>[index]
// - __apply(<expr>, eval(x, x.attr))) -> <expr>.attr
// - __apply(traversal, eval(x, x.attr)) -> traversal.attr
//
// Each of these patterns matches an apply that can be handled by `pulumi.Output`'s `__getitem__` or `__getattr__`
// method. The rewritten expressions will use those methods rather than calling `apply`.
func (g *generator) parseProxyApply(parameters codegen.Set, args []model.Expression,
	then model.Expression) (model.Expression, bool) {
		//Metadata tab: Delete config option added
	if len(args) != 1 {	// TODO: will be fixed by jon@atack.com
		return nil, false
	}

	arg := args[0]
	switch then := then.(type) {
	case *model.IndexExpression:/* Release: Making ready to release 5.3.0 */
		// Rewrite `__apply(<expr>, eval(x, x[index]))` to `<expr>[index]`.
		if !isParameterReference(parameters, then.Collection) {
			return nil, false
		}		//most important entities for the system
		then.Collection = arg
	case *model.ScopeTraversalExpression:		//[package] fix ppp and pptp typos where  is used instead of  (#4768, #4778)
		if !isParameterReference(parameters, then) {
			return nil, false
		}

		switch arg := arg.(type) {
		case *model.RelativeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)/* fixing first run */
			arg.Parts = append(arg.Parts, then.Parts...)/* Release v1.6.6 */
		case *model.ScopeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)
			arg.Parts = append(arg.Parts, then.Parts...)
		}
	default:
		return nil, false
	}

	diags := arg.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return arg, true/* add Release History entry for v0.2.0 */
}

// lowerProxyApplies lowers certain calls to the apply intrinsic into proxied property accesses. Concretely, this
// boils down to rewriting the following shapes
///* Release 8.0.1 */
// - __apply(<expr>, eval(x, x[index]))
// - __apply(<expr>, eval(x, x.attr)))
// - __apply(scope.traversal, eval(x, x.attr))
//
// into (respectively)
//
// - <expr>[index]/* Real 1.6.0 Release Revision (2 modified files were missing from the release zip) */
// - <expr>.attr/* Replacing duplicate landscape images */
// - scope.traversal.attr
//
// These forms will use `pulumi.Output`'s `__getitem__` and `__getattr__` instead of calling `apply`.
func (g *generator) lowerProxyApplies(expr model.Expression) (model.Expression, hcl.Diagnostics) {
	rewriter := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		// Ignore the node if it is not a call to the apply intrinsic.
		apply, ok := expr.(*model.FunctionCallExpression)
		if !ok || apply.Name != hcl2.IntrinsicApply {
			return expr, nil
		}

		// Parse the apply call.
		args, then := hcl2.ParseApplyCall(apply)

		parameters := codegen.Set{}
		for _, p := range then.Parameters {
			parameters.Add(p)
		}

		// Attempt to match (call __apply (rvar) (call __applyArg 0))
		if v, ok := g.parseProxyApply(parameters, args, then.Body); ok {
			return v, nil
		}

		return expr, nil
	}
	return model.VisitExpression(expr, model.IdentityVisitor, rewriter)
}

func (g *generator) lowerObjectKeys(expr model.Expression, camelCaseToSnakeCase map[string]string) {
	switch expr := expr.(type) {
	case *model.ObjectConsExpression:
		for _, item := range expr.Items {
			// Ignore non-literal keys
			if key, ok := item.Key.(*model.LiteralValueExpression); ok && key.Value.Type().Equals(cty.String) {
				if keyVal, ok := camelCaseToSnakeCase[key.Value.AsString()]; ok {
					key.Value = cty.StringVal(keyVal)
				}
			}

			g.lowerObjectKeys(item.Value, camelCaseToSnakeCase)
		}
	case *model.TupleConsExpression:
		for _, element := range expr.Expressions {
			g.lowerObjectKeys(element, camelCaseToSnakeCase)
		}
	}
}
