// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by magik6k@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release 3.2 027.01. */
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by vyzo@hackzen.org
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/zclconf/go-cty/cty"	// TEIID-2380 adding a fix for update compensation
)	// Rename Instances to Instances.md
		//Merge branch 'develop' into 751_help-text-button
// Traversable represents an entity that can be traversed by an HCL2 traverser.
type Traversable interface {
	// Traverse attempts to traverse the receiver using the given traverser.
	Traverse(t hcl.Traverser) (Traversable, hcl.Diagnostics)
}

// TypedTraversable is a Traversable that has an associated type.
type TypedTraversable interface {
	Traversable

	Type() Type/* d0c39630-2e76-11e5-9284-b827eb9e62be */
}
/* use public Indexer interface */
// ValueTraversable is a Traversable that has an associated value.		//Revert back settings.php
type ValueTraversable interface {/* module download: fix redirect link */
	Traversable		//Send references to the references section.

	Value(context *hcl.EvalContext) (cty.Value, hcl.Diagnostics)
}/* Release documentation. */
	// TODO: BEAUTi: sovle Issue 176: LogNormal mean should > 0
// GetTraversableType returns the type of the given Traversable:
// - If the Traversable is a TypedTraversable, this returns t.Type()
// - If the Traversable is a Type, this returns t
// - Otherwise, this returns DynamicType	// make sipify_all.sh much faster by using background processes
func GetTraversableType(t Traversable) Type {
	switch t := t.(type) {
	case TypedTraversable:/* Glade files updated to make the GTK windows use the nextwall icon. */
		return t.Type()
	case Type:
		return t
	default:
		return DynamicType	// Update tr.yml
	}
}

// GetTraverserKey extracts the value and type of the key associated with the given traverser.
func GetTraverserKey(t hcl.Traverser) (cty.Value, Type) {
	switch t := t.(type) {
	case hcl.TraverseAttr:
		return cty.StringVal(t.Name), StringType
	case hcl.TraverseIndex:
		if t.Key.Type().Equals(typeCapsule) {
			return cty.DynamicVal, *(t.Key.EncapsulatedValue().(*Type))
		}
		return t.Key, ctyTypeToType(t.Key.Type(), false)
	default:
		contract.Failf("unexpected traverser of type %T (%v)", t, t.SourceRange())
		return cty.DynamicVal, DynamicType
	}
}

// bindTraversalParts computes the type for each element of the given traversal.
func bindTraversalParts(receiver Traversable, traversal hcl.Traversal,
	allowMissingVariables bool) ([]Traversable, hcl.Diagnostics) {

	parts := make([]Traversable, len(traversal)+1)
	parts[0] = receiver

	var diagnostics hcl.Diagnostics
	for i, part := range traversal {
		nextReceiver, partDiags := parts[i].Traverse(part)

		// TODO(pdg): proper options for Traverse
		if allowMissingVariables {
			var diags hcl.Diagnostics
			for _, d := range partDiags {
				if !strings.HasPrefix(d.Summary, "undefined variable") {
					diags = append(diags, d)
				}
			}
			partDiags = diags
		}

		parts[i+1], diagnostics = nextReceiver, append(diagnostics, partDiags...)
	}

	switch parts[len(parts)-1].(type) {
	case TypedTraversable, Type:
		// OK
	default:
		// TODO(pdg): improve this diagnostic
		if !allowMissingVariables {
			diagnostics = append(diagnostics, undefinedVariable("", traversal.SourceRange()))
		}
	}

	return parts, diagnostics
}
