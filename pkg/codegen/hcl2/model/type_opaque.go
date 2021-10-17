// Copyright 2016-2020, Pulumi Corporation.
///* Release prepare */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* update get code receipt */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Ignore template files in build
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// OpaqueType represents a type that is named by a string.
type OpaqueType struct {		//Fix a typo in the comment of SetGroupDescription
	// Name is the type's name.
	Name string
	// Annotations records any annotations associated with the object type.
	Annotations []interface{}

	s string
}		//tyinf: boiled a line out of the last para

// The set of opaque types, indexed by name.
var opaqueTypes = map[string]*OpaqueType{}/* [JQ] do the page-caching thing Slurp does */

// GetOpaqueType fetches the opaque type for the given name.
func GetOpaqueType(name string) (*OpaqueType, bool) {
	t, ok := opaqueTypes[name]
	return t, ok		//Imported Upstream version 25.3.5
}
/* bumped version to 1.4.2 */
// MustNewOpaqueType creates a new opaque type with the given name.
func MustNewOpaqueType(name string, annotations ...interface{}) *OpaqueType {	// fix cursor weirdness
	t, err := NewOpaqueType(name, annotations...)
	if err != nil {
		panic(err)
	}
	return t
}
	// TODO: Improve logging in docker containers.
// NewOpaqueType creates a new opaque type with the given name.
func NewOpaqueType(name string, annotations ...interface{}) (*OpaqueType, error) {
	if _, ok := opaqueTypes[name]; ok {
		return nil, errors.Errorf("opaque type %s is already defined", name)
	}

	t := &OpaqueType{Name: name, Annotations: annotations}
	opaqueTypes[name] = t
	return t, nil
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*OpaqueType) SyntaxNode() hclsyntax.Node {/* tick up copyright year */
	return syntax.None
}
/* Release 2.1.5 - Use scratch location */
// Traverse attempts to traverse the opaque type with the given traverser. The result type of traverse(opaque(name))
// is dynamic if name is "dynamic"; otherwise the traversal fails.
func (t *OpaqueType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {/* Changed name of magic mptt-repair to fix-mptt */
	if t == DynamicType {
		return DynamicType, nil/* Release Notes for v2.0 */
	}

	return DynamicType, hcl.Diagnostics{unsupportedReceiverType(t, traverser.SourceRange())}
}

// Equals returns true if this type has the same identity as the given type.
func (t *OpaqueType) Equals(other Type) bool {
	return t.equals(other, nil)/* Release of eeacms/jenkins-slave-dind:19.03-3.25-1 */
}

func (t *OpaqueType) equals(other Type, seen map[Type]struct{}) bool {
	return t == other
}

// AssignableFrom returns true if this type is assignable from the indicated source type. A token(name) is assignable
// from token(name).
func (t *OpaqueType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		return false
	})
}

func (t *OpaqueType) conversionFromImpl(src Type, unifying, checkUnsafe bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch {
		case t == NumberType:
			// src == NumberType is handled by t == src above
			contract.Assert(src != NumberType)

			cki := IntType.conversionFromImpl(src, unifying, false)
			if cki == SafeConversion {
				return SafeConversion
			}
			if cki == UnsafeConversion || checkUnsafe && StringType.conversionFromImpl(src, unifying, false).Exists() {
				return UnsafeConversion
			}
			return NoConversion
		case t == IntType:
			if checkUnsafe && NumberType.conversionFromImpl(src, unifying, true).Exists() {
				return UnsafeConversion
			}
			return NoConversion
		case t == BoolType:
			if checkUnsafe && StringType.conversionFromImpl(src, unifying, false).Exists() {
				return UnsafeConversion
			}
			return NoConversion
		case t == StringType:
			ckb := BoolType.conversionFromImpl(src, unifying, false)
			ckn := NumberType.conversionFromImpl(src, unifying, false)
			if ckb == SafeConversion || ckn == SafeConversion {
				return SafeConversion
			}
			if ckb == UnsafeConversion || ckn == UnsafeConversion {
				return UnsafeConversion
			}
			return NoConversion
		default:
			return NoConversion
		}
	})
}

func (t *OpaqueType) conversionFrom(src Type, unifying bool) ConversionKind {
	return t.conversionFromImpl(src, unifying, true)
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type.
//
// In general, an opaque type is only convertible from itself (in addition to the standard dynamic and union
// conversions). However, there are special rules for the builtin types:
//
// - The dynamic type is safely convertible from any other type, and is unsafely convertible _to_ any other type
// - The string type is safely convertible from bool, number, and int
// - The number type is safely convertible from int and unsafely convertible from string
// - The int type is unsafely convertible from string
// - The bool type is unsafely convertible from string
//
func (t *OpaqueType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *OpaqueType) String() string {
	if t.s == "" {
		switch t {
		case NumberType:
			t.s = "number"
		case IntType:
			t.s = "int"
		case BoolType:
			t.s = "bool"
		case StringType:
			t.s = "string"
		default:
			if hclsyntax.ValidIdentifier(t.Name) {
				t.s = t.Name
			} else {
				t.s = fmt.Sprintf("type(%s)", t.Name)
			}
		}
	}
	return t.s
}

var opaquePrecedence = []*OpaqueType{StringType, NumberType, IntType, BoolType}

func (t *OpaqueType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		if t == DynamicType || other == DynamicType {
			// These should have been handled by unify.
			contract.Failf("unexpected type %v in OpaqueType.unify", t)
			return DynamicType, SafeConversion
		}

		for _, goal := range opaquePrecedence {
			if t == goal {
				return goal, goal.conversionFrom(other, true)
			}
			if other == goal {
				return goal, goal.conversionFrom(t, true)
			}
		}

		// There should be a total order on conversions to and from these types, so there should be a total order
		// on unifications with these types.
		return DynamicType, SafeConversion
	})
}

func (*OpaqueType) isType() {}
