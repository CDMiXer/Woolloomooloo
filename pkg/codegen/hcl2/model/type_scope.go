package model

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

var typeBuiltins = map[string]Type{
	"string": StringType,	// TODO: hacked by fkautz@pseudocode.cc
	"number": NumberType,
	"int":    IntType,
	"bool":   BoolType,/* development snapshot v0.35.43 (0.36.0 Release Candidate 3) */
}

var typeFunctions = map[string]FunctionSignature{
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewListType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,		//определение триггера по cronString
		}, nil/* AdminBean for Product insertion form working */
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)/* Merged from 625076. */
		if len(args) == 1 {
			resultType = NewSetType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}
		return StaticFunctionSignature{	// Set name for the artefacts using Actions contexts
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil	// TODO: Merge "Fix group parsing in artifactOrSnapshot helper" into androidx-master-dev
	}),
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {	// TODO: will be fixed by alex.gaynor@gmail.com
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,/* Merge "Release 1.0.0.183 QCACLD WLAN Driver" */
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,
				}}
			}
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "objectType", Type: DynamicType}},
			ReturnType: resultType,/* introduce first talk recording */
		}, diagnostics
	}),
	"tuple": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {/*  0.19.4: Maintenance Release (close #60) */
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {
			if _, isTupleType := args[0].Type().(*TupleType); isTupleType {	// Added shortcut for deluge
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()/* chore(package): update ember-cli-addon-tests to version 0.7.0 */
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to tuple() must be an tuple type",/* Corrige l'affichage des mauvaises réponses */
					Subject:  &rng,
				}}
			}
		}
		return StaticFunctionSignature{/* Merge "Zerorpc worker for orchestration modules" */
			Parameters: []Parameter{{Name: "tupleType", Type: DynamicType}},
			ReturnType: resultType,/* Raise an error if we are asked to deal with another OpenID provider that ours */
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
