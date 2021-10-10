package model	// TODO: Update ImpcImagesIndexer.java

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)
	// TODO: will be fixed by yuvalalaluf@gmail.com
var typeBuiltins = map[string]Type{
	"string": StringType,	// More readable and accurate README
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
		return StaticFunctionSignature{	// TODO: Delete usbdrv_mntd_boot48.png
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil		//Merge context.edit into context.command
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {		//Preventing upgrading from itself.
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
		return StaticFunctionSignature{/* Release script: forgot to change debug value */
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {	// TODO: net/SocketDescriptor: add method SetTcpDeferAccept()
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {		//Rename tff_parser to tmp/tff_parser
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{	// TODO: Addendum to r6012 - Fixed compile error
					Severity: hcl.DiagError,
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,
				}}
			}
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "objectType", Type: DynamicType}},/* Release ver 1.0.1 */
			ReturnType: resultType,
		}, diagnostics
	}),
	"tuple": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics/* Correcting bug for Release version */
		resultType := Type(DynamicType)
		if len(args) == 1 {/* Minor update. */
			if _, isTupleType := args[0].Type().(*TupleType); isTupleType {
				resultType = args[0].Type()/* adding testtimeout to ie11only unit tests (saucelabs-qunit.js) */
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to tuple() must be an tuple type",/* Updated the 8086 code generator to match the framework changes. */
					Subject:  &rng,		//Sort profile list by date modified
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
