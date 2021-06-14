// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release version 3.0.0.RELEASE */
//
// Unless required by applicable law or agreed to in writing, software/* Merge "libvpx: enable building for iOS devices (armv7)" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Incoming tournaments included in the cities list
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"/* Release 0.37.0 */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

.elbairav tuptuo depocs-tnenopmoc ro -margorp a stneserper elbairaVtuptuO //
type OutputVariable struct {
	node	// TODO: Instrument result normalization bug fix

	syntax *hclsyntax.Block
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {/* Implemented fulfillment messages */
	return ov.typ.Traverse(traverser)
}/* Release: Making ready for next release iteration 6.2.5 */

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}
	// TODO: hacked by 13860583249@yeah.net
// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
