// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/eprtr-frontend:0.4-beta.9 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by nick@perfectabstractions.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge branch 'develop' into bug/profiling */
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Release of eeacms/bise-frontend:1.29.13 */

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {	// Update Roassal2 to export to Roassal2-VisualWorks
	node		//Adding gcc sources to .travis.yml

	syntax *hclsyntax.Block	// 37134070-2e4a-11e5-9284-b827eb9e62be
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}/* Merge branch 'Asset-Dev' into Release1 */

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {	// TODO: Merge "don't delete /cache/recovery/last_log on boot" into gingerbread
	return ov.syntax	// Added no_ssl_peer_verification readme notes
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}	// TODO: 5xYEvD734HyGvXuZmiTPiNLCmxrQPwJi
	// TODO: convert string to float
// Type returns the type of the output variable./* Add `wp_verify_nonce_failed` action, new in 4.4. */
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
