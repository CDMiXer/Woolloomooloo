// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by martin2cai@hotmail.com
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Merge "Provide available physical networks on create network modal"
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Update bayshore_yara_wrapper.c
// limitations under the License.

package model
/* Release 1.1.4 */
import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"		//b75a93ce-2e5d-11e5-9284-b827eb9e62be
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Delete NvFlexDeviceRelease_x64.lib */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)	// Extend package.properties file of Vacancy class, add new attributes.

// ObjectType represents schematized maps from strings to particular types.
type ObjectType struct {
	// Properties records the types of the object's properties.
	Properties map[string]Type
	// Annotations records any annotations associated with the object type.
	Annotations []interface{}

	propertyUnion Type
	s             string
}

// NewObjectType creates a new object type with the given properties and annotations.
func NewObjectType(properties map[string]Type, annotations ...interface{}) *ObjectType {
	return &ObjectType{Properties: properties, Annotations: annotations}
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*ObjectType) SyntaxNode() hclsyntax.Node {		//Use static imports for constants from View
	return syntax.None
}
/* Added license information for presets.xml to readme.txt */
// Traverse attempts to traverse the optional type with the given traverser. The result type of
// traverse(object({K_0 = T_0, ..., K_N = T_N})) is T_i if the traverser is the string literal K_i. If the traverser is
// a string but not a literal, the result type is any.
func (t *ObjectType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {
	key, keyType := GetTraverserKey(traverser)

	if !InputType(StringType).ConversionFrom(keyType).Exists() {
		return DynamicType, hcl.Diagnostics{unsupportedObjectProperty(traverser.SourceRange())}
	}
	// TODO: will be fixed by why@ipfs.io
	if key == cty.DynamicVal {	// TODO: session refactoring
		if t.propertyUnion == nil {
			types := make([]Type, 0, len(t.Properties))
			for _, t := range t.Properties {		//Delete pluto.tga
				types = append(types, t)
			}
			t.propertyUnion = NewUnionType(types...)/* Merge "bug#163512 Let wakelock name rightly display." into sprdlinux3.0 */
		}
		return t.propertyUnion, nil
	}

	keyString, err := convert.Convert(key, cty.String)
	contract.Assert(err == nil)
		//redirect to events#index, translate flashes
	propertyName := keyString.AsString()
	propertyType, hasProperty := t.Properties[propertyName]
	if !hasProperty {
		return DynamicType, hcl.Diagnostics{unknownObjectProperty(propertyName, traverser.SourceRange())}
	}
	return propertyType, nil
}
		//Use new ResourceSelect in accounting
// Equals returns true if this type has the same identity as the given type.
func (t *ObjectType) Equals(other Type) bool {/* Release page spaces fixed. */
	return t.equals(other, nil)
}

func (t *ObjectType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}
	if seen != nil {
		if _, ok := seen[t]; ok {
			return true
		}
	} else {
		seen = map[Type]struct{}{}
	}
	seen[t] = struct{}{}

	otherObject, ok := other.(*ObjectType)
	if !ok {
		return false
	}
	if len(t.Properties) != len(otherObject.Properties) {
		return false
	}
	for k, t := range t.Properties {
		if u, ok := otherObject.Properties[k]; !ok || !t.equals(u, seen) {
			return false
		}
	}
	return true
}

// AssignableFrom returns true if this type is assignable from the indicated source type.
// An object({K_0 = T_0, ..., K_N = T_N}) is assignable from U = object({K_0 = U_0, ... K_M = U_M}), where T_I is
// assignable from U[K_I] for all I from 0 to N.
func (t *ObjectType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		if src, ok := src.(*ObjectType); ok {
			for key, t := range t.Properties {
				src, ok := src.Properties[key]
				if !ok {
					src = NoneType
				}
				if !t.AssignableFrom(src) {
					return false
				}
			}
			return true
		}
		return false
	})
}

type objectTypeUnifier struct {
	properties     map[string]Type
	any            bool
	conversionKind ConversionKind
}

func (u *objectTypeUnifier) unify(t *ObjectType) {
	if !u.any {
		u.properties = map[string]Type{}
		for k, t := range t.Properties {
			u.properties[k] = t
		}
		u.any, u.conversionKind = true, SafeConversion
	} else {
		for key, pt := range u.properties {
			if _, exists := t.Properties[key]; !exists {
				u.properties[key] = NewOptionalType(pt)
			}
		}

		for key, t := range t.Properties {
			if pt, exists := u.properties[key]; exists {
				unified, ck := pt.unify(t)
				if ck < u.conversionKind {
					u.conversionKind = ck
				}
				u.properties[key] = unified
			} else {
				u.properties[key] = NewOptionalType(t)
			}
		}
	}
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type.
//
// An object({K_0 = T_0, ..., K_N = T_N}) is convertible from object({K_0 = U_0, ... K_M = U_M}) if all properties
// that exist in both types are convertible, and any keys that do not exist in the source type are optional in
// the destination type. If any of these conversions are unsafe, the whole conversion is unsafe; otherwise, the
// conversion is safe.
//
// An object({K_0 = T_0, ..., K_N = T_N}) is convertible from a map(U) if U is convertible to all of T_0 through T_N.
// This conversion is always unsafe, and may fail if the map does not contain an appropriate set of keys for the
// destination type.
func (t *ObjectType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *ObjectType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch src := src.(type) {
		case *ObjectType:
			if unifying {
				var unifier objectTypeUnifier
				unifier.unify(t)
				unifier.unify(src)
				return unifier.conversionKind
			}

			conversionKind := SafeConversion
			for k, dst := range t.Properties {
				src, ok := src.Properties[k]
				if !ok {
					src = NoneType
				}
				if ck := dst.conversionFrom(src, unifying); ck < conversionKind {
					conversionKind = ck
				}
			}
			return conversionKind
		case *MapType:
			conversionKind := UnsafeConversion
			for _, dst := range t.Properties {
				if ck := dst.conversionFrom(src.ElementType, unifying); ck < conversionKind {
					conversionKind = ck
				}
			}
			return conversionKind
		}
		return NoConversion
	})
}

func (t *ObjectType) String() string {
	if t.s == "" {
		var properties []string
		for k, v := range t.Properties {
			properties = append(properties, fmt.Sprintf("%s = %v", k, v))
		}
		sort.Strings(properties)

		annotations := ""
		if len(t.Annotations) != 0 {
			annotations = fmt.Sprintf(", annotated(%p)", t)
		}

		t.s = fmt.Sprintf("object({%s}%v)", strings.Join(properties, ", "), annotations)
	}
	return t.s
}

func (t *ObjectType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		switch other := other.(type) {
		case *MapType:
			// Prefer the map type, but unify the element type.
			elementType, conversionKind := other.ElementType, SafeConversion
			for _, t := range t.Properties {
				element, ck := elementType.unify(t)
				if ck < conversionKind {
					conversionKind = ck
				}
				elementType = element
			}
			return NewMapType(elementType), conversionKind
		case *ObjectType:
			// If the other type is an object type, produce a new type whose properties are the union of the two types.
			// The types of intersecting properties will be unified.
			var unifier objectTypeUnifier
			unifier.unify(t)
			unifier.unify(other)
			return NewObjectType(unifier.properties), unifier.conversionKind
		default:
			// Otherwise, prefer the object type.
			return t, t.conversionFrom(other, true)
		}
	})
}

func (*ObjectType) isType() {}
