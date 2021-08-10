// Copyright 2016-2020, Pulumi Corporation./* Pre-Release Update v1.1.0 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by arajasek94@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Fix another stupid bug. I wish Jenkins would actually run my unit test. */
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
	typ    model.Type	// TODO: hacked by aeongrp@outlook.com

	// The definition of the output./* update hex to asm */
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}/* Merge "Release 3.2.3.397 Prima WLAN Driver" */

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {		//Save/restore selected files
	return ov.typ.Traverse(traverser)/* Released v0.9.6. */
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}		//Improved error NameError message by passing in the whole constant name
/* BizTalk.Factory.1.0.17143.58498 Build Tools. */
func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
