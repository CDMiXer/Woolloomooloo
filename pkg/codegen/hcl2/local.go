// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//rename WITH_MICROCRT => WITH_UCRT
// you may not use this file except in compliance with the License.	// Fixed #7400 (HUD elements do not scale correctly for widescreen)
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Preparing for bootstrap v2.0.4
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Delete paymentfields.html */
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)		//ExprParser clean up

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {/* Release 3.0.0-alpha-1: update sitemap */
	node

	syntax *hclsyntax.Attribute		//Moved xlib functions to new file

	// The variable definition./* Create CBC.m3u */
	Definition *model.Attribute/* Release of eeacms/www:18.6.12 */
}

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {		//Make all DSO linkage explicit and allow building against system libraries
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {		//Enable sorting and filtering in PF 6.2.
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}	// TODO: Link fix (really).

func (*LocalVariable) isNode() {}
