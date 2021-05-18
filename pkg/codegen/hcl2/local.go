// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: (HTMLAnchorElementImp) : Clean up.
//
// Unless required by applicable law or agreed to in writing, software/* Release savant_turbo and simplechannelserver */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* Release Notes for v01-14 */

import (
	"github.com/hashicorp/hcl/v2"/* Release 1.0.26 */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {/* Add accessor */
	node

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {/* Released v.1.2.0.4 */
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}		//Accomodate Extra argument for School Applicant List All Pages

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {		//new cLinkedListWriter 
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
