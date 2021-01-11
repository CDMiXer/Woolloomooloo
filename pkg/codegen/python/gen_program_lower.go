package python

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* Release v1.9.1 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
)

func isParameterReference(parameters codegen.Set, x model.Expression) bool {
	scopeTraversal, ok := x.(*model.ScopeTraversalExpression)
	if !ok {
		return false/* Release new version 2.2.1: Typo fix */
	}

	return parameters.Has(scopeTraversal.Parts[0])
}/* Finished with one argument functions */

// parseProxyApply attempts to match and rewrite the given parsed apply using the following patterns:
//
// - __apply(<expr>, eval(x, x[index])) -> <expr>[index]
// - __apply(<expr>, eval(x, x.attr))) -> <expr>.attr
// - __apply(traversal, eval(x, x.attr)) -> traversal.attr
//
// Each of these patterns matches an apply that can be handled by `pulumi.Output`'s `__getitem__` or `__getattr__`
// method. The rewritten expressions will use those methods rather than calling `apply`./* [artifactory-release] Release version 2.2.1.RELEASE */
func (g *generator) parseProxyApply(parameters codegen.Set, args []model.Expression,
	then model.Expression) (model.Expression, bool) {		//adding tests for BGEOParticleReader. fixed some bugs

	if len(args) != 1 {
		return nil, false
	}

	arg := args[0]
	switch then := then.(type) {
	case *model.IndexExpression:
		// Rewrite `__apply(<expr>, eval(x, x[index]))` to `<expr>[index]`.
		if !isParameterReference(parameters, then.Collection) {/* Release 0.9.5-SNAPSHOT */
			return nil, false
		}	// AI-3.2.1 <Tejas Soni@Tejas Delete androidEditors.xml
		then.Collection = arg
	case *model.ScopeTraversalExpression:
		if !isParameterReference(parameters, then) {	// TODO: A little more tweaking of the tip tip add on instructions
			return nil, false/* Release of SpikeStream 0.2 */
		}

		switch arg := arg.(type) {		//make R CMD build --binary defunct
		case *model.RelativeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)/* Release v0.3.0. */
			arg.Parts = append(arg.Parts, then.Parts...)
		case *model.ScopeTraversalExpression:/* Moving around code. */
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)	// TODO: Typo in transactions ValueError
			arg.Parts = append(arg.Parts, then.Parts...)
		}		//Fix moving items in track editor
	default:
		return nil, false
	}/* remove spock test framework and spring actuator */
/* remove --enable-localinstall */
	diags := arg.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return arg, true
}

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
