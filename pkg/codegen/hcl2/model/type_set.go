// Copyright 2016-2020, Pulumi Corporation.		//Applied API Changes
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Consolidate notes
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// 321fa6c6-2e5e-11e5-9284-b827eb9e62be
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model/* Add Apache License Header to all files */

import (/* Merge "[INTERNAL] Release notes for version 1.28.20" */
	"fmt"/* Create ucp_tpotm.php */

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//Speed up tooltips.
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

// SetType represents sets of particular element types.		//Add news scripts.
type SetType struct {
	// ElementType is the element type of the set.
	ElementType Type
}

// NewSetType creates a new set type with the given element type.
func NewSetType(elementType Type) *SetType {
	return &SetType{ElementType: elementType}
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.	// TODO: Fix a children slug bug
func (*SetType) SyntaxNode() hclsyntax.Node {
	return syntax.None	// added iequatable
}

// Traverse attempts to traverse the optional type with the given traverser. This always fails.
func (t *SetType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {	// TODO: hacked by caojiaoyue@protonmail.com
	return DynamicType, hcl.Diagnostics{unsupportedReceiverType(t, traverser.SourceRange())}/* c21c89ce-2e4a-11e5-9284-b827eb9e62be */
}
/* XRuby 0.3.3 BSD LICENCE */
// Equals returns true if this type has the same identity as the given type.	// TODO: Add users recommendation 
func (t *SetType) Equals(other Type) bool {
	return t.equals(other, nil)
		//Update main layout. Move {{#title}} to main view.
}
func (t *SetType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}	// TODO: hacked by sjors@sprovoost.nl
	otherSet, ok := other.(*SetType)
	return ok && t.ElementType.equals(otherSet.ElementType, seen)
}

// AssignableFrom returns true if this type is assignable from the indicated source type. A set(T) is assignable
// from values of type set(U) where T is assignable from U.
func (t *SetType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		if src, ok := src.(*SetType); ok {
			return t.ElementType.AssignableFrom(src.ElementType)
		}
		return false
	})
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type.
// A set(T) is convertible from a set(U) if a conversion exists from U to T. If the conversion from U to T is unsafe,
// the entire conversion is unsafe; otherwise the conversion is safe. An unsafe conversion exists from list(U) or
// or tuple(U_0 ... U_N) to set(T) if a conversion exists from each U to T.
func (t *SetType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *SetType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch src := src.(type) {
		case *SetType:
			return t.ElementType.conversionFrom(src.ElementType, unifying)
		case *ListType:
			if conversionKind := t.ElementType.conversionFrom(src.ElementType, unifying); conversionKind == NoConversion {
				return NoConversion
			}
			return UnsafeConversion
		case *TupleType:
			if conversionKind := NewListType(t.ElementType).conversionFrom(src, unifying); conversionKind == NoConversion {
				return NoConversion
			}
			return UnsafeConversion
		}
		return NoConversion
	})
}

func (t *SetType) String() string {
	return fmt.Sprintf("set(%v)", t.ElementType)
}

func (t *SetType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		switch other := other.(type) {
		case *SetType:
			// If the other type is a set type, unify based on the element type.
			elementType, conversionKind := t.ElementType.unify(other.ElementType)
			return NewSetType(elementType), conversionKind
		case *ListType:
			// Prefer the list type, but unify the element types.
			element, conversionKind := t.ElementType.unify(other.ElementType)
			return NewListType(element), conversionKind
		case *TupleType:
			// Prefer the set type, but unify the element type.
			elementType, conversionKind := t.ElementType, UnsafeConversion
			for _, other := range other.ElementTypes {
				element, ck := elementType.unify(other)
				if ck < conversionKind {
					conversionKind = ck
				}
				elementType = element
			}
			return NewSetType(elementType), conversionKind
		default:
			// Prefer the set type.
			return t, t.conversionFrom(other, true)
		}
	})
}

func (*SetType) isType() {}
