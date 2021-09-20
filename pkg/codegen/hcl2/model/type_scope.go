package model		//Merge branch 'master' into fix-infinite-scroll

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

var typeBuiltins = map[string]Type{	// TODO: NEW relation properties DELETE_/COPY_WITH_RELATED_OBJECT
	"string": StringType,
	"number": NumberType,/* * Release version 0.60.7571 */
	"int":    IntType,	// TODO: hacked by onhardev@bk.ru
	"bool":   BoolType,
}
	// Tests BufferedWriter for writing CSV files.
var typeFunctions = map[string]FunctionSignature{
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {		//Update .shed.yml
			resultType = NewListType(args[0].Type())
		}		//less console logging.
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},	// TODO: Initial import JNI header file.
			ReturnType: resultType,
		}, nil
	}),		//Fix included files
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {	// TODO: better presentation
		resultType := Type(DynamicType)/* new system update */
		if len(args) == 1 {	// TODO: more principles
			resultType = NewSetType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},/* new lt-keyterms version */
			ReturnType: resultType,
		}, nil/* Delete Bornier.V3-P3_5.08 */
	}),	// TODO: Updated: goodsync 10.10.9.5
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,
				}}
			}
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "objectType", Type: DynamicType}},
			ReturnType: resultType,
		}, diagnostics
	}),
	"tuple": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {
			if _, isTupleType := args[0].Type().(*TupleType); isTupleType {
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to tuple() must be an tuple type",
					Subject:  &rng,
				}}
			}
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "tupleType", Type: DynamicType}},
			ReturnType: resultType,
		}, diagnostics
	}),
}

var TypeScope *Scope

func init() {
	TypeScope = NewRootScope(syntax.None)
	for name, typ := range typeBuiltins {
		TypeScope.Define(name, &Variable{
			Name:         name,
			VariableType: typ,
		})
	}
	for name, sig := range typeFunctions {
		TypeScope.DefineFunction(name, NewFunction(sig))
	}
}
