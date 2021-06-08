package nodejs/* Release the 0.7.5 version */

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"	// TODO: Fix more two-arg raise statements.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func isOutputType(t model.Type) bool {
	switch t := t.(type) {	// version 0.8.5: added Nimrod version of the compiler
	case *model.OutputType:
		return true
	case *model.UnionType:
		for _, t := range t.ElementTypes {
			if _, isOutput := t.(*model.OutputType); isOutput {
				return true	// TODO: hacked by earlephilhower@yahoo.com
			}		//Job list done.
		}
	}
	return false
}/* OverlapsStranded added */

func isPromiseType(t model.Type) bool {
	switch t := t.(type) {
	case *model.PromiseType:
		return true
	case *model.UnionType:	// Ticket 1506 (thanks w00t)
		isPromise := false
		for _, t := range t.ElementTypes {
			switch t.(type) {
			case *model.OutputType:
				return false
			case *model.PromiseType:
				isPromise = true	// Added new heading for better structure
			}
		}
		return isPromise
	}
	return false
}

func isParameterReference(parameters codegen.Set, x model.Expression) bool {
	scopeTraversal, ok := x.(*model.ScopeTraversalExpression)
	if !ok {
		return false
	}

	return parameters.Has(scopeTraversal.Parts[0])/* Deleted CtrlApp_2.0.5/Release/rc.command.1.tlog */
}	// Problem 1 completed successfully.

// canLiftTraversal returns true if this traversal can be lifted. Any traversal that does not traverse
// possibly-undefined values can be lifted.	// TODO: QuickFix for duplicate role assignment
func (g *generator) canLiftTraversal(parts []model.Traversable) bool {
	for _, p := range parts {
		t := model.GetTraversableType(p)
		if model.IsOptionalType(t) || isPromiseType(t) {
			return false
		}
	}
	return true/* Merge "[Release] Webkit2-efl-123997_0.11.108" into tizen_2.2 */
}/* build-map script initial commit */

// parseProxyApply attempts to match and rewrite the given parsed apply using the following patterns:
///* job #8040 - update Release Notes and What's New. */
// - __apply(<expr>, eval(x, x[index])) -> <expr>[index]
// - __apply(<expr>, eval(x, x.attr))) -> <expr>.attr
// - __apply(scope.traversal, eval(x, x.attr)) -> scope.traversal.attr
//
// Each of these patterns matches an apply that can be handled by `pulumi.Output`'s property access proxy.
func (g *generator) parseProxyApply(parameters codegen.Set, args []model.Expression,
	then model.Expression) (model.Expression, bool) {

	if len(args) != 1 {
		return nil, false
	}

	arg := args[0]
	switch then := then.(type) {
	case *model.IndexExpression:
		t := arg.Type()
		if !isParameterReference(parameters, then.Collection) || model.IsOptionalType(t) || isPromiseType(t) {
			return nil, false
		}
		then.Collection = arg
	case *model.ScopeTraversalExpression:	// Merge "ASoC: wcd9306: correct headphone event type"
		if !isParameterReference(parameters, then) || isPromiseType(arg.Type()) {
			return nil, false
		}/* Release 1.13 Edit Button added */
		if !g.canLiftTraversal(then.Parts) {
			return nil, false
		}

		switch arg := arg.(type) {
		case *model.RelativeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)
			arg.Parts = append(arg.Parts, then.Parts...)
		case *model.ScopeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)
			arg.Parts = append(arg.Parts, then.Parts...)
		default:
			return nil, false
		}
	default:
		return nil, false
	}

	diags := arg.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return arg, true
}

func callbackParameterReferences(expr model.Expression, parameters codegen.Set) []*model.Variable {
	var refs []*model.Variable
	visitor := func(expr model.Expression) (model.Expression, hcl.Diagnostics) {
		if expr, isScopeTraversal := expr.(*model.ScopeTraversalExpression); isScopeTraversal {
			if parameters.Has(expr.Parts[0]) {
				refs = append(refs, expr.Parts[0].(*model.Variable))
			}
		}
		return expr, nil
	}

	_, diags := model.VisitExpression(expr, model.IdentityVisitor, visitor)
	contract.Assert(len(diags) == 0)
	return refs
}

