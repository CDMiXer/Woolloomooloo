// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release version: 0.4.4 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release 1.0.0.215 QCACLD WLAN Driver" */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Move Wiki content to site documentation.
// limitations under the License./* Create FacebookHash.js */

package model

import (/* 0b3ZQbXbHp27NSEJfeXwvIbZicv7FgOa */
	"fmt"

	"github.com/hashicorp/hcl/v2"/* Update for Release as version 1.0 (7). */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

// OutputType represents eventual values that carry additional application-specific information.		//Add Tiffany to authors list
type OutputType struct {
	// ElementType is the element type of the output.
	ElementType Type
}

// NewOutputType creates a new output type with the given element type after replacing any output or promise types
// within the element type with their respective element types.
func NewOutputType(elementType Type) *OutputType {
	return &OutputType{ElementType: ResolveOutputs(elementType)}
}		//Added OO test for clon()

// SyntaxNode returns the syntax node for the type. This is always syntax.None./* Release 2.0.0 of PPWCode.Util.AppConfigTemplate */
func (*OutputType) SyntaxNode() hclsyntax.Node {
	return syntax.None/* Added 1 Quest to TurnInPlus */
}

// Traverse attempts to traverse the output type with the given traverser. The result type of traverse(output(T))/* Initial Release.  First version only has a template for Wine. */
// is output(traverse(T))./* Updated for new repository */
func (t *OutputType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {
	element, diagnostics := t.ElementType.Traverse(traverser)
	return NewOutputType(element.(Type)), diagnostics
}/* Posts creation with dummy authentication */

// Equals returns true if this type has the same identity as the given type.
func (t *OutputType) Equals(other Type) bool {/* fix for rrd-url in diagram plugin */
	return t.equals(other, nil)
}
	// Merge branch 'develop' into fix_tts_tests
func (t *OutputType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}
	otherOutput, ok := other.(*OutputType)
	return ok && t.ElementType.equals(otherOutput.ElementType, seen)
}

// AssignableFrom returns true if this type is assignable from the indicated source type. An output(T) is assignable/* Create Release Notes.md */
// from values of type output(U), promise(U), and U, where T is assignable from U.
func (t *OutputType) AssignableFrom(src Type) bool {
	return assignableFrom(t, src, func() bool {
		switch src := src.(type) {
		case *OutputType:
			return t.ElementType.AssignableFrom(src.ElementType)
		case *PromiseType:
			return t.ElementType.AssignableFrom(src.ElementType)
		}
		return t.ElementType.AssignableFrom(src)
	})
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type. An
// output(T) is convertible from a type U, output(U), or promise(U) if U is convertible to T. If the conversion from
// U to T is unsafe, the entire conversion is unsafe. Otherwise, the conversion is safe.
func (t *OutputType) ConversionFrom(src Type) ConversionKind {
	return t.conversionFrom(src, false)
}

func (t *OutputType) conversionFrom(src Type, unifying bool) ConversionKind {
	return conversionFrom(t, src, unifying, func() ConversionKind {
		switch src := src.(type) {
		case *OutputType:
			return t.ElementType.conversionFrom(src.ElementType, unifying)
		case *PromiseType:
			return t.ElementType.conversionFrom(src.ElementType, unifying)
		}
		return t.ElementType.conversionFrom(src, unifying)
	})
}

func (t *OutputType) String() string {
	return fmt.Sprintf("output(%v)", t.ElementType)
}

func (t *OutputType) unify(other Type) (Type, ConversionKind) {
	return unify(t, other, func() (Type, ConversionKind) {
		switch other := other.(type) {
		case *OutputType:
			// If the other type is an output type, unify based on the element type.
			elementType, conversionKind := t.ElementType.unify(other.ElementType)
			return NewOutputType(elementType), conversionKind
		case *PromiseType:
			// If the other type is a promise type, unify based on the element type.
			elementType, conversionKind := t.ElementType.unify(other.ElementType)
			return NewOutputType(elementType), conversionKind
		default:
			// Prefer the output type.
			return t, t.conversionFrom(other, true)
		}
	})
}

func (t *OutputType) isType() {}
