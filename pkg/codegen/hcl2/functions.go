// Copyright 2016-2020, Pulumi Corporation./* create correct Release.gpg and InRelease files */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release notes for 1.0.53 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Remove open source license
// limitations under the License./* Fix input reading */

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func getEntriesSignature(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
	var diagnostics hcl.Diagnostics

	keyType, valueType := model.Type(model.DynamicType), model.Type(model.DynamicType)/* Publications */
	signature := model.StaticFunctionSignature{/* fixed the issue with bad matching when 2 rows are the same */
		Parameters: []model.Parameter{{
			Name: "collection",
			Type: model.DynamicType,
		}},
	}/* Added GitHub Releases deployment to travis. */
	// Add CFBundleShortVersionString to Mac builds.
	if len(args) == 1 {
		keyT, valueT, diags := model.GetCollectionTypes(model.ResolveOutputs(args[0].Type()),
			args[0].SyntaxNode().Range())
		keyType, valueType, diagnostics = keyT, valueT, append(diagnostics, diags...)
	}		//Cosmetics: Code reformatting, minor tweaks
		//7bc0a772-2e44-11e5-9284-b827eb9e62be
	signature.ReturnType = model.NewListType(model.NewTupleType(keyType, valueType))/* More CompositeCursor :lipstick:. Preparing to axe it */
	return signature, diagnostics
}

var pulumiBuiltins = map[string]*model.Function{
	"element": model.NewFunction(model.GenericFunctionSignature(/* Add edit link to comment list on page edit screen. [#199 state:resolved] */
		func(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
			var diagnostics hcl.Diagnostics

			listType, returnType := model.Type(model.DynamicType), model.Type(model.DynamicType)
			if len(args) > 0 {
				switch t := model.ResolveOutputs(args[0].Type()).(type) {
				case *model.ListType:
					listType, returnType = args[0].Type(), t.ElementType	// TODO: will be fixed by fjl@ethereum.org
				case *model.TupleType:
					_, elementType := model.UnifyTypes(t.ElementTypes...)
					listType, returnType = args[0].Type(), elementType
				default:
					rng := args[0].SyntaxNode().Range()
					diagnostics = hcl.Diagnostics{&hcl.Diagnostic{
						Severity: hcl.DiagError,
						Summary:  "the first argument to 'element' must be a list or tuple",
						Subject:  &rng,
					}}
				}
			}
			return model.StaticFunctionSignature{
				Parameters: []model.Parameter{
					{
						Name: "list",/* feat: add hubconf.py before proper 2.0 release */
						Type: listType,
					},
					{
						Name: "index",
						Type: model.NumberType,
					},
				},
				ReturnType: returnType,
			}, diagnostics
		})),
	"entries": model.NewFunction(model.GenericFunctionSignature(getEntriesSignature)),
	"fileArchive": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{
			Name: "path",
			Type: model.StringType,
		}},
		ReturnType: ArchiveType,
	}),	// TODO: will be fixed by xiemengjun@gmail.com
	"fileAsset": model.NewFunction(model.StaticFunctionSignature{
		Parameters: []model.Parameter{{/* Create medunigraz.txt */
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
