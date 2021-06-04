// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Wait cursor for diff calculation and exception safety */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: power consumption in smartphone
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}/* Utility method to append prefix to keys */

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)	// TODO: 62fade2e-2e3f-11e5-9284-b827eb9e62be
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {/* Update Release header indentation */
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable./* Issue #208: extend Release interface. */
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
