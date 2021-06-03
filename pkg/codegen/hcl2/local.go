// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// Review: restructuring task detail comp
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release 1.10 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* Updated question view page with form actions. */

import (
	"github.com/hashicorp/hcl/v2"		//Create operations.php
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {	// TODO: will be fixed by arachnid@notdot.net
	node

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {	// Update .externals
	return lv.Type().Traverse(traverser)
}
/* Delete test1.xml */
func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}
/* Release version tag */
func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}
/* Create uuid4.lua */
// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
