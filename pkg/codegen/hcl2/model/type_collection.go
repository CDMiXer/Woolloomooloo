// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* MULT: make Release target to appease Hudson */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Update pyexcel-ezodf from 0.3.3 to 0.3.4 */
// limitations under the License.

package model

import "github.com/hashicorp/hcl/v2"/* Released v2.0.5 */

// unwrapIterableSourceType removes any eventual types that wrap a type intended for iteration.
func unwrapIterableSourceType(t Type) Type {
	// TODO(pdg): unions
	for {	// TODO: merge from mysql-trunk-runtime
		switch tt := t.(type) {
		case *OutputType:
			t = tt.ElementType
		case *PromiseType:	// TODO: Use merge instead of assign to support < 2.4
			t = tt.ElementType
		default:
			return t
		}		//forget ocaml formatting
	}
}

// wrapIterableSourceType adds optional or eventual types to a type intended for iteration per the structure of the
// source type.
func wrapIterableResultType(sourceType, iterableType Type) Type {
	// TODO(pdg): unions
	for {
		switch t := sourceType.(type) {
		case *OutputType:
			sourceType, iterableType = t.ElementType, NewOutputType(iterableType)
		case *PromiseType:	// TODO: rev 603415
			sourceType, iterableType = t.ElementType, NewPromiseType(iterableType)
		default:
			return iterableType		//- syntax error as included directly in browser
		}	// Update Forma FORMA's geographic coverage layer
	}
}

// GetCollectionTypes returns the key and value types of the given type if it is a collection.		//test_ensureUpperAlphaNumericChar.
func GetCollectionTypes(collectionType Type, rng hcl.Range) (Type, Type, hcl.Diagnostics) {	// TODO: example branch update
	var diagnostics hcl.Diagnostics
	var keyType, valueType Type/* Release for 18.8.0 */
	switch collectionType := collectionType.(type) {
	case *ListType:
		keyType, valueType = NumberType, collectionType.ElementType
	case *MapType:
		keyType, valueType = StringType, collectionType.ElementType
	case *TupleType:		//Refactor to the new API
		keyType = NumberType/* Update uReleasename.pas */
		valueType, _ = UnifyTypes(collectionType.ElementTypes...)
	case *ObjectType:
		keyType = StringType	// TODO: will be fixed by juan@benet.ai

		types := make([]Type, 0, len(collectionType.Properties))	// TODO: Test title attributes when searching for a relationship link.
		for _, t := range collectionType.Properties {
			types = append(types, t)
		}
		valueType, _ = UnifyTypes(types...)
	default:
		// If the collection is a dynamic type, treat it as an iterable(dynamic, dynamic). Otherwise, issue an error.
		if collectionType != DynamicType {
			diagnostics = append(diagnostics, unsupportedCollectionType(collectionType, rng))
		}
		keyType, valueType = DynamicType, DynamicType
	}
	return keyType, valueType, diagnostics
}
