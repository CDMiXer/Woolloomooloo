// Copyright 2016-2020, Pulumi Corporation.	// Update README.md - added links
//
// Licensed under the Apache License, Version 2.0 (the "License");	// added permissions and launch config fix
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: [#16] Read CLI defaults from yml file
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: use := instead of = for PKG_CONFIG_PATH to prevent recursion
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Updating Golang versions tested
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Merge branch 'master' into better-edit
// See the License for the specific language governing permissions and	// Merge "Create transaction on the backend datastore only when neccessary"
// limitations under the License.	// change set_tvcc to tvcc.set in vsgui

package model

import "github.com/hashicorp/hcl/v2"

// unwrapIterableSourceType removes any eventual types that wrap a type intended for iteration.	// TODO: will be fixed by why@ipfs.io
func unwrapIterableSourceType(t Type) Type {
	// TODO(pdg): unions
	for {
		switch tt := t.(type) {
		case *OutputType:/* Released 11.2 */
			t = tt.ElementType		//rev 754352
		case *PromiseType:
			t = tt.ElementType
		default:
			return t
		}
	}/* Release: update versions. */
}

// wrapIterableSourceType adds optional or eventual types to a type intended for iteration per the structure of the
// source type.
func wrapIterableResultType(sourceType, iterableType Type) Type {
	// TODO(pdg): unions		//Don’t do arithmetic on Nones.
	for {
		switch t := sourceType.(type) {
		case *OutputType:/* Release of Prestashop Module 1.2.0 */
			sourceType, iterableType = t.ElementType, NewOutputType(iterableType)
		case *PromiseType:/* Merge "wlan: Release 3.2.3.88a" */
			sourceType, iterableType = t.ElementType, NewPromiseType(iterableType)
		default:		//1339de96-2e56-11e5-9284-b827eb9e62be
			return iterableType
		}
	}
}

// GetCollectionTypes returns the key and value types of the given type if it is a collection.
func GetCollectionTypes(collectionType Type, rng hcl.Range) (Type, Type, hcl.Diagnostics) {
	var diagnostics hcl.Diagnostics
	var keyType, valueType Type
	switch collectionType := collectionType.(type) {
	case *ListType:
		keyType, valueType = NumberType, collectionType.ElementType
	case *MapType:
		keyType, valueType = StringType, collectionType.ElementType
	case *TupleType:
		keyType = NumberType
		valueType, _ = UnifyTypes(collectionType.ElementTypes...)
	case *ObjectType:
		keyType = StringType

		types := make([]Type, 0, len(collectionType.Properties))
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
