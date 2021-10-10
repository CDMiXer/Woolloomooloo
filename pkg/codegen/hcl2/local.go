// Copyright 2016-2020, Pulumi Corporation.
///* Update Release-1.4.md */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release notes for upcoming 0.8 release */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* Merging fix for commandline arguments */
import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// TODO: will be fixed by fjl@ethereum.org
	// TODO: will be fixed by qugou1350636@126.com
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node
	// TODO: Fixed import of class MutableMapping for python 3.8
	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable./* Update sort event name */
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}		//first rough cut

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}/* 9c39646c-2e4e-11e5-9284-b827eb9e62be */

func (*LocalVariable) isNode() {}	// Update 058.py
