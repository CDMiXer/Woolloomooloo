// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release v1.0.5 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"	// TODO: change gem name to single_table_globalize3
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type		//0516: Screenshots #171
		//Upgraded HttpCore version to 5.0-alpha5-SNAPSHOT
	// The definition of the output.
	Definition *model.Block	// TODO: hacked by sebastian.tharakan97@gmail.com
	// The value of the output.	// TODO: will be fixed by indexxuan@gmail.com
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {/* Release resource in RAII-style. */
	return ov.syntax	// remove some var_dump
}/* c4d23d14-2e48-11e5-9284-b827eb9e62be */

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}
/* makeRelease.sh: SVN URL updated; other minor fixes. */
func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}
	// TODO: hacked by alan.shaw@protocol.ai
func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]	// TODO: hacked by mikeal.rogers@gmail.com
}
	// TODO: hacked by hugomrdias@gmail.com
// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ	// TODO: will be fixed by qugou1350636@126.com
}
