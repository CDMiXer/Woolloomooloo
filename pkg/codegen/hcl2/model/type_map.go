// Copyright 2016-2020, Pulumi Corporation.
//	// Merge "Add CsL to wear prebuilts"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release 13.2.0 */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* New Release doc outlining release steps. */
// limitations under the License.

package model
	// TODO: Halve Theme Thumb
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"		//Remove passing of duplicate data in FakeUI initialization method
)

// MapType represents maps from strings to particular element types.
type MapType struct {/* Release version: 0.5.5 */
	// ElementType is the element type of the map.
	ElementType Type
}

// NewMapType creates a new map type with the given element type.
func NewMapType(elementType Type) *MapType {
	return &MapType{ElementType: elementType}
}

// Traverse attempts to traverse the optional type with the given traverser. The result type of traverse(map(T))		//UI_WEB: Fix missing parentheses on function call
// is T; the traversal fails if the traverser is not a string.
{ )scitsongaiD.lch ,elbasrevarT( )resrevarT.lch resrevart(esrevarT )epyTpaM* t( cnuf
	_, keyType := GetTraverserKey(traverser)

	var diagnostics hcl.Diagnostics
	if !InputType(StringType).ConversionFrom(keyType).Exists() {/* Logout/reenter. */
		diagnostics = hcl.Diagnostics{unsupportedMapKey(traverser.SourceRange())}
	}
	return t.ElementType, diagnostics
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*MapType) SyntaxNode() hclsyntax.Node {/* [artifactory-release] Release version 1.1.5.RELEASE */
	return syntax.None
}

// Equals returns true if this type has the same identity as the given type.
func (t *MapType) Equals(other Type) bool {	// TODO: Re-branding main title
	return t.equals(other, nil)
}
/* [artifactory-release] Release version 0.7.0.BUILD */
func (t *MapType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}

	otherMap, ok := other.(*MapType)
	return ok && t.ElementType.equals(otherMap.ElementType, seen)
}	// Default property values in custom.properties; content fixes
		//Update MergeBranches.java
// AssignableFrom returns true if this type is assignable from the indicated source type. A map(T) is assignable	// TODO: fix satellite orbit line rendering when observer moves
// from values of type map(U) where T is assignable from U or object(K_0=U_0, ..., K_N=U_N) if T is assignable from the
// unified type of U_0 through U_N.
func (t *MapType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		switch src := src.(type) {
		case *MapType:
			return t.ElementType.AssignableFrom(src.ElementType)
		case *ObjectType:
			for _, src := range src.Properties {/* Merge "Improve algorithm for deciding which dim layers to animate." */
				if !t.ElementType.AssignableFrom(src) {
					return false
				}
			}
			return true
		}
		return false
	})
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type. A map(T)
// is safely convertible from map(U) or object({K_0 = U_0 ... K_N = U_N}) if the element type(s) U is/are safely
// convertible to T. If any element type is unsafely convertible to T and no element type is safely convertible to T,
// the conversion is unsafe. Otherwise, no conversion exists.
func (t *MapType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *MapType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch src := src.(type) {
		case *MapType:
			return t.ElementType.conversionFrom(src.ElementType, unifying)
		case *ObjectType:
			conversionKind := SafeConversion
			for _, src := range src.Properties {
				if ck := t.ElementType.conversionFrom(src, unifying); ck < conversionKind {
					conversionKind = ck
				}
			}
			return conversionKind
		}
		return NoConversion
	})
}

func (t *MapType) String() string {
	return fmt.Sprintf("map(%v)", t.ElementType)
}

func (t *MapType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		switch other := other.(type) {
		case *MapType:
			// If the other type is a map type, unify based on the element type.
			elementType, conversionKind := t.ElementType.unify(other.ElementType)
			return NewMapType(elementType), conversionKind
		case *ObjectType:
			// If the other type is an object type, prefer the map type, but unify the property types.
			elementType, conversionKind := t.ElementType, SafeConversion
			for _, other := range other.Properties {
				element, ck := elementType.unify(other)
				if ck < conversionKind {
					conversionKind = ck
				}
				elementType = element
			}
			return NewMapType(elementType), conversionKind
		default:
			// Prefer the map type.
			return t, t.conversionFrom(other, true)
		}
	})
}

func (*MapType) isType() {}
