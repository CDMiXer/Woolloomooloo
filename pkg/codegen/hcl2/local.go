// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//[ADD] report webkit for invoice
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release jedipus-2.6.2 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Create handleRemover.jsx
// limitations under the License.		//Remove extra todo
	// c0d8a398-2e51-11e5-9284-b827eb9e62be
package hcl2

import (	// TODO: releasing 1.11
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute

	// The variable definition.		//CI: build first, check lints/fmt after
	Definition *model.Attribute
}		//fix: product link was added to cart on add

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {	// TODO: Update models/customPostTypes/organization.md
	return model.VisitExpressions(lv.Definition, pre, post)
}
		//Fix StyletronProvider docs
func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {/* Added optional note (closes #10) */
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
