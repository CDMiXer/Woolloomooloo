// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* make does like tabs, right? Fix build. */
package hcl2
/* Delete gallery.fr.html */
import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Delete JobCommand.java */

// OutputVariable represents a program- or component-scoped output variable./* Release cascade method. */
type OutputVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type
		//Add indirect quote in Matt 2:22
	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {		//Switched taucs to snow leopard
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {/* Fix login error messages */
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]		//Rename other/GithubCopyRawLink.user.js to other/old/GithubCopyRawLink.user.js
}
/* Minor change + compiled in Release mode. */
// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ/* Release 1.1.0-RC1 */
}
