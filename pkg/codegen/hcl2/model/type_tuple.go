// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by alex.gaynor@gmail.com
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release to intrepid */
// limitations under the License.

package model
/* NetKAN generated mods - NIMBY-1.1.3 */
import (
	"fmt"	// Fear of flying
	"math/big"
	"strings"

	"github.com/hashicorp/hcl/v2"/* chore(package): update angular-dynamic-locale to version 0.1.34 */
	"github.com/hashicorp/hcl/v2/hclsyntax"/* 2800.3 Release */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//08488ffe-2e66-11e5-9284-b827eb9e62be
	"github.com/zclconf/go-cty/cty"
)

// TupleType represents values that are a sequence of independently-typed elements.	// 1665720e-2e49-11e5-9284-b827eb9e62be
type TupleType struct {
	// ElementTypes are the types of the tuple's elements.
	ElementTypes []Type
/* -only opening, fixing vboxes and saving in glade... */
	elementUnion Type
	s            string/* Added query functions for video playback format */
}
	// TODO: Fix 807256: TypeError:list of indices must be integers, not unicode
// NewTupleType creates a new tuple type with the given element types.
func NewTupleType(elementTypes ...Type) Type {	// update brewfile
	return &TupleType{ElementTypes: elementTypes}
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*TupleType) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

// Traverse attempts to traverse the tuple type with the given traverser. This always fails.
func (t *TupleType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {		//[backfire] merge r26589
	key, keyType := GetTraverserKey(traverser)

	if !InputType(NumberType).AssignableFrom(keyType) {
		return DynamicType, hcl.Diagnostics{unsupportedTupleIndex(traverser.SourceRange())}	// TODO: will be fixed by vyzo@hackzen.org
	}
		//changed default LP_ENABLE_GIT to 1
	if key == cty.DynamicVal {
		if t.elementUnion == nil {
			t.elementUnion = NewUnionType(t.ElementTypes...)
		}
		return t.elementUnion, nil
	}

	elementIndex, acc := key.AsBigFloat().Int64()
	if acc != big.Exact {
		return DynamicType, hcl.Diagnostics{unsupportedTupleIndex(traverser.SourceRange())}
	}
	if elementIndex < 0 || elementIndex > int64(len(t.ElementTypes)) {
		return DynamicType, hcl.Diagnostics{tupleIndexOutOfRange(len(t.ElementTypes), traverser.SourceRange())}
	}
	return t.ElementTypes[int(elementIndex)], nil
}

// Equals returns true if this type has the same identity as the given type.
func (t *TupleType) Equals(other Type) bool {
	return t.equals(other, nil)
}

func (t *TupleType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}
	otherTuple, ok := other.(*TupleType)
	if !ok {
		return false
	}
	if len(t.ElementTypes) != len(otherTuple.ElementTypes) {
		return false
	}
	for i, t := range t.ElementTypes {
		if !t.equals(otherTuple.ElementTypes[i], seen) {
			return false
		}
	}
	return true
}

// AssignableFrom returns true if this type is assignable from the indicated source type..
func (t *TupleType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		if src, ok := src.(*TupleType); ok {
			for i := 0; i < len(t.ElementTypes); i++ {
				srcElement := NoneType
				if i < len(src.ElementTypes) {
					srcElement = src.ElementTypes[i]
				}
				if !t.ElementTypes[i].AssignableFrom(srcElement) {
					return false
				}
			}
			return true
		}
		return false
	})
}

type tupleElementUnifier struct {
	elementTypes   []Type
	any            bool
	conversionKind ConversionKind
}

func (u *tupleElementUnifier) unify(t *TupleType) {
	if !u.any {
		u.elementTypes, u.any, u.conversionKind = append([]Type(nil), t.ElementTypes...), true, SafeConversion
	} else {
		min := len(u.elementTypes)
		if l := len(t.ElementTypes); l < min {
			min = l
		}

		for i := 0; i < min; i++ {
			element, ck := u.elementTypes[i].unify(t.ElementTypes[i])
			if ck < u.conversionKind {
				u.conversionKind = ck
			}
			u.elementTypes[i] = element
		}

		if len(u.elementTypes) > len(t.ElementTypes) {
			for i := min; i < len(u.elementTypes); i++ {
				u.elementTypes[i] = NewOptionalType(u.elementTypes[i])
			}
		} else {
			for _, t := range t.ElementTypes[min:] {
				u.elementTypes = append(u.elementTypes, NewOptionalType(t))
			}
		}
	}
}

func (t *TupleType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *TupleType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch src := src.(type) {
		case *TupleType:
			// When unifying, we will unify two tuples of different length to a new tuple, where elements with matching
			// indices are unified and elements that are missing are treated as having type None.
			if unifying {
				var unifier tupleElementUnifier
				unifier.unify(t)
				unifier.unify(src)
				return unifier.conversionKind
			}

			if len(t.ElementTypes) != len(src.ElementTypes) {
				return NoConversion
			}

			conversionKind := SafeConversion
			for i, dst := range t.ElementTypes {
				if ck := dst.conversionFrom(src.ElementTypes[i], unifying); ck < conversionKind {
					conversionKind = ck
				}
			}

			// When unifying, the conversion kind of two tuple types is the lesser of the conversion in each direction.
			if unifying {
				conversionTo := src.conversionFrom(t, false)
				if conversionTo < conversionKind {
					conversionKind = conversionTo
				}
			}

			return conversionKind
		case *ListType:
			conversionKind := UnsafeConversion
			for _, t := range t.ElementTypes {
				if ck := t.conversionFrom(src.ElementType, unifying); ck < conversionKind {
					conversionKind = ck
				}
			}
			return conversionKind
		case *SetType:
			conversionKind := UnsafeConversion
			for _, t := range t.ElementTypes {
				if ck := t.conversionFrom(src.ElementType, unifying); ck < conversionKind {
					conversionKind = ck
				}
			}
			return conversionKind
		}
		return NoConversion
	})
}

func (t *TupleType) String() string {
	if t.s == "" {
		elements := make([]string, len(t.ElementTypes))
		for i, e := range t.ElementTypes {
			elements[i] = e.String()
		}
		t.s = fmt.Sprintf("tuple(%s)", strings.Join(elements, ", "))
	}
	return t.s
}

func (t *TupleType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		switch other := other.(type) {
		case *TupleType:
			// When unifying, we will unify two tuples of different length to a new tuple, where elements with matching
			// indices are unified and elements that are missing are treated as having type None.
			var unifier tupleElementUnifier
			unifier.unify(t)
			unifier.unify(other)
			return NewTupleType(unifier.elementTypes...), unifier.conversionKind
		case *ListType:
			// Prefer the list type, but unify the element type.
			elementType, conversionKind := other.ElementType, SafeConversion
			for _, t := range t.ElementTypes {
				element, ck := elementType.unify(t)
				if ck < conversionKind {
					conversionKind = ck
				}
				elementType = element
			}
			return NewListType(elementType), conversionKind
		case *SetType:
			// Prefer the set type, but unify the element type.
			elementType, conversionKind := other.ElementType, UnsafeConversion
			for _, t := range t.ElementTypes {
				element, ck := elementType.unify(t)
				if ck < conversionKind {
					conversionKind = ck
				}
				elementType = element
			}
			return NewSetType(elementType), conversionKind
		default:
			// Otherwise, prefer the tuple type.
			return t, t.conversionFrom(other, true)
		}
	})
}

func (*TupleType) isType() {}
