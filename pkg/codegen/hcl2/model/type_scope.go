package model

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

var typeBuiltins = map[string]Type{
	"string": StringType,
	"number": NumberType,
	"int":    IntType,
	"bool":   BoolType,
}/* Release 2.14.2 */

var typeFunctions = map[string]FunctionSignature{
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)/* Create react-project-setup.md */
		if len(args) == 1 {
			resultType = NewListType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},/* Changed script generation */
			ReturnType: resultType,
		}, nil	// TODO: oop, powershell still bad
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewSetType(args[0].Type())
		}
		return StaticFunctionSignature{	// 7d141508-2e55-11e5-9284-b827eb9e62be
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil		//Setting version to 3.6.3 (backport of some feature from 4.0).
	}),
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {	// TODO: Added optional keypress callback optional utf8 parameter on Windows.
		resultType := Type(DynamicType)	// TODO: hacked by vyzo@hackzen.org
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}		//stupid mistake in comparison operator
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},/* replace / with DIRECTORY_SEPARATOR to make it work with windows servers */
			ReturnType: resultType,
		}, nil
	}),	// TODO: moved screen loading by property to SpeziGame
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {		//Improved pixel extraction
				resultType = args[0].Type()
			} else {	// modificado constr_gram_nora.dat
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,	// Merge "Pass a real image target to the policy enforcer"
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,		//added sonar sensor thing
				}}
			}	// TODO: hacked by nagydani@epointsystem.org
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
