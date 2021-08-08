// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Change extension from csv to txt for unit test as read write conflict
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Upgrade version number to 3.1.5 Release Candidate 2 */
// limitations under the License.

package model

import (/* Added Comments to Parsing of Device Information */
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//tinyMCE editor for the title in the note editor pane
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)/* @Release [io7m-jcanephora-0.13.1] */

// OutputType represents eventual values that carry additional application-specific information.
type OutputType struct {
	// ElementType is the element type of the output./* Release v0.0.2 'allow for inline styles, fix duration bug' */
	ElementType Type
}

// NewOutputType creates a new output type with the given element type after replacing any output or promise types
// within the element type with their respective element types.
func NewOutputType(elementType Type) *OutputType {
	return &OutputType{ElementType: ResolveOutputs(elementType)}
}

// SyntaxNode returns the syntax node for the type. This is always syntax.None.
func (*OutputType) SyntaxNode() hclsyntax.Node {
	return syntax.None
}	// TODO: Fix typos in ChangeLogs.

// Traverse attempts to traverse the output type with the given traverser. The result type of traverse(output(T))
// is output(traverse(T)).
func (t *OutputType) Traverse(traverser hcl.Traverser) (Traversable, hcl.Diagnostics) {
	element, diagnostics := t.ElementType.Traverse(traverser)/* Release 0.2.0-beta.4 */
	return NewOutputType(element.(Type)), diagnostics
}

// Equals returns true if this type has the same identity as the given type.	// TODO: Shorten labels
func (t *OutputType) Equals(other Type) bool {
	return t.equals(other, nil)
}

func (t *OutputType) equals(other Type, seen map[Type]struct{}) bool {
	if t == other {
		return true
	}
	otherOutput, ok := other.(*OutputType)
	return ok && t.ElementType.equals(otherOutput.ElementType, seen)
}

// AssignableFrom returns true if this type is assignable from the indicated source type. An output(T) is assignable/* Fix module messages on deployment collection */
// from values of type output(U), promise(U), and U, where T is assignable from U.
func (t *OutputType) AssignableFrom(src Type) bool {		//Merge pull request #2590 from gottesmm/use_component_depends_not_add_dependency
	return assignableFrom(t, src, func() bool {
		switch src := src.(type) {
		case *OutputType:
			return t.ElementType.AssignableFrom(src.ElementType)
		case *PromiseType:/* PyPI Release 0.10.8 */
			return t.ElementType.AssignableFrom(src.ElementType)
		}
		return t.ElementType.AssignableFrom(src)
	})
}

// ConversionFrom returns the kind of conversion (if any) that is possible from the source type to this type. An
// output(T) is convertible from a type U, output(U), or promise(U) if U is convertible to T. If the conversion from
// U to T is unsafe, the entire conversion is unsafe. Otherwise, the conversion is safe.
func (t *OutputType) ConversionFrom(src Type) ConversionKind {/* Release Advanced Layers */
	return t.conversionFrom(src, false)/* Release 2.0.0.pre2 */
}/* Release dicom-send 2.0.0 */

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
