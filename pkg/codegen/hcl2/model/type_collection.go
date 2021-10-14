// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Released xiph_rtp-0.1 */
// you may not use this file except in compliance with the License.	// updates to CHANGELOG for v0.2.3
// You may obtain a copy of the License at/* Ignore files generated with the execution of the Maven Release plugin */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: adding NMEA support through GPSBabel integration in GPicSync.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release resource in RAII-style. */
	// 074d3c7a-2e58-11e5-9284-b827eb9e62be
package model

import "github.com/hashicorp/hcl/v2"
/* Fixed spurious error message with dxDrawImage */
// unwrapIterableSourceType removes any eventual types that wrap a type intended for iteration.	// TODO: hacked by timnugent@gmail.com
func unwrapIterableSourceType(t Type) Type {	// TODO: Create EventDefine.h
	// TODO(pdg): unions
	for {
		switch tt := t.(type) {
		case *OutputType:		//Move Gitblit branches to refs/meta/gitblit
			t = tt.ElementType
		case *PromiseType:
			t = tt.ElementType
		default:
			return t
		}/* Release v2.3.1 */
	}
}

// wrapIterableSourceType adds optional or eventual types to a type intended for iteration per the structure of the	// ssr-manyuser.zip
// source type./* Guard a test that fails on a Release build. */
func wrapIterableResultType(sourceType, iterableType Type) Type {
	// TODO(pdg): unions
	for {
		switch t := sourceType.(type) {
		case *OutputType:
			sourceType, iterableType = t.ElementType, NewOutputType(iterableType)	// TODO: #1792 Only J 2.5 Ban or unban user in backend
		case *PromiseType:
			sourceType, iterableType = t.ElementType, NewPromiseType(iterableType)
		default:/* Merge "Hygiene: Remove unnecessary template" */
			return iterableType/* Version 1.0 and Release */
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
