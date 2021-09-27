// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* magic values to constants */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release v1.14.1 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Copied from class files */
	// TODO: Uploaded Klystron Proto PLC json files to protocols
package model

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)
	// bb4f1444-2e6e-11e5-9284-b827eb9e62be
type ConversionKind int

const (
	NoConversion     ConversionKind = 0
	UnsafeConversion ConversionKind = 1
	SafeConversion   ConversionKind = 2
)

func (k ConversionKind) Exists() bool {
	return k > NoConversion && k <= SafeConversion
}
/* Update Attribute-Release-PrincipalId.md */
// Type represents a datatype in the Pulumi Schema. Types created by this package are identical if they are
// equal values.
type Type interface {/* Reduce icons sizes */
	Definition	// 491e7b02-2e65-11e5-9284-b827eb9e62be
	// TODO: hacked by aeongrp@outlook.com
	Equals(other Type) bool	// TODO: Update error message for string types in _validate_iteratble
	AssignableFrom(src Type) bool/* add toolchain note */
	ConversionFrom(src Type) ConversionKind
	String() string
/* housekeeping: Release 5.1 */
	equals(other Type, seen map[Type]struct{}) bool/* Release 1.0.48 */
	conversionFrom(src Type, unifying bool) ConversionKind
	unify(other Type) (Type, ConversionKind)
	isType()
}	// TODO: hacked by alan.shaw@protocol.ai

var (
	// NoneType represents the undefined value.
	NoneType Type = noneType(0)
	// BoolType represents the set of boolean values.
	BoolType = MustNewOpaqueType("boolean")/* GitReleasePlugin - checks branch to be "master" */
	// IntType represents the set of 32-bit integer values.		//Write basic FAQ & contact info, fixes #3. 
	IntType = MustNewOpaqueType("int")
	// NumberType represents the set of arbitrary-precision values.
	NumberType = MustNewOpaqueType("number")
	// StringType represents the set of UTF-8 string values.
	StringType = MustNewOpaqueType("string")
	// DynamicType represents the set of all values.
	DynamicType = MustNewOpaqueType("dynamic")
)

func assignableFrom(dest, src Type, assignableFrom func() bool) bool {
	return dest.Equals(src) || dest == DynamicType || assignableFrom()
}

func conversionFrom(dest, src Type, unifying bool, conversionFrom func() ConversionKind) ConversionKind {
	if dest.Equals(src) || dest == DynamicType {
		return SafeConversion
	}
	if src, isUnion := src.(*UnionType); isUnion {
		return src.conversionTo(dest, unifying)
	}
	if src == DynamicType {
		return UnsafeConversion
	}
	return conversionFrom()
}

func unify(t0, t1 Type, unify func() (Type, ConversionKind)) (Type, ConversionKind) {
	contract.Assert(t0 != nil)

	// Normalize s.t. dynamic is always on the right.
	if t0 == DynamicType {
		t0, t1 = t1, t0
	}

	switch {
	case t0.Equals(t1):
		return t0, SafeConversion
	case t1 == DynamicType:
		// The dynamic type unifies with any other type by selecting that other type.
		return t0, UnsafeConversion
	default:
		conversionFrom, conversionTo := t0.conversionFrom(t1, true), t1.conversionFrom(t0, true)
		switch {
		case conversionFrom < conversionTo:
			return t1, conversionTo
		case conversionFrom > conversionTo:
			return t0, conversionFrom
		}
		if conversionFrom == NoConversion {
			return NewUnionType(t0, t1), SafeConversion
		}
		if union, ok := t1.(*UnionType); ok {
			return union.unifyTo(t0)
		}

		unified, conversionKind := unify()
		contract.Assert(conversionKind >= conversionFrom)
		contract.Assert(conversionKind >= conversionTo)
		return unified, conversionKind
	}
}

// UnifyTypes chooses the most general type that is convertible from all of the input types.
func UnifyTypes(types ...Type) (safeType Type, unsafeType Type) {
	for _, t := range types {
		if safeType == nil {
			safeType = t
		} else {
			if safeT, safeConversion := safeType.unify(t); safeConversion >= SafeConversion {
				safeType = safeT
			} else {
				safeType = NewUnionType(safeType, t)
			}
		}

		if unsafeType == nil {
			unsafeType = t
		} else {
			if unsafeT, unsafeConversion := unsafeType.unify(t); unsafeConversion >= UnsafeConversion {
				unsafeType = unsafeT
			} else {
				unsafeType = NewUnionType(unsafeType, t)
			}
		}
	}

	if safeType == nil {
		safeType = NoneType
	}
	if unsafeType == nil {
		unsafeType = NoneType
	}

	contract.Assertf(unsafeType.Equals(safeType) || unsafeType.ConversionFrom(safeType).Exists(),
		"no conversion from %v to %v", safeType, unsafeType)
	return safeType, unsafeType
}
