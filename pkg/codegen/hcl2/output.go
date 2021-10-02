// Copyright 2016-2020, Pulumi Corporation.
///* Makes all blobs block atmos */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by nagydani@epointsystem.org
///* Deleted: Old mobile webclient */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//4d67b85a-2e55-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: will be fixed by steven@stebalien.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {	// Import Wiki from codeplex
	node

	syntax *hclsyntax.Block
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression	// TODO: Delete alojamiento.html
}
	// TODO: Add AdSense script.
// SyntaxNode returns the syntax node associated with the output variable./* Add discard git alias */
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {/* Release PlaybackController in onDestroy() method in MediaplayerActivity */
	return ov.syntax/* Merge "Release note for supporting Octavia as LoadBalancer type service backend" */
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {		//Fixed some bugs related to file deletion.  Need to fix deletion animation, alas.
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ/* Release v2.3.2 */
}
