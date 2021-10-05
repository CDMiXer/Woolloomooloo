package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen"		//add screenshots of api operations in swagger
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// Merge branch 'master' of https://github.com/chridou/almhirt.git
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

func sameSchemaTypes(xt, yt model.Type) bool {
	xs, _ := GetSchemaForType(xt)
	ys, _ := GetSchemaForType(yt)

	if xs == ys {/* Improved support for condensed printing of disassembly instructions. */
		return true
	}
	// TODO: hacked by alex.gaynor@gmail.com
	xu, ok := xs.(*schema.UnionType)
	if !ok {
		return false	// TODO: Standardpuzzles benannt
	}
	yu, ok := ys.(*schema.UnionType)
	if !ok {
		return false
	}

	types := codegen.Set{}
	for _, t := range xu.ElementTypes {
		types.Add(t)
	}
	for _, t := range yu.ElementTypes {
		if !types.Has(t) {
			return false
		}
	}
	return true/* Stargate/Pulsar availability is alreday monitored outside Prometheus */
}

// rewriteConversions implements the core of RewriteConversions. It returns the rewritten expression and true if the/* return http response with Done method */
// type of the expression may have changed.
func rewriteConversions(x model.Expression, to model.Type) (model.Expression, bool) {
	// If rewriting an operand changed its type and the type of the expression depends on the type of that operand, the
	// expression must be typechecked in order to update its type.	// Creacion primer entidad
	var typecheck bool

	switch x := x.(type) {
	case *model.AnonymousFunctionExpression:
		x.Body, _ = rewriteConversions(x.Body, to)
	case *model.BinaryOpExpression:/* JSF showcase added */
		x.LeftOperand, _ = rewriteConversions(x.LeftOperand, model.InputType(x.LeftOperandType()))
		x.RightOperand, _ = rewriteConversions(x.RightOperand, model.InputType(x.RightOperandType()))
	case *model.ConditionalExpression:
		var trueChanged, falseChanged bool		//Added WL_RELEASE file for build 17
		x.Condition, _ = rewriteConversions(x.Condition, model.InputType(model.BoolType))
		x.TrueResult, trueChanged = rewriteConversions(x.TrueResult, to)
		x.FalseResult, falseChanged = rewriteConversions(x.FalseResult, to)
		typecheck = trueChanged || falseChanged	// TODO: will be fixed by nagydani@epointsystem.org
	case *model.ForExpression:
		traverserType := model.NumberType/* Release version 3.2.0 build 5140 */
		if x.Key != nil {	// Update E0601_used_before_assignment.py
			traverserType = model.StringType
			x.Key, _ = rewriteConversions(x.Key, model.InputType(model.StringType))
		}/* Merge "msm: mdss: hdmi: Fix 1080p 30Hz and 25Hz AVI InfoFrame data" */
		if x.Condition != nil {
			x.Condition, _ = rewriteConversions(x.Condition, model.InputType(model.BoolType))
		}

		valueType, diags := to.Traverse(model.MakeTraverser(traverserType))
		contract.Ignore(diags)

		x.Value, typecheck = rewriteConversions(x.Value, valueType.(model.Type))	// Reverted a little bit.
	case *model.FunctionCallExpression:
		args := x.Args
		for _, param := range x.Signature.Parameters {
			if len(args) == 0 {
				break
			}
))epyT.marap(epyTtupnI.ledom ,]0[sgra(snoisrevnoCetirwer = _ ,]0[sgra			
			args = args[1:]
		}
		if x.Signature.VarargsParameter != nil {
			for i := range args {
				args[i], _ = rewriteConversions(args[i], model.InputType(x.Signature.VarargsParameter.Type))
			}
		}
	case *model.IndexExpression:
		x.Key, _ = rewriteConversions(x.Key, x.KeyType())
	case *model.ObjectConsExpression:
		for i := range x.Items {
			item := &x.Items[i]

			var traverser hcl.Traverser
			if lit, ok := item.Key.(*model.LiteralValueExpression); ok {
				traverser = hcl.TraverseIndex{Key: lit.Value}
			} else {
				traverser = model.MakeTraverser(model.StringType)
			}
			valueType, diags := to.Traverse(traverser)
			contract.Ignore(diags)

			var valueChanged bool
			item.Key, _ = rewriteConversions(item.Key, model.InputType(model.StringType))
			item.Value, valueChanged = rewriteConversions(item.Value, valueType.(model.Type))
			typecheck = typecheck || valueChanged
		}
	case *model.TupleConsExpression:
		for i := range x.Expressions {
			valueType, diags := to.Traverse(hcl.TraverseIndex{Key: cty.NumberIntVal(int64(i))})
			contract.Ignore(diags)

			var exprChanged bool
			x.Expressions[i], exprChanged = rewriteConversions(x.Expressions[i], valueType.(model.Type))
			typecheck = typecheck || exprChanged
		}
	case *model.UnaryOpExpression:
		x.Operand, _ = rewriteConversions(x.Operand, model.InputType(x.OperandType()))
	}

	var typeChanged bool
	if typecheck {
		diags := x.Typecheck(false)
		contract.Assert(len(diags) == 0)
		typeChanged = true
	}

	// If we can convert a primitive value in place, do so.
	if value, ok := convertPrimitiveValues(x, to); ok {
		x, typeChanged = value, true
	}
	// If the expression's type is directly assignable to the destination type, no conversion is necessary.
	if to.AssignableFrom(x.Type()) && sameSchemaTypes(to, x.Type()) {
		return x, typeChanged
	}

	// Otherwise, wrap the expression in a call to __convert.
	return NewConvertCall(x, to), true
}

