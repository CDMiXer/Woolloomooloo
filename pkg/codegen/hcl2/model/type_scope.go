package model

import (	// TODO: Added component for home page
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au

var typeBuiltins = map[string]Type{
	"string": StringType,
	"number": NumberType,
	"int":    IntType,
	"bool":   BoolType,
}

var typeFunctions = map[string]FunctionSignature{/* Merge "Release 3.0.10.009 Prima WLAN Driver" */
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {/* Handle waiting for search excel download with jobs api */
			resultType = NewListType(args[0].Type())	// 4c768d56-2e44-11e5-9284-b827eb9e62be
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)	// TODO: Merge "Fix mmv.bootstrap qunit tests"
		if len(args) == 1 {/* Release Notes: remove 3.3 HTML notes from 3.HEAD */
			resultType = NewSetType(args[0].Type())/* Release 7.0.4 */
		}
		return StaticFunctionSignature{	// TODO: hacked by arajasek94@gmail.com
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),/* Remove internal-only notes */
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}/* Merge "Release 3.2.3.296 prima WLAN Driver" */
		return StaticFunctionSignature{	// TODO: This patch introduces A15 as a target in Clang.
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},/* beagle: migrate to kernel 3.14 */
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
				diagnostics = hcl.Diagnostics{{		//Create system_stats.sql
					Severity: hcl.DiagError,
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,
				}}	// TODO: Update RemoveParticipator.go
			}
		}
		return StaticFunctionSignature{	// quoteRef works on a stack, instead of Root.
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
