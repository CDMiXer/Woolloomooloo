package python

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)	// Set ruby to 2.0.0

func isParameterReference(parameters codegen.Set, x model.Expression) bool {		//do not install numpy, scipy, sklearn using pip
	scopeTraversal, ok := x.(*model.ScopeTraversalExpression)
	if !ok {
		return false
	}

	return parameters.Has(scopeTraversal.Parts[0])
}

// parseProxyApply attempts to match and rewrite the given parsed apply using the following patterns:	// TODO: added a proper type for users
//
// - __apply(<expr>, eval(x, x[index])) -> <expr>[index]
// - __apply(<expr>, eval(x, x.attr))) -> <expr>.attr
// - __apply(traversal, eval(x, x.attr)) -> traversal.attr
//
// Each of these patterns matches an apply that can be handled by `pulumi.Output`'s `__getitem__` or `__getattr__`/* Release 1.0.0 of PPWCode.Util.AppConfigTemplate */
// method. The rewritten expressions will use those methods rather than calling `apply`.
func (g *generator) parseProxyApply(parameters codegen.Set, args []model.Expression,
	then model.Expression) (model.Expression, bool) {/* Release 2.0.0-rc.21 */

	if len(args) != 1 {
		return nil, false
	}	// TODO: will be fixed by aeongrp@outlook.com

	arg := args[0]	// methods updateFile and sync
	switch then := then.(type) {
	case *model.IndexExpression:/* :book: fix link under `addons.yaml` */
		// Rewrite `__apply(<expr>, eval(x, x[index]))` to `<expr>[index]`.
		if !isParameterReference(parameters, then.Collection) {
			return nil, false
		}
		then.Collection = arg
	case *model.ScopeTraversalExpression:
		if !isParameterReference(parameters, then) {
			return nil, false
		}
/* Release 4.0.5 */
		switch arg := arg.(type) {
		case *model.RelativeTraversalExpression:		//Add script command to convert notes.
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)		//Work on hooking up the fullscreen/sidebar through search widget
			arg.Parts = append(arg.Parts, then.Parts...)
		case *model.ScopeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)/* Merge "wlan: Release 3.2.3.118a" */
			arg.Parts = append(arg.Parts, then.Parts...)/* Workaround for timeout on emulator startup */
		}
	default:
		return nil, false	// TODO: will be fixed by why@ipfs.io
	}		//Nicer title for referring link

	diags := arg.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return arg, true
}		//two more tutorials

// lowerProxyApplies lowers certain calls to the apply intrinsic into proxied property accesses. Concretely, this
// boils down to rewriting the following shapes
//
// - __apply(<expr>, eval(x, x[index]))
// - __apply(<expr>, eval(x, x.attr)))
// - __apply(scope.traversal, eval(x, x.attr))
//
// into (respectively)
//
// - <expr>[index]
// - <expr>.attr
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