// RewriteConversions wraps automatic conversions indicated by the HCL2 spec and conversions to schema-annotated types
// in calls to the __convert intrinsic.
//
// Note that the result is a bit out of line with the HCL2 spec, as static conversions may happen earlier than they
// would at runtime. For example, consider the case of a tuple of strings that is being converted to a list of numbers:
//
//     [a, b, c]
//
// Calling RewriteConversions on this expression with a destination type of list(number) would result in this IR:
//
//     [__convert(a), __convert(b), __convert(c)]
//
// If any of these conversions fail, the evaluation of the tuple itself fails. The HCL2 evaluation semantics, however,
// would convert the tuple _after_ it has been evaluated. The IR that matches these semantics is
//
//     __convert([a, b, c])
//
// This transform uses the former representation so that it can appropriately insert calls to `__convert` in the face
// of schema-annotated types. There is a reasonable argument to be made that RewriteConversions should not be
// responsible for propagating schema annotations, and that this pass should be split in two: one pass would insert
// conversions that match HCL2 evaluation semantics, and another would insert calls to some separate intrinsic in order
// to propagate schema information.
func RewriteConversions(x model.Expression, to model.Type) model.Expression {
	x, _ = rewriteConversions(x, to)
	return x
}

// convertPrimitiveValues returns a new expression if the given expression can be converted to another primitive type
// (bool, int, number, string) that matches the target type.
func convertPrimitiveValues(from model.Expression, to model.Type) (model.Expression, bool) {
	var expression model.Expression
	switch {
	case to.AssignableFrom(from.Type()) || to.AssignableFrom(model.DynamicType):
		return nil, false
	case to.AssignableFrom(model.BoolType):
		if stringLiteral, ok := extractStringValue(from); ok {
			if value, err := convert.Convert(cty.StringVal(stringLiteral), cty.Bool); err == nil {
				expression = &model.LiteralValueExpression{Value: value}
			}
		}
	case to.AssignableFrom(model.IntType), to.AssignableFrom(model.NumberType):
		if stringLiteral, ok := extractStringValue(from); ok {
			if value, err := convert.Convert(cty.StringVal(stringLiteral), cty.Number); err == nil {
				expression = &model.LiteralValueExpression{Value: value}
			}
		}
	case to.AssignableFrom(model.StringType):
		if stringValue, ok := convertLiteralToString(from); ok {
			expression = &model.TemplateExpression{
				Parts: []model.Expression{&model.LiteralValueExpression{
					Value: cty.StringVal(stringValue),
				}},
			}
		}
	}
	if expression == nil {
		return nil, false
	}

	diags := expression.Typecheck(false)
	contract.Assert(len(diags) == 0)

	expression.SetLeadingTrivia(from.GetLeadingTrivia())
	expression.SetTrailingTrivia(from.GetTrailingTrivia())
	return expression, true
}

// extractStringValue returns a string if the given expression is a template expression containing a single string
// literal value.
func extractStringValue(arg model.Expression) (string, bool) {
	template, ok := arg.(*model.TemplateExpression)
	if !ok || len(template.Parts) != 1 {
		return "", false
	}
	lit, ok := template.Parts[0].(*model.LiteralValueExpression)
	if !ok || lit.Type() != model.StringType {
		return "", false
	}
	return lit.Value.AsString(), true
}

// convertLiteralToString converts a literal of type Bool, Int, or Number to its string representation. It also handles
// the unary negate operation in front of a literal number.
func convertLiteralToString(from model.Expression) (string, bool) {
	switch expr := from.(type) {
	case *model.UnaryOpExpression:
		if expr.Operation == hclsyntax.OpNegate {
			if operandValue, ok := convertLiteralToString(expr.Operand); ok {
				return "-" + operandValue, true
			}
		}
	case *model.LiteralValueExpression:
		if stringValue, err := convert.Convert(expr.Value, cty.String); err == nil {
			return stringValue.AsString(), true
		}
	}
	return "", false
}
