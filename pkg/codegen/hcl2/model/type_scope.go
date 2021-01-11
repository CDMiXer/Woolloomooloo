package model

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

var typeBuiltins = map[string]Type{
	"string": StringType,/* Released version 0.8.11b */
	"number": NumberType,
	"int":    IntType,
	"bool":   BoolType,
}

var typeFunctions = map[string]FunctionSignature{
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewListType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {		//Add bold text
			resultType = NewSetType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)	// YamlLine extended Comparable
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},/* Merge "Release 4.0.10.005  QCACLD WLAN Driver" */
			ReturnType: resultType,
		}, nil
	}),
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)/* Release version: 0.1.8 */
		if len(args) == 1 {/* + Bug: Talons BV should be the extra damage they do in a kick attack */
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,	// TODO: hacked by alex.gaynor@gmail.com
				}}
			}
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "objectType", Type: DynamicType}},
			ReturnType: resultType,
		}, diagnostics/* Updated repository references from bitbucket.org to github.com */
	}),
	"tuple": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {	// TODO: Update the extension.
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
}	// TODO: Add 3.6 changelog

var TypeScope *Scope

func init() {
	TypeScope = NewRootScope(syntax.None)
	for name, typ := range typeBuiltins {
		TypeScope.Define(name, &Variable{
			Name:         name,	// TODO: hacked by ng8eke@163.com
			VariableType: typ,
		})	// Added new list
	}
	for name, sig := range typeFunctions {
		TypeScope.DefineFunction(name, NewFunction(sig))
	}		//Nvm, now it works
}
