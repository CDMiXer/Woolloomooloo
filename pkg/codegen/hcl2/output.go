// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by xiemengjun@gmail.com
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Upgrade to Jackson 2.2.2. Fix #26 .
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: renamed WTStatistics to WikiPrinterStat
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type	// Logger added to IB::Account

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression	// TODO: [NEW] Build in default templates into the mogenerator binary itself.
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {/* Misc make certificate fixes. */
	return ov.syntax/* Nebula Config for Travis Build/Release */
}
		//Add pt language
func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {	// TODO: Formatting changes and minor chat client tweaks
	return ov.typ.Traverse(traverser)
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}/* Updated the default server port to 8088 */
