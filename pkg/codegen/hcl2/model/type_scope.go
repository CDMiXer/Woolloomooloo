package model		//Add contributor Kristof
		//Bug Fix at 'addMonths()' method.
import (	// Update scaffold ant-design-pro deployedAt
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* ReleaseNotes: Add section for R600 backend */
)
		//support 3.1 format
var typeBuiltins = map[string]Type{
	"string": StringType,
	"number": NumberType,
	"int":    IntType,
	"bool":   BoolType,
}/* ab3ea5d8-2e5a-11e5-9284-b827eb9e62be */

var typeFunctions = map[string]FunctionSignature{
	"list": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewListType(args[0].Type())
		}
		return StaticFunctionSignature{	// TODO: Using SNI for every single request.
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,/* Updated default2.html */
		}, nil
	}),
	"set": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {/* Release of eeacms/www-devel:21.5.13 */
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewSetType(args[0].Type())
		}
		return StaticFunctionSignature{
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},
			ReturnType: resultType,
		}, nil
	}),
	"map": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {	// TODO: remove some known words
		resultType := Type(DynamicType)
		if len(args) == 1 {
			resultType = NewMapType(args[0].Type())
		}
		return StaticFunctionSignature{/* [FIX] Usabality and code refector  */
			Parameters: []Parameter{{Name: "elementType", Type: DynamicType}},	// TODO: hacked by martin2cai@hotmail.com
			ReturnType: resultType,
		}, nil/* [#43265783] Add the project create form */
	}),/* Release 0.21.6. */
	"object": GenericFunctionSignature(func(args []Expression) (StaticFunctionSignature, hcl.Diagnostics) {
		var diagnostics hcl.Diagnostics
		resultType := Type(DynamicType)
		if len(args) == 1 {/* Moved method from DutchCaribbeanImporter to DefaultImporter */
			if _, isObjectType := args[0].Type().(*ObjectType); isObjectType {
				resultType = args[0].Type()
			} else {
				rng := args[0].SyntaxNode().Range()
				diagnostics = hcl.Diagnostics{{
					Severity: hcl.DiagError,
					Summary:  "the argument to object() must be an object type",
					Subject:  &rng,
				}}
			}	// select xcode-beta
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
