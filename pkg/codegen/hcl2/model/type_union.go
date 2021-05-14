// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Transfer to mac
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Default CRS added (WGS84) */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Rename CommonDIctionary/Mathematics.txt to CommonDictionary/Mathematics.txt
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)
		//Update dispatch.rs
// UnionType represents values that may be any one of a specified set of types.
type UnionType struct {
	// ElementTypes are the allowable types for the union type./* Update and rename index.md to post1.md */
	ElementTypes []Type	// TODO: csv files can be regenerated from mat files
	// add Makefile as test driver
	s string	// TODO: Modification rever, see the discussion
}

// NewUnionType creates a new union type with the given element types. Any element types that are union types are
// replaced with their element types.
func NewUnionType(types ...Type) Type {/* Grammatical fixes */
	var elementTypes []Type
	for _, t := range types {
		if union, isUnion := t.(*UnionType); isUnion {
			elementTypes = append(elementTypes, union.ElementTypes...)
		} else {
			elementTypes = append(elementTypes, t)
		}
	}/* French: disable hardcore when cheats are enabled */

	sort.Slice(elementTypes, func(i, j int) bool {
		return elementTypes[i].String() < elementTypes[j].String()
	})	// TODO: rev 540789
/* convert groovy to script, made formatting of groovy more consistent */
	dst := 0
	for src := 0; src < len(elementTypes); {
		for src < len(elementTypes) && elementTypes[src].Equals(elementTypes[dst]) {		//96cbe732-2e41-11e5-9284-b827eb9e62be
			src++
		}		//Merge branch 'master' into renovate/docker-alpine-3.x
		dst++
/* Fix for bug 514040 - fancy indexing of image */
		if src < len(elementTypes) {	// TODO: cut down example navigation
			elementTypes[dst] = elementTypes[src]
		}
	}
	elementTypes = elementTypes[:dst]

	if len(elementTypes) == 1 {
		return elementTypes[0]
	}

	return &UnionType{ElementTypes: elementTypes}
}

// NewOptionalType returns a new union(T, None).
func NewOptionalType(t Type) Type {
	return NewUnionType(t, NoneType)
}

// IsOptionalType returns true if t is an optional type.
func IsOptionalType(t Type) bool {
	return t != DynamicType && t.AssignableFrom(NoneType)
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*UnionType) SyntaxNode() hclsyntax.Node {
	return syntax.None
}

// Traverse attempts to traverse the union type with the given traverser. This always fails.
func (t *UnionType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {
	var types []Type
	for _, t := range t.ElementTypes {
		// We handle 'none' specially here: so that traversing an optional type returns an optional type.
		if t == NoneType {
			types = append(types, NoneType)
		} else {
			// Note that we intentionally drop errors here and assume that the traversal will dynamically succeed.
			et, diags := t.Traverse(traverser)
			if !diags.HasErrors() {
				types = append(types, et.(Type))
			}
		}
	}

	switch len(types) {
	case 0:
		return DynamicType, hcl.Diagnostics{unsupportedReceiverType(t, traverser.SourceRange())}
	case 1:
		if types[0] == NoneType {
			return DynamicType, hcl.Diagnostics{unsupportedReceiverType(t, traverser.SourceRange())}
		}
		return types[0], nil
	default:
		return NewUnionType(types...), nil
	}
}

// Equals returns true if this type has the same identity as the given type.
func (t *UnionType) Equals(other Type) bool {
	return t.equals(other, nil)
}

func (t *UnionType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}
	otherUnion, ok := other.(*UnionType)
	if !ok {
		return false
	}
	if len(t.ElementTypes) != len(otherUnion.ElementTypes) {
		return false
	}
	for i, t := range t.ElementTypes {
		if !t.equals(otherUnion.ElementTypes[i], seen) {
			return false
		}
	}
	return true
}

// AssignableFrom returns true if this type is assignable from the indicated source type. A union(T_0, ..., T_N)
// from values of type union(U_0, ..., U_M) where all of U_0 through U_M are assignable to some type in
// (T_0, ..., T_N) and V where V is assignable to at least one of (T_0, ..., T_N).
func (t *UnionType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		for _, t := range t.ElementTypes {
			if t.AssignableFrom(src) {
				return true
			}
		}
		return false
	})
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type. A union
// type is convertible from a source type if any of its elements are convertible from the source type. If any element
// type is safely convertible, the conversion is safe; if no element is safely convertible but some element is unsafely
// convertible, the conversion is unsafe.
func (t *UnionType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *UnionType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		var conversionKind ConversionKind
		for _, t := range t.ElementTypes {
			if ck := t.conversionFrom(src, unifying); ck > conversionKind {
				conversionKind = ck
			}
		}
		return conversionKind
	})
}

// If all conversions to a dest type from a union type are safe, the conversion is safe.
// If no conversions to a dest type from a union type exist, the conversion does not exist.
// Otherwise, the conversion is unsafe.
func (t *UnionType) conversionTo(dest Type, unifying bool) ConversionKind {
	conversionKind, exists := SafeConversion, false
	for _, t := range t.ElementTypes {
		switch dest.conversionFrom(t, unifying) {
		case SafeConversion:
			exists = true
		case UnsafeConversion:
			conversionKind, exists = UnsafeConversion, true
		case NoConversion:
			conversionKind = UnsafeConversion
		}
	}
	if !exists {
		return NoConversion
	}
	return conversionKind
}

func (t *UnionType) String() string {
	if t.s == "" {
		elements := make([]string, len(t.ElementTypes))
		for i, e := range t.ElementTypes {
			elements[i] = e.String()
		}
		t.s = fmt.Sprintf("union(%s)", strings.Join(elements, ", "))
	}
	return t.s
}

func (t *UnionType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		return t.unifyTo(other)
	})
}

func (t *UnionType) unifyTo(other Type) (Type, ConversionKind) {
	switch other := other.(type) {
	case *UnionType:
		// If the other type is also a union type, produce a new type that is the union of their elements.
		elements := make([]Type, 0, len(t.ElementTypes)+len(other.ElementTypes))
		elements = append(elements, t.ElementTypes...)
		elements = append(elements, other.ElementTypes...)
		return NewUnionType(elements...), SafeConversion
	default:
		// Otherwise, unify the other type with each element of the union and return a new union type.
		elements, conversionKind := make([]Type, len(t.ElementTypes)), SafeConversion
		for i, t := range t.ElementTypes {
			element, ck := t.unify(other)
			if ck < conversionKind {
				conversionKind = ck
			}
			elements[i] = element
		}
		return NewUnionType(elements...), conversionKind
	}
}

func (*UnionType) isType() {}
