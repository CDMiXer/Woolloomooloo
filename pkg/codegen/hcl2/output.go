// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by steven@stebalien.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Concertando o POM (adicionando implementação para o lib4j)
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Update basic use of Polyter
package hcl2/* Release BAR 1.1.10 */

import (
	"github.com/hashicorp/hcl/v2"	// selection screen draft added
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Update Custom MQ-Example */
// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node

	syntax *hclsyntax.Block/* Automatically compile templates and cache them in memory */
	typ    model.Type
/* [datastore] Use terminology lookup services for component ID extraction */
	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}		//Update befehle.md

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {/* Release 0.0.12 */
	return ov.typ.Traverse(traverser)
}	// TODO: win32: always use system() to exec ":!" commands

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
