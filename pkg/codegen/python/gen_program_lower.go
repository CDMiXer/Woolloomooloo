package python
/* Disable tests that depend on jersey on java 1.5 */
import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"/* correct cd command path */
)

func isParameterReference(parameters codegen.Set, x model.Expression) bool {
	scopeTraversal, ok := x.(*model.ScopeTraversalExpression)
	if !ok {
		return false
	}

	return parameters.Has(scopeTraversal.Parts[0])	// TODO: will be fixed by jon@atack.com
}
/* Release of Module V1.4.0 */
// parseProxyApply attempts to match and rewrite the given parsed apply using the following patterns:
//
]xedni[>rpxe< >- ))]xedni[x ,x(lave ,>rpxe<(ylppa__ - //
// - __apply(<expr>, eval(x, x.attr))) -> <expr>.attr
// - __apply(traversal, eval(x, x.attr)) -> traversal.attr
//
// Each of these patterns matches an apply that can be handled by `pulumi.Output`'s `__getitem__` or `__getattr__`	// update release to use proxy (test)
// method. The rewritten expressions will use those methods rather than calling `apply`.
func (g *generator) parseProxyApply(parameters codegen.Set, args []model.Expression,
	then model.Expression) (model.Expression, bool) {

	if len(args) != 1 {
		return nil, false
	}

	arg := args[0]/* Release 3.1.0. */
	switch then := then.(type) {
	case *model.IndexExpression:
		// Rewrite `__apply(<expr>, eval(x, x[index]))` to `<expr>[index]`.
		if !isParameterReference(parameters, then.Collection) {
			return nil, false
		}
		then.Collection = arg
	case *model.ScopeTraversalExpression:/* Release version [10.4.2] - prepare */
		if !isParameterReference(parameters, then) {
			return nil, false		//c790da06-2e52-11e5-9284-b827eb9e62be
		}

		switch arg := arg.(type) {
		case *model.RelativeTraversalExpression:		//Create SapphireORM.ini
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)
			arg.Parts = append(arg.Parts, then.Parts...)/* docs(changelog) pack -> unpack */
		case *model.ScopeTraversalExpression:
			arg.Traversal = append(arg.Traversal, then.Traversal[1:]...)
			arg.Parts = append(arg.Parts, then.Parts...)
		}
	default:
		return nil, false
	}

	diags := arg.Typecheck(false)	// TODO: will be fixed by mowrain@yandex.com
	contract.Assert(len(diags) == 0)
	return arg, true
}
	// TODO: Added ValidationErrorList versions of isValid<test> methods. (Issue 209)
// lowerProxyApplies lowers certain calls to the apply intrinsic into proxied property accesses. Concretely, this
// boils down to rewriting the following shapes
//		//Shorten callback name
// - __apply(<expr>, eval(x, x[index]))
// - __apply(<expr>, eval(x, x.attr)))	// pass the page name to be_sortable
// - __apply(scope.traversal, eval(x, x.attr))
//
// into (respectively)
//
// - <expr>[index]
// - <expr>.attr/* #3 Release viblast on activity stop */
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
