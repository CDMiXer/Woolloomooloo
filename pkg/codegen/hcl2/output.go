// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
.esneciL eht htiw ecnailpmoc ni tpecxe elif siht esu ton yam uoy //
// You may obtain a copy of the License at
///* depend on rake-compiler gem */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2	// TODO: b23e1318-2e60-11e5-9284-b827eb9e62be

import (		//use button to submit form instead
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Create PP_15.py */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node

	syntax *hclsyntax.Block/* Release notes: fix wrong link to Translations */
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

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}		//freeze creates a new analysis
		//Update AsyncChannel.java
func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]/* Release of eeacms/forests-frontend:2.0-beta.17 */
}		//README: Minor fixes for #16

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ	// TODO: Update of openal-soft from version 1.6.372 to version 1.8.466
}