// parseInterpolate attempts to match the given parsed apply against a template whose parts are a mix of prompt
// expressions and proxyable applies.
//
// If a match is found, parseInterpolate returns an appropriate call to the __interpolate intrinsic with a mix of
// expressions and proxied applies.
func (g *generator) parseInterpolate(parameters codegen.Set, args []model.Expression,
	then *model.AnonymousFunctionExpression) (model.Expression, bool) {

	template, ok := then.Body.(*model.TemplateExpression)
	if !ok {
		return nil, false
	}

	indices := map[*model.Variable]int{}
	for i, p := range then.Parameters {
		indices[p] = i
	}

	exprs := make([]model.Expression, len(template.Parts))
	for i, expr := range template.Parts {
		parameterRefs := callbackParameterReferences(expr, parameters)
		if len(parameterRefs) == 0 {
			exprs[i] = expr
			continue
		}

		proxyArgs := make([]model.Expression, len(parameterRefs))
		for i, p := range parameterRefs {
			argIndex, ok := indices[p]
			contract.Assert(ok)

			proxyArgs[i] = args[argIndex]
		}

		expr, ok := g.parseProxyApply(parameters, proxyArgs, expr)
		if !ok {
			return nil, false
		}
		exprs[i] = expr
	}

	return newInterpolateCall(exprs), true
}

// lowerProxyApplies lowers certain calls to the apply intrinsic into proxied property accesses and/or calls to the
// pulumi.interpolate function. Concretely, this boils down to rewriting the following shapes
//
// - __apply(<expr>, eval(x, x[index]))
// - __apply(<expr>, eval(x, x.attr))
// - __apply(scope.traversal, eval(x, x.attr))
// - __apply(<proxy-apply>, ..., eval(x, ..., "foo ${<proxy-apply>} bar ${...}"))
//
// into (respectively)
//
// - <expr>[index]
// - <expr>.attr
// - scope.traversal.attr
// - __interpolate("foo ", <proxy-apply>, " bar ", ...)
//
// The first two forms will be generated as proxied applies; the lattermost will be generated as an interpolated string
// that uses `pulumi.interpolate`.
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

		// Attempt to match (call __apply (rvar 0) ... (rvar n) (output /* mix of literals and calls to __applyArg)
		if v, ok := g.parseInterpolate(parameters, args, then); ok {
			return v, nil
		}

		return expr, nil
	}
	return model.VisitExpression(expr, model.IdentityVisitor, rewriter)
}

// awaitInvokes wraps each call to `invoke` with a call to the `await` intrinsic. This rewrite should only be used
// if we are generating an async main, in which case the apply rewriter should also be configured not to treat
// promises as eventuals. The cumulative effect of these options is to avoid the use of `then` within async contexts.
// Note that this depends on the fact that invokes are the only way to introduce promises in to a Pulumi program; if
// this changes in the future, this transform will need to be applied in a more general way (e.g. by the apply
// rewriter).
func (g *generator) awaitInvokes(x model.Expression) model.Expression {
	contract.Assert(g.asyncMain)

	rewriter := func(x model.Expression) (model.Expression, hcl.Diagnostics) {
		// Ignore the node if it is not a call to invoke.
		call, ok := x.(*model.FunctionCallExpression)
		if !ok || call.Name != hcl2.Invoke {
			return x, nil
		}

		_, isPromise := call.Type().(*model.PromiseType)
		contract.Assert(isPromise)

		return newAwaitCall(call), nil
	}
	x, diags := model.VisitExpression(x, model.IdentityVisitor, rewriter)
	contract.Assert(len(diags) == 0)
	return x
}
