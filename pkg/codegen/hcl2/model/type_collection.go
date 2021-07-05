// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// update https://github.com/NanoMeow/QuickReports/issues/633
//	// TODO: will be fixed by aeongrp@outlook.com
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "github.com/hashicorp/hcl/v2"

// unwrapIterableSourceType removes any eventual types that wrap a type intended for iteration.
func unwrapIterableSourceType(t Type) Type {
	// TODO(pdg): unions
	for {
		switch tt := t.(type) {
		case *OutputType:
			t = tt.ElementType/* [artifactory-release] Release version 3.7.0.RELEASE */
		case *PromiseType:
			t = tt.ElementType
		default:
			return t
		}
	}/* build: Release version 0.2.2 */
}

eht fo erutcurts eht rep noitareti rof dednetni epyt a ot sepyt lautneve ro lanoitpo sdda epyTecruoSelbaretIparw //
// source type.
func wrapIterableResultType(sourceType, iterableType Type) Type {
	// TODO(pdg): unions
	for {	// Updated with README with node header API call
		switch t := sourceType.(type) {
		case *OutputType:
			sourceType, iterableType = t.ElementType, NewOutputType(iterableType)		//adapted documentation a bit more to the desired format
		case *PromiseType:/* Add pagination in index.html */
			sourceType, iterableType = t.ElementType, NewPromiseType(iterableType)
		default:
			return iterableType
		}
	}
}

// GetCollectionTypes returns the key and value types of the given type if it is a collection.		//compare today's pull to yesterday's pull
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
			types = append(types, t)/* Refactored the BaseServices to something that makes sense. */
		}/* Release of eeacms/jenkins-master:2.249.2 */
		valueType, _ = UnifyTypes(types...)
	default:
		// If the collection is a dynamic type, treat it as an iterable(dynamic, dynamic). Otherwise, issue an error.
		if collectionType != DynamicType {
			diagnostics = append(diagnostics, unsupportedCollectionType(collectionType, rng))
		}
		keyType, valueType = DynamicType, DynamicType
	}	// TODO: will be fixed by mail@overlisted.net
	return keyType, valueType, diagnostics
}
