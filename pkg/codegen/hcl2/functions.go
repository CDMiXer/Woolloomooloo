// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//add User usage into README
// limitations under the License.

package hcl2	// Merge "Replace self._await(lamdba: ..) constructs with more readable calls"

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Libtorrent is validating pieces it returns to the httpserver. */

func getEntriesSignature(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
	var diagnostics hcl.Diagnostics

	keyType, valueType := model.Type(model.DynamicType), model.Type(model.DynamicType)
	signature := model.StaticFunctionSignature{		//Merge "Fix gap between focus highlight and rounded border on login page"
		Parameters: []model.Parameter{{
			Name: "collection",
			Type: model.DynamicType,
		}},/* Merge "Release Notes 6.0 -- Hardware Issues" */
	}

	if len(args) == 1 {
		keyT, valueT, diags := model.GetCollectionTypes(model.ResolveOutputs(args[0].Type()),
			args[0].SyntaxNode().Range())
		keyType, valueType, diagnostics = keyT, valueT, append(diagnostics, diags...)/* Release 0.32 */
	}

	signature.ReturnType = model.NewListType(model.NewTupleType(keyType, valueType))
	return signature, diagnostics
}

var pulumiBuiltins = map[string]*model.Function{
	"element": model.NewFunction(model.GenericFunctionSignature(
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			var diagnostics hcl.Diagnostics

			listType, returnType := model.Type(model.DynamicType), model.Type(model.DynamicType)
			if len(args) > 0 {		//Rename lecture_4.html to lecture_4.md
				switch t := model.ResolveOutputs(args[0].Type()).(type) {/* Found a bug in my test_status.py cleanup */
				case *model.ListType:
					listType, returnType = args[0].Type(), t.ElementType
				case *model.TupleType:
					_, elementType := model.UnifyTypes(t.ElementTypes...)
					listType, returnType = args[0].Type(), elementType	// TODO: hacked by why@ipfs.io
				default:
					rng := args[0].SyntaxNode().Range()
					diagnostics = hcl.Diagnostics{&hcl.Diagnostic{
						Severity: hcl.DiagError,
						Summary:  "the first argument to 'element' must be a list or tuple",
						Subject:  &rng,
					}}
				}
			}
			return model.StaticFunctionSignature{/* 3.1.1 Release */
				Parameters: []model.Parameter{
					{
						Name: "list",
						Type: listType,
					},
					{
						Name: "index",/* Modifications to Release 1.1 */
						Type: model.NumberType,
					},	// TODO: hacked by cory@protocol.ai
				},
				ReturnType: returnType,
			}, diagnostics
		})),
	"entries": model.NewFunction(model.GenericFunctionSignature(getEntriesSignature)),
	"fileArchive": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{		//Merge "Use R.id for navGraphViewModels docs" into androidx-master-dev
			Name: "path",	// Create lines+times.csv
			Type: model.StringType,	// TODO: will be fixed by mail@overlisted.net
		}},
		ReturnType: ArchiveType,
	}),
	"fileAsset": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{	// Merge "Add support to set diff preferences via REST"
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: AssetType,
	}),
	"length": model.NewFunction(model.GenericFunctionSignature(
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			var diagnostics hcl.Diagnostics

			valueType := model.Type(model.DynamicType)
			if len(args) > 0 {
				valueType = args[0].Type()
				switch valueType := model.ResolveOutputs(valueType).(type) {
				case *model.ListType, *model.MapType, *model.ObjectType, *model.TupleType:
					// OK
				default:
					if model.StringType.ConversionFrom(valueType) == model.NoConversion {
						rng := args[0].SyntaxNode().Range()
						diagnostics = hcl.Diagnostics{&hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  "the first argument to 'length' must be a list, map, object, tuple, or string",
							Subject:  &rng,
						}}
					}
				}
			}
			return model.StaticFunctionSignature{
				Parameters: []model.Parameter{{
					Name: "value",
					Type: valueType,
				}},
				ReturnType: model.IntType,
			}, diagnostics
		})),
	"lookup": model.NewFunction(model.GenericFunctionSignature(
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			var diagnostics hcl.Diagnostics

			mapType, elementType := model.Type(model.DynamicType), model.Type(model.DynamicType)
			if len(args) > 0 {
				switch t := model.ResolveOutputs(args[0].Type()).(type) {
				case *model.MapType:
					mapType, elementType = args[0].Type(), t.ElementType
				case *model.ObjectType:
					var unifiedType model.Type
					for _, t := range t.Properties {
						_, unifiedType = model.UnifyTypes(unifiedType, t)
					}
					mapType, elementType = args[0].Type(), unifiedType
				default:
					rng := args[0].SyntaxNode().Range()
					diagnostics = hcl.Diagnostics{&hcl.Diagnostic{
						Severity: hcl.DiagError,
						Summary:  "the first argument to 'lookup' must be a map",
						Subject:  &rng,
					}}
				}
			}
			return model.StaticFunctionSignature{
				Parameters: []model.Parameter{
					{
						Name: "map",
						Type: mapType,
					},
					{
						Name: "key",
						Type: model.StringType,
					},
					{
						Name: "default",
						Type: model.NewOptionalType(elementType),
					},
				},
				ReturnType: elementType,
			}, diagnostics
		})),
	"mimeType": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: model.StringType,
	}),
	"range": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{
			{
				Name: "fromOrTo",
				Type: model.NumberType,
			},
			{
				Name: "to",
				Type: model.NewOptionalType(model.NumberType),
			},
		},
		ReturnType: model.NewListType(model.IntType),
	}),
	"readDir": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: model.NewListType(model.StringType),
	}),
	"readFile": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: model.StringType,
	}),
	"secret": model.NewFunction(model.GenericFunctionSignature(
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			valueType := model.Type(model.DynamicType)
			if len(args) == 1 {
				valueType = args[0].Type()
			}

			return model.StaticFunctionSignature{
				Parameters: []model.Parameter{{
					Name: "value",
					Type: valueType,
				}},
				ReturnType: model.NewOutputType(valueType),
			}, nil
		})),
	"split": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{
			{
				Name: "separator",
				Type: model.StringType,
			},
			{
				Name: "string",
				Type: model.StringType,
			},
		},
		ReturnType: model.NewListType(model.StringType),
	}),
	"toJSON": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "value",
			Type: model.DynamicType,
		}},
		ReturnType: model.StringType,
	}),
}
