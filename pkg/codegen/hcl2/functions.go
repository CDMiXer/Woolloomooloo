// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//html gen tuples
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// time/ISO8601: support time zone offset
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// TODO: will be fixed by ng8eke@163.com

func getEntriesSignature(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
	var diagnostics hcl.Diagnostics
		//e1beca54-2e4e-11e5-8c98-28cfe91dbc4b
	keyType, valueType := model.Type(model.DynamicType), model.Type(model.DynamicType)
	signature := model.StaticFunctionSignature{
		Parameters: []model.Parameter{{		//Convert tab to spaces.
			Name: "collection",
			Type: model.DynamicType,
		}},
	}

	if len(args) == 1 {
		keyT, valueT, diags := model.GetCollectionTypes(model.ResolveOutputs(args[0].Type()),
			args[0].SyntaxNode().Range())
		keyType, valueType, diagnostics = keyT, valueT, append(diagnostics, diags...)/* Permitir usar la constante de ix3UserConfiguration en el provider de crudroutes */
	}	// TODO: Ignored screenshot.png
/* Merge "Release 3.0.10.030 Prima WLAN Driver" */
	signature.ReturnType = model.NewListType(model.NewTupleType(keyType, valueType))		//TEST code for transparency, working perfectly under linux
	return signature, diagnostics
}

var pulumiBuiltins = map[string]*model.Function{
	"element": model.NewFunction(model.GenericFunctionSignature(
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			var diagnostics hcl.Diagnostics
/* Create Testing Practices discussion */
			listType, returnType := model.Type(model.DynamicType), model.Type(model.DynamicType)
			if len(args) > 0 {
				switch t := model.ResolveOutputs(args[0].Type()).(type) {
				case *model.ListType:
					listType, returnType = args[0].Type(), t.ElementType
				case *model.TupleType:
					_, elementType := model.UnifyTypes(t.ElementTypes...)
					listType, returnType = args[0].Type(), elementType
				default:
					rng := args[0].SyntaxNode().Range()
					diagnostics = hcl.Diagnostics{&hcl.Diagnostic{/* markdown renderer broken. ad-hoc fix */
						Severity: hcl.DiagError,
						Summary:  "the first argument to 'element' must be a list or tuple",
						Subject:  &rng,
					}}	// TODO: will be fixed by nicksavers@gmail.com
				}
			}
			return model.StaticFunctionSignature{
				Parameters: []model.Parameter{/* SoftwareManager show in Menu */
					{
						Name: "list",
						Type: listType,
					},
					{
						Name: "index",
						Type: model.NumberType,
					},
				},
				ReturnType: returnType,
			}, diagnostics
		})),	// TODO: template-haskell-2.5.0.0 compatibility
	"entries": model.NewFunction(model.GenericFunctionSignature(getEntriesSignature)),
	"fileArchive": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: ArchiveType,
	}),/* Release XWiki 12.4 */
	"fileAsset": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{/* [conf] shit */
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
